package schema

type (
	// Applyer takes schema writes and applies them.  This is used both as a database
	// is being loaded, and schema is loaded by store as well as responsible for applying
	// schema changes such as Alters.  In distributed db's this is very, very huge part
	// of work so is a very important interface that is under flux.
	Applyer interface {
		// Init initialize the applyer with registry.
		Init(r *Registry)
		// AddOrUpdateOnSchema Add or Update object (Table, Index)
		AddOrUpdateOnSchema(s *Schema, obj interface{}) error
		// Drop an object from schema
		Drop(s *Schema, obj interface{}) error
	}

	// SchemaSourceProvider is factory for creating schema storage
	SchemaSourceProvider func(s *Schema) Source

	// InMemApplyer applies schema changes in memory.  As changes to
	// schema come in (such as ALTER statements, new tables, new databases)
	// we need to apply them to the underlying schema.
	InMemApplyer struct {
		reg          *Registry
		schemaSource SchemaSourceProvider
	}
)

// NewApplyer new in memory applyer.  For distributed db's we would need
// a different applyer (Raft).
func NewApplyer(sp SchemaSourceProvider) Applyer { _ = "STUB: not implemented"; return *new(Applyer) }

// Init store the registry as part of in-mem applyer which needs it.
func (m *InMemApplyer) Init(r *Registry) {
	_ = "STUB: not implemented"

	// AddOrUpdateOnSchema we have a schema change to apply.  A schema change is
	// a new table, index, or whole new schema being registered.  We provide the first
	// argument which is which schema it is being applied to (ie, add table x to schema y).
	return
}

func (m *InMemApplyer) AddOrUpdateOnSchema(s *Schema, v interface{}) error {
	_ = "STUB: not implemented"

	// All Schemas must also have an info-schema
	return nil
}

// The info-schema if new will need an actual store, the provider
// will add it to the schema.

// Find the type of operation being updated.

// Wipe out cache, it is invalid

// s==v means schema has been updated

// since s != v then this is a child schema

// Drop we have a schema change to apply.
func (m *InMemApplyer) Drop(s *Schema, v interface{}) error {
	_ = "STUB: not implemented"

	// Find the type of operation being updated.
	return nil
}

// s==v means schema is being dropped

// s==v means schema is being dropped
