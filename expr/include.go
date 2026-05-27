package expr

import (
	"fmt"
)

const (
	maxIncludeDepth = 100
)

var (
	// If we hit max depth
	ErrMaxDepth = fmt.Errorf("Recursive Evaluation Error")
)

// FindIncludes recursively descend down a node looking for all Include identities
func FindIncludes(node Node) []string { _ = "STUB: not implemented"; return nil }

// InlineIncludes take an expression and resolve any includes so that
// the included expression is "Inline"
func InlineIncludes(ctx Includer, n Node) (Node, error) {
	_ = "STUB: not implemented"
	return *new(Node), nil
}

func doInlineIncludes(ctx Includer, n Node, depth int) (Node, error) {
	_ = "STUB: not implemented"
	// We need to make a copy, so we lazily use the To/From pb
	// We need the copy because we are going to mutate this node
	// but AST is assumed to be immuteable, and shared, since we are breaking
	// this contract we copy
	return *new(Node), nil
}

func inlineIncludesDepth(ctx Includer, arg Node, depth int) (Node, error) {
	_ = "STUB: not implemented"
	return *new(Node), nil
}

// FuncNode, BinaryNode, BooleanNode, TriNode, UnaryNode, ArrayNode

//*NumberNode, *IdentityNode, *StringNode, nil,
//*ValueNode, *NullNode:

func resolveInclude(ctx Includer, inc *IncludeNode, depth int) (Node, error) {
	_ = "STUB: not implemented"

	// if inc.inlineExpr != nil {
	// 	return inc.inlineExpr, nil
	// }
	return *new(Node), nil
}

// Now inline, the inlines

func findAllIncludes(node Node, current []string) []string { _ = "STUB: not implemented"; return nil }
