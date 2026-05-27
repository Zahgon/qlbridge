package builtins

import (
	u "github.com/araddon/gou"

	"github.com/araddon/qlbridge/expr"
	"github.com/araddon/qlbridge/value"
)

var _ = u.EMPTY

// hash.sip() hash a value to a 64 bit int
//
//	hash.sip("/blog/index.html")  =>  1234
type HashSip struct{}

// Type int
func (m *HashSip) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *HashSip) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func hashSipEval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// HashMd5Func Hash a value to MD5 string
//
//	hash.md5("/blog/index.html")  =>  abc345xyz
type HashMd5 struct{}

// Type string
func (m *HashMd5) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *HashMd5) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func hashMd5Eval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// HashSha1Func Hash a value to SHA256 string
//
//	hash.sha1("/blog/index.html")  =>  abc345xyz
type HashSha1 struct{}

// Type string
func (m *HashSha1) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *HashSha1) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func hashSha1Eval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// HashSha256Func Hash a value to SHA256 string
//
//	hash.sha256("/blog/index.html")  =>  abc345xyz
type HashSha256 struct{}

// Type string
func (m *HashSha256) Type() value.ValueType {
	_ = "STUB: not implemented"
	return *new(value.ValueType)
}
func (m *HashSha256) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func hashSha256Eval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// HashSha512Func Hash a value to SHA512 string
//
//	hash.sha512("/blog/index.html")  =>  abc345xyz
type HashSha512 struct{}

// Type string
func (m *HashSha512) Type() value.ValueType {
	_ = "STUB: not implemented"
	return *new(value.ValueType)
}
func (m *HashSha512) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func hashSha512Eval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// Base 64 encoding function
//
//	encoding.b64encode("hello world=")  =>  aGVsbG8gd29ybGQ=
type EncodeB64Encode struct{}

// Type string
func (m *EncodeB64Encode) Type() value.ValueType {
	_ = "STUB: not implemented"
	return *new(value.ValueType)
}
func (m *EncodeB64Encode) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func encodeB64EncodeEval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// Base 64 encoding function
//
//	encoding.b64decode("aGVsbG8gd29ybGQ=")  =>  "hello world"
type EncodeB64Decode struct{}

// Type string
func (m *EncodeB64Decode) Type() value.ValueType {
	_ = "STUB: not implemented"
	return *new(value.ValueType)
}
func (m *EncodeB64Decode) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func encodeB64DecodeEval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}
