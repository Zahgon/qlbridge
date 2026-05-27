// Package schema implements core Relational Algrebra schema objects such as Table,
// Schema, DataSource, Fields, Headers, Index.
package schema

import (
	"database/sql/driver"
	"sync"
	"time"

	u "github.com/araddon/gou"
	"github.com/golang/protobuf/proto"

	"github.com/araddon/qlbridge/value"
)

var (
	// SchemaRefreshInterval default schema Refresh Interval
	SchemaRefreshInterval = -time.Minute * 5

	// Static list of common field names for describe header on Show, Describe
	EngineFullCols       = []string{"Engine", "Support", "Comment", "Transactions", "XA", "Savepoints"}
	ProdedureFullCols    = []string{"Db", "Name", "Type", "Definer", "Modified", "Created", "Security_type", "Comment", "character_set_client ", "collation_connection", "Database Collation"}
	DescribeFullCols     = []string{"Field", "Type", "Collation", "Null", "Key", "Default", "Extra", "Privileges", "Comment"}
	DescribeFullColMap   = map[string]int{"Field": 0, "Type": 1, "Collation": 2, "Null": 3, "Key": 4, "Default": 5, "Extra": 6, "Privileges": 7, "Comment": 8}
	DescribeCols         = []string{"Field", "Type", "Null", "Key", "Default", "Extra"}
	DescribeColMap       = map[string]int{"Field": 0, "Type": 1, "Null": 2, "Key": 3, "Default": 4, "Extra": 5}
	ShowTableColumns     = []string{"Table", "Table_Type"}
	ShowVariablesColumns = []string{"Variable_name", "Value"}
	ShowDatabasesColumns = []string{"Database"}
	ShowTableColumnMap   = map[string]int{"Table": 0}
	ShowIndexCols        = []string{"Table", "Non_unique", "Key_name", "Seq_in_index", "Column_name", "Collation", "Cardinality", "Sub_part", "Packed", "Null", "Index_type", "Index_comment"}
	DescribeFullHeaders  = NewDescribeFullHeaders()
	DescribeHeaders      = NewDescribeHeaders()

	// We use Fields, and Tables as messages in Schema (SHOW, DESCRIBE)
	_ Message = (*Field)(nil)
	_ Message = (*Table)(nil)

	// Enforce interfaces
	_ SourceTableColumn = (*Table)(nil)

	// Schema In Mem must implement applyer
	_ Applyer = (*InMemApplyer)(nil)

	// Enforce proto marshalling
	_ proto.Marshaler = (*Table)(nil)
	//_ proto.Unmarshaler = (*Table)(nil)

	_ = u.EMPTY
)

const (
	// NoNulls defines if we allow nulls
	NoNulls = false
	// AllowNulls ?
	AllowNulls = true
)

type (
	// DialectWriter knows how to format the schema output specific to a dialect
	// such as postgres, mysql, bigquery all have different identity, value escape characters.
	DialectWriter interface {
		// Dialect ie "mysql", "postgres", "cassandra", "bigquery"
		Dialect() string
		Table(tbl *Table) string
		FieldType(t value.ValueType) string
	}

	// Alter interface for schema storage sources
	Alter interface {
		// DropTable drop given table
		DropTable(table string) error
	}

	// Schema is a "Virtual" Schema and may have multiple different backing sources.
	// - Multiple DataSource(s) (each may be discrete source type such as mysql, elasticsearch, etc)
	// - each schema supplies tables to the virtual table pool
	// - each table name across schemas must be unique (or aliased)
	Schema struct {
		Name          string             // Name of schema
		Conf          *ConfigSource      // source configuration
		DS            Source             // This datasource Interface
		InfoSchema    *Schema            // represent this Schema as sql schema like "information_schema"
		SchemaRef     *Schema            // IF this is infoschema, the schema it refers to
		parent        *Schema            // parent schema (optional) if nested.
		schemas       map[string]*Schema // map[schema-name]:Children Schemas
		tableSchemas  map[string]*Schema // Tables to schema map for parent/child
		tableMap      map[string]*Table  // Tables and their field info, flattened from all child schemas
		tableNames    []string           // List Table names, flattened all schemas into one list
		lastRefreshed time.Time          // Last time we refreshed this schema
		mu            sync.RWMutex       // lock for schema mods
	}

	// Table represents traditional definition of Database Table.  It belongs to a Schema
	// and can be used to create a Datasource used to read this table.
	Table struct {
		TablePb
		Fields         []*Field               // List of Fields, in order
		Context        map[string]interface{} // During schema discovery of underlying source, may need to store additional info
		FieldPositions map[string]int         // Maps name of column to ordinal position in array of []driver.Value's
		FieldMap       map[string]*Field      // Map of Field-name -> Field
		Schema         *Schema                // The schema this is member of
		Source         Source                 // The source
		tblID          uint64                 // internal tableid, hash of table name + schema?
		cols           []string               // array of column names
		lastRefreshed  time.Time              // Last time we refreshed this schema
		rows           [][]driver.Value
	}

	// Field Describes the column info, name, data type, defaults, index, null
	// - dialects (mysql, mongo, cassandra) have their own descriptors for these,
	//   so this is generic meant to be converted to Frontend at runtime
	Field struct {
		idx uint64         // Positional index in array of fields
		row []driver.Value // memoized values of this fields descriptors for describe
		FieldPb
		Context map[string]interface{} // During schema discovery of underlying source, may need to store additional info
	}
	// FieldData is the byte value of a "Described" field ready to write to the wire so we don't have
	// to continually re-serialize it.
	FieldData []byte

	// ConfigSchema is the json/config block for Schema, the data-sources
	// that make up this Virtual Schema.  Must have a name and list
	// of sources to include.
	ConfigSchema struct {
		Name       string   `json:"name"`    // Virtual Schema Name, must be unique
		Sources    []string `json:"sources"` // List of sources , the names of the "Db" in source
		ConfigNode []string `json:"-"`       // List of backend Servers
	}

	// ConfigSource are backend datasources ie : storage/database/csvfiles
	// Each represents a single source type/config.  May belong to more
	// than one schema.
	ConfigSource struct {
		Name         string            `json:"name"`            // Name
		Schema       string            `json:"schema"`          // Schema Name if different than Name, will join existing schema
		SourceType   string            `json:"type"`            // [mysql,elasticsearch,csv,etc] Name in DataSource Registry
		TablesToLoad []string          `json:"tables_to_load"`  // if non empty, only load these tables
		TableAliases map[string]string `json:"table_aliases"`   // if non empty, only load these tables
		Nodes        []*ConfigNode     `json:"nodes"`           // List of nodes
		Hosts        []string          `json:"hosts"`           // List of hosts, replaces older "nodes"
		Settings     u.JsonHelper      `json:"settings"`        // Arbitrary settings specific to each source type
		Partitions   []*TablePartition `json:"partitions"`      // List of partitions per table (optional)
		PartitionCt  uint32            `json:"partition_count"` // Instead of array of per table partitions, raw partition count
	}

	// ConfigNode are Servers/Services, ie a running instance of said Source
	// - each must represent a single source type
	// - normal use is a server, describing partitions of servers
	// - may have arbitrary config info in Settings.
	ConfigNode struct {
		Name     string       `json:"name"`     // Name of this Node optional
		Source   string       `json:"source"`   // Name of source this node belongs to
		Address  string       `json:"address"`  // host/ip
		Settings u.JsonHelper `json:"settings"` // Arbitrary settings
	}
)

// NewSchema create a new empty schema with given name.
func NewSchema(schemaName string) *Schema { _ = "STUB: not implemented"; return nil }

// NewInfoSchema create a new empty schema with given name.
func NewInfoSchema(schemaName string, s *Schema) *Schema { _ = "STUB: not implemented"; return nil }

// NewSchemaSource create a new empty schema with given name and source.
func NewSchemaSource(schemaName string, ds Source) *Schema { _ = "STUB: not implemented"; return nil }

// Since Is this schema object been refreshed within time window described by @dur time ago ?
func (m *Schema) Since(dur time.Duration) bool { _ = "STUB: not implemented"; return false }

// Current Is this schema up to date?
func (m *Schema) Current() bool { _ = "STUB: not implemented"; return false }

// Tables gets list of all tables for this schema.
func (m *Schema) Tables() []string {
	_ = "STUB: not implemented"

	// Table gets Table definition for given table name
	return nil
}

func (m *Schema) Table(tableIn string) (*Table, error) { _ = "STUB: not implemented"; return nil, nil }

// u.Debugf("%p looking up %q", m, tableName)

// Lets see if it is   `schema`.`table` format

// OpenConn get a connection from this schema by table name.
func (m *Schema) OpenConn(tableName string) (Conn, error) {
	_ = "STUB: not implemented"
	return *new(Conn), nil
}

// Schema Find a child Schema for given schema name,
func (m *Schema) Schema(schemaName string) (*Schema, error) {
	_ = "STUB: not implemented"
	// We always lower-case schema names
	return nil, nil
}

// SchemaForTable Find a Schema for given Table
func (m *Schema) SchemaForTable(tableName string) (*Schema, error) {
	_ = "STUB: not implemented"

	// We always lower-case table names
	return nil, nil
}

// addChildSchema add a child schema to this one.  Schemas can be tree-in-nature
// with schema of multiple backend datasources being combined into parent Schema, but each
// child has their own unique defined schema.
func (m *Schema) addChildSchema(child *Schema) { _ = "STUB: not implemented"; return }

/*
// AddSchemaForTable add table.

	func (m *Schema) addSchemaForTable(tableName string, ss *Schema) {
		m.mu.Lock()
		defer m.mu.Unlock()
		m.addschemaForTableUnlocked(tableName, ss)
	}
*/
func (m *Schema) refreshSchemaUnlocked() { _ = "STUB: not implemented"; return }

//u.Debugf("%p:%s  DS T:%T table name %s", m, m.Name, m.DS, tableName)

//u.Infof("schema  %p:%s", ss, ss.Name)

//tbl := ss.tableMap[tableName]
//u.Debugf("s:%p ss:%p add table name %s  tbl:%#v", m, ss, tableName, tbl)

func (m *Schema) dropTable(tbl *Table) error {
	_ = "STUB: not implemented"

	// u.Warnf("%p drop %s %v", m, m.Name, m.Tables())
	// u.Infof("infoschema %#v", m.InfoSchema)
	return nil
}

func (m *Schema) addTable(tbl *Table) error {
	_ = "STUB: not implemented"

	// u.Debugf("schema:%p AddTable %#v", m, tbl)
	return nil
}

// create consistent-hash-id of this table name, and or table+schema

// Assign partitions

//u.Infof("add table: %v partitionct:%v conf:%+v", tbl.Name, tbl.PartitionCt, m.Conf)

func (m *Schema) addschemaForTableUnlocked(tableName string, ss *Schema) {
	_ = "STUB: not implemented"
	return
}

// u.Debugf("%p:%s Schema addschemaForTableUnlocked %q  ", m, m.Name, tableName)

// ignoreable errors

func (m *Schema) loadTable(tableName string) error {
	_ = "STUB: not implemented"

	// u.Infof("%p schema.%v loadTable(%q)", m, m.Name, tableName)
	return nil
}

// Add partitions

// NewTable create a new table for a schema.
func NewTable(table string) *Table { _ = "STUB: not implemented"; return nil }

func (m *Table) init(s *Schema) {
	m.Schema = s
}

// HasField does this table have given field/column?
func (m *Table) HasField(name string) bool { _ = "STUB: not implemented"; return false }

// FieldsAsMessages get list of all fields as interface Message
// used in schema as sql "describe table"
func (m *Table) FieldsAsMessages() []Message { _ = "STUB: not implemented"; return nil }

// Id satisifieds Message Interface
func (m *Table) Id() uint64 {
	_ = "STUB: not implemented"

	// Body satisifies Message Interface
	return 0
}

func (m *Table) Body() interface{} {
	_ = "STUB: not implemented"

	// AddField register a new field
	return nil
}

func (m *Table) AddField(fld *Field) { _ = "STUB: not implemented"; return }

// AddFieldType describe and register a new column
func (m *Table) AddFieldType(name string, valType value.ValueType) {
	_ = "STUB: not implemented"
	return
}

// Column get the Underlying data type.
func (m *Table) Column(col string) (value.ValueType, bool) {
	_ = "STUB: not implemented"
	return *new(value.ValueType), false
}

// SetColumns Explicityly set column names.
func (m *Table) SetColumns(cols []string) { _ = "STUB: not implemented"; return }

//col = strings.ToLower(col)

// SetColumnsFromFields Explicityly set column names from fields.
func (m *Table) SetColumnsFromFields() { _ = "STUB: not implemented"; return }

// Columns list of all column names.
func (m *Table) Columns() []string {
	_ = "STUB: not implemented"

	// AsRows return all fields suiteable as list of values for Describe/Show statements.
	return nil
}

func (m *Table) AsRows() [][]driver.Value { _ = "STUB: not implemented"; return nil }

// SetRows set rows aka values for this table.  Used for schema/testing.
func (m *Table) SetRows(rows [][]driver.Value) {
	_ = "STUB: not implemented"

	// FieldNamesPositions List of Field Names and ordinal position in Column list
	return
}

func (m *Table) FieldNamesPositions() map[string]int { _ = "STUB: not implemented"; return nil }

// Current Is this schema object current?  ie, have we refreshed it from
// source since refresh interval.
func (m *Table) Current() bool { _ = "STUB: not implemented"; return false }

// SetRefreshed update the refreshed date to now.
func (m *Table) SetRefreshed() { _ = "STUB: not implemented"; return }

// Since Is this schema object within time window described by @dur time ago ?
func (m *Table) Since(dur time.Duration) bool { _ = "STUB: not implemented"; return false }

// AddContext add key/value pairs to context (settings, metatadata).
func (m *Table) AddContext(key string, value interface{}) { _ = "STUB: not implemented"; return }

func (m *Table) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewFieldBase(name string, valType value.ValueType, size int, desc string) *Field {
	_ = "STUB: not implemented"
	return nil
}

// You need to over-ride this to change it

func NewField(name string, valType value.ValueType, size int, allowNulls bool, defaultVal driver.Value, key, collation, description string) *Field {
	_ = "STUB: not implemented"
	return nil
}

func (m *Field) ValueType() value.ValueType {
	_ = "STUB: not implemented"
	return *new(value.ValueType)
}
func (m *Field) Id() uint64            { _ = "STUB: not implemented"; return 0 }
func (m *Field) Body() interface{}     { _ = "STUB: not implemented"; return nil }
func (m *Field) AsRow() []driver.Value { _ = "STUB: not implemented"; return nil }

// []string{"Field", "Type", "Collation", "Null", "Key", "Default", "Extra", "Privileges", "Comment"}

// should we send this through a dialect-writer?  bc dialect specific?

// should we put native type in here?

func (m *Field) AddContext(key string, value interface{}) { _ = "STUB: not implemented"; return }

func (m *Field) String() string { _ = "STUB: not implemented"; return "" }

func NewDescribeFullHeaders() []*Field { _ = "STUB: not implemented"; return nil }

//[]string{"Field", "Type", "Collation", "Null", "Key", "Default", "Extra", "Privileges", "Comment"}

func NewDescribeHeaders() []*Field { _ = "STUB: not implemented"; return nil }

//[]string{"Field", "Type",  "Null", "Key", "Default", "Extra"}

func NewSourceConfig(name, sourceType string) *ConfigSource { _ = "STUB: not implemented"; return nil }

func (m *ConfigSource) String() string { _ = "STUB: not implemented"; return "" }
