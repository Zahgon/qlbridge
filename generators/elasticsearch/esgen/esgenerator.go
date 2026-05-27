package esgen

import (
	"time"

	u "github.com/araddon/gou"

	"github.com/araddon/qlbridge/expr"
	"github.com/araddon/qlbridge/generators/elasticsearch/gentypes"
	"github.com/araddon/qlbridge/rel"
)

var (
	// MaxDepth specifies the depth at which we are certain the filter generator is in an endless loop
	// This *shouldn't* happen, but is better than a stack overflow
	MaxDepth = 1000

	_ = u.EMPTY
)

// copy-pasta from entity to avoid the import
// when we actually parameterize this we will need to do it differently anyway
func DayBucket(dt time.Time) int { _ = "STUB: not implemented"; return 0 }

type FilterGenerator struct {
	ts     time.Time
	inc    expr.Includer
	schema gentypes.SchemaColumns
}

func NewGenerator(ts time.Time, inc expr.Includer, s gentypes.SchemaColumns) *FilterGenerator {
	_ = "STUB: not implemented"
	return nil
}

func (fg *FilterGenerator) fieldType(n expr.Node) (*gentypes.FieldType, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (fg *FilterGenerator) Walk(stmt *rel.FilterStatement) (*gentypes.Payload, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//TODO order by -> sort

// expr dispatches to node-type-specific methods
func (fg *FilterGenerator) walkExpr(node expr.Node, depth int) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//u.Debugf("%d fg.expr T:%T  %#v", depth, node, node)

// Urnaries do their own negation

// Also do their own negation

//HACK As a special case support true as "match_all"; we could support
//    false -> MatchNone, but that seems useless and wasteful of ES cpu.

// Convert MissingField errors to a logical `false`

//u.Debugf("depth=%d filters=%s missing field: %s", depth, node, err)

func (fg *FilterGenerator) unaryExpr(node *expr.UnaryNode, depth int) (interface{}, error) {
	_ = "STUB: not implemented"
	//u.Debugf("urnary %v", node.Operator.T.String())
	return nil, nil
}

//u.Debugf("exists err: %q", err)

//u.Debugf("exists %s", ft)

// filters returns a boolean expression
func (fg *FilterGenerator) booleanExpr(bn *expr.BooleanNode, depth int) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Convert MissingField errors to a logical `false`

//u.Debugf("depth=%d filters=%s missing field: %s", depth, fs, err)

// Simply skip missing fields in ORs

// Convert ANDs to false

// Be nice and omit the useless boolean filter since there's only 1 item

func (fg *FilterGenerator) binaryExpr(node *expr.BinaryNode, depth int) (interface{}, error) {
	_ = "STUB: not implemented"
	// Type check binary expression arguments as they must be:
	// Identifier-Operator-Literal
	return nil, nil
}

// the VM supports both = and ==

//return nil, fmt.Errorf("qlindex: == not supported for nested types %q", lhs.String())

// ident(0) != literal(1)

// ident CONTAINS literal

// ident LIKE literal

// Build up list of arguments

func (fg *FilterGenerator) triExpr(node *expr.TriNode, depth int) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// a BETWEEN b AND c
// Type check ternary expression arguments as they must be:
// Identifier(0) BETWEEN Literal(1) AND Literal(2)

func (fg *FilterGenerator) funcExpr(node *expr.FuncNode, depth int) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// see entity.EvalTimeWindow for code implementation. Checks if the contextual time is within the time buckets provided
// by the parameters

//  We are applying the function to the named field, but the caller *can't* just use the fieldname (which would
// evaluate to nothing, as the field isn't
