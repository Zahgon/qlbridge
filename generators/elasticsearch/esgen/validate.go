package esgen

import (
	u "github.com/araddon/gou"

	"github.com/araddon/qlbridge/expr"
	"github.com/araddon/qlbridge/generators/elasticsearch/gentypes"
	"github.com/araddon/qlbridge/rel"
)

var (
	_ = u.EMPTY

	// Ensure our schema implments filter validation
	fakeValidator gentypes.FilterValidate
)

func init() {
	tv := &TypeValidator{}
	fakeValidator = tv.FilterValidate
}

type TypeValidator struct {
	schema gentypes.SchemaColumns
}

func NewValidator(s gentypes.SchemaColumns) *TypeValidator { _ = "STUB: not implemented"; return nil }

func (m *TypeValidator) FilterValidate(stmt *rel.FilterStatement) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *TypeValidator) walkNode(node expr.Node) error {
	_ = "STUB: not implemented"

	// u.Debugf("%d m.expr T:%T  %#v", depth, node, node)
	return nil
}

// We assume included statement has don't its own validation

func (m *TypeValidator) identityNode(n *expr.IdentityNode) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *TypeValidator) urnaryNode(n *expr.UnaryNode) error { _ = "STUB: not implemented"; return nil }

// TODO:   validate that rhs = bool ?

func (m *TypeValidator) booleanNode(bn *expr.BooleanNode) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *TypeValidator) binaryNode(node *expr.BinaryNode) error {
	_ = "STUB: not implemented"

	// Type check binary expression arguments as they must be:
	// Identifier-Operator-Literal
	return nil
}

//rhs := exprValueType(m.schema, node.Args[1])

// es 5 now enforces that lhs, rhs must be same type no mixed

// If left hand is number right hand needs to be number

// the VM supports both = and ==

// ident(0) != literal(1)

// ident CONTAINS literal

// ident LIKE literal

// Build up list of arguments

func (m *TypeValidator) triNode(node *expr.TriNode) error { _ = "STUB: not implemented"; return nil }

func (m *TypeValidator) funcExpr(node *expr.FuncNode) error { _ = "STUB: not implemented"; return nil }
