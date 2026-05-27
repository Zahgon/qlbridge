package schema

import (
	"sync"

	"github.com/araddon/qlbridge/lex"
)

var (
	// the global data sources registry mutex
	registryMu sync.RWMutex
	// default registry for schema, datasources
	registry *Registry

	// DisableRecover If true, we will not capture/suppress panics.
	// Test only feature hopefully
	DisableRecover bool
)

type (
	// Registry  is a global or namespace registry of datasources and schema.
	// Datasources have a "sourcetype" and define somewhat the driver.
	// Schemas are made up of one or more underlying source-types and have normal
	// schema info about tables etc.
	Registry struct {
		applyer Applyer
		// Map of source name, each source name is name of db-TYPE
		// such as elasticsearch, mongo, csv etc
		sources     map[string]Source
		schemas     map[string]*Schema
		schemaNames []string
		mu          sync.RWMutex
	}
)

// CreateDefaultRegistry create the default registry.
func CreateDefaultRegistry(applyer Applyer) { _ = "STUB: not implemented"; return }

// OpenConn a schema-source Connection, Global open connection function using
// default schema registry.
func OpenConn(schemaName, table string) (Conn, error) {
	_ = "STUB: not implemented"
	return *new(Conn), nil
}

// RegisterSourceType makes a datasource type available by the provided @sourceType
// If Register is called twice with the same name or if source is nil, it panics.
//
// Sources are specific schemas of type csv, elasticsearch, etc containing
// multiple tables.
func RegisterSourceType(sourceType string, source Source) { _ = "STUB: not implemented"; return }

// RegisterSourceAsSchema means you have a datasource, that is going to act
// as a named schema.  ie, this will not be a nested schema with sub-schemas
// and the source will not be re-useable as a source-type.
func RegisterSourceAsSchema(name string, source Source) error {
	_ = "STUB: not implemented"

	// Since registry is a global, lets first lock that.
	return nil
}

// RegisterSchema makes a named schema available by the provided @name
// If Register is called twice with the same name or if source is nil, it panics.
//
// Sources are specific schemas of type csv, elasticsearch, etc containing
// multiple tables.
func RegisterSchema(schema *Schema) error { _ = "STUB: not implemented"; return nil }

// DefaultRegistry get access to the shared/global
// registry of all datasource implementations
func DefaultRegistry() *Registry {
	_ = "STUB: not implemented"

	// NewRegistry create schema registry.
	return nil
}

func NewRegistry(applyer Applyer) *Registry { _ = "STUB: not implemented"; return nil }

func (m *Registry) addSourceType(sourceType string, source Source) {
	_ = "STUB: not implemented"
	return
}

// SchemaDrop removes a schema
func (m *Registry) SchemaDrop(schema, name string, objectType lex.TokenType) error {
	_ = "STUB: not implemented"
	return nil
}

// SchemaRefresh means reload the schema from underlying store.  Possibly
// requires introspection.
func (m *Registry) SchemaRefresh(name string) error { _ = "STUB: not implemented"; return nil }

// Init pre-schema load call any sources that need pre-schema init
func (m *Registry) Init() {
	_ = "STUB: not implemented"
	// TODO:  this is a race, we need a lock on sources
	return
}

// SchemaAddFromConfig means you have a Schema-Source you want to add
func (m *Registry) SchemaAddFromConfig(conf *ConfigSource) error {
	_ = "STUB: not implemented"
	return nil
}

// If we specify a parent schema to add this child schema to

// Schema Get schema for given name.
func (m *Registry) Schema(schemaName string) (*Schema, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// SchemaAdd Add a new Schema
func (m *Registry) SchemaAdd(s *Schema) error { _ = "STUB: not implemented"; return nil }

// SchemaAddChild Add a new Child Schema
func (m *Registry) SchemaAddChild(name string, child *Schema) error {
	_ = "STUB: not implemented"
	return nil
}

// Schemas returns a list of schema names
func (m *Registry) Schemas() []string { _ = "STUB: not implemented"; return nil }

// GetSource Find a DataSource by SourceType
func (m *Registry) GetSource(sourceType string) (Source, error) {
	_ = "STUB: not implemented"
	return *new(Source), nil
}

func (m *Registry) getDepth(depth int, sourceType string) (Source, error) {
	_ = "STUB: not implemented"
	return *new(Source), nil
}

// String describe contents of registry.
func (m *Registry) String() string { _ = "STUB: not implemented"; return "" }

// Create a schema from given named source
// we will find Source for that name and introspect
func discoverSchemaFromSource(s *Schema, applyer Applyer) error {
	_ = "STUB: not implemented"
	return nil
}

// For each table in source schema
