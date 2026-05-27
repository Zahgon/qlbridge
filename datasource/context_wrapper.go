// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package datasource

import (
	"fmt"
	"reflect"
	"time"

	"github.com/araddon/qlbridge/expr"
	"github.com/araddon/qlbridge/value"
)

type ContextWrapper struct {
	val reflect.Value
	s   *state
}

func NewContextWrapper(val interface{}) *ContextWrapper { _ = "STUB: not implemented"; return nil }

func (m *ContextWrapper) Get(key string) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// Now if it's a method, it gets the arguments.

func (m *ContextWrapper) Row() map[string]value.Value { _ = "STUB: not implemented"; return nil }
func (m *ContextWrapper) Ts() time.Time               { _ = "STUB: not implemented"; return *new(time.Time) }

// unwind pointers, etc to find either the value or flag indicating was nil
func findValue(v reflect.Value) (reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), false
}

var zero reflect.Value

type state struct {
	stack []namedvar
	err   error
}

// our stack vars that have come from strings in vm eval engine
//
//	such as "user.Name" will try to find struct value with .Name
type namedvar struct {
	name  string
	value reflect.Value
}

var (
	errorType       = reflect.TypeOf((*error)(nil)).Elem()
	fmtStringerType = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
)

func (s *state) errorf(format string, args ...interface{}) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

// evalFieldChain evaluates .X.Y.Z possibly followed by arguments.
// dot is the environment in which to evaluate arguments, while
// receiver is the value being walked along the chain.
func (s *state) evalFieldChain(dot, receiver reflect.Value, node *expr.IdentityNode, ident []string, args []expr.Node, final reflect.Value) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

// Now if it's a method, it gets the arguments.

// func (s *state) evalFunction(dot reflect.Value, node *expr.IdentityNode, cmd expr.Node, args []expr.Node, final reflect.Value) reflect.Value {
// 	name := node.Text
// 	function, ok := findFunction(name, s.tmpl)
// 	if !ok {
// 		return s.errorf("%q is not a defined function", name)
// 	}
// 	return s.evalCall(dot, function, cmd, name, args, final)
// }

func lowerFieldMatch(fieldName string) func(string) bool { _ = "STUB: not implemented"; return nil }

//u.Debugf("check: %s == %s ?", field, lowerField)

// evalField evaluates an expression like (.Field) or (.Field arg1 arg2).
// The 'final' argument represents the return value from the preceding
// value of the pipeline, if any.
func (s *state) evalField(dot reflect.Value, fieldName string, node expr.Node, args []expr.Node, final, receiver reflect.Value) reflect.Value {
	_ = "STUB: not implemented"

	//u.Debugf("evalField: valid?%v", receiver.IsValid())
	return *new(reflect.Value)
}

//u.Warnf("bailing")

// Unless it's an interface, need to get to a value of type *T to guarantee
// we see all methods of T and *T.

//u.Warnf("unimplemented method: %v", fieldName)

// It's not a method; must be a field of a struct or an element of a map. The receiver must not be nil.

//u.Debugf("fld:%s  receiver kind():%v  val: %v", fieldName, receiver.Kind(), receiver)

// Wow, this is pretty bruttaly expensive
// Iterate over all available fields and read the tag value

// Get the field, returns https://golang.org/pkg/reflect/#StructField

// Get the field tag value

//u.Infof("got field? %v", fieldName, tField)

// field is unexported

// If it's a function, we must call it.

//context reader doesn't care about empty values

// If it's a map, attempt to use the field name as a key.

// switch s.tmpl.option.missingKey {
// case mapInvalid:
// 	// Just use the invalid value.
// case mapZeroValue:
// 	result = reflect.Zero(receiver.Type().Elem())
// case mapError:
// 	s.errorf("map has no entry for key %q", fieldName)
// }

func (s *state) evalCall(dot, fun reflect.Value, node expr.Node, name string, args []expr.Node, final reflect.Value) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

// TODO: This could still be a confusing error; maybe goodFunc should provide info.

// Build the arg list.

// If we have an error that is not nil, stop execution and return that error to the caller.

/*
// evalCall executes a function or method call. If it's a method, fun already has the receiver bound, so
// it looks just like a function call.  The arg list, if non-nil, includes (in the manner of the shell), arg[0]
// as the function itself.
func (s *state) evalCall(dot, fun reflect.Value, node expr.Node, name string, args []expr.Node, final reflect.Value) reflect.Value {
	if args != nil {
		args = args[1:] // Zeroth arg is function name/node; not passed to function.
	}
	typ := fun.Type()
	numIn := len(args)
	if final.IsValid() {
		numIn++
	}
	numFixed := len(args)
	if typ.IsVariadic() {
		numFixed = typ.NumIn() - 1 // last arg is the variadic one.
		if numIn < numFixed {
			s.errorf("wrong number of args for %s: want at least %d got %d", name, typ.NumIn()-1, len(args))
		}
	} else if numIn < typ.NumIn()-1 || !typ.IsVariadic() && numIn != typ.NumIn() {
		s.errorf("wrong number of args for %s: want %d got %d", name, typ.NumIn(), len(args))
	}
	if !goodFunc(typ) {
		// TODO: This could still be a confusing error; maybe goodFunc should provide info.
		s.errorf("can't call method/function %q with %d results", name, typ.NumOut())
	}
	// Build the arg list.
	argv := make([]reflect.Value, numIn)
	// Args must be evaluated. Fixed args first.
	i := 0
	for ; i < numFixed && i < len(args); i++ {
		argv[i] = s.evalArg(dot, typ.In(i), args[i])
	}
	// Now the ... args.
	if typ.IsVariadic() {
		argType := typ.In(typ.NumIn() - 1).Elem() // Argument is a slice.
		for ; i < len(args); i++ {
			argv[i] = s.evalArg(dot, argType, args[i])
		}
	}
	// Add final value if necessary.
	if final.IsValid() {
		t := typ.In(typ.NumIn() - 1)
		if typ.IsVariadic() {
			if numIn-1 < numFixed {
				// The added final argument corresponds to a fixed parameter of the function.
				// Validate against the type of the actual parameter.
				t = typ.In(numIn - 1)
			} else {
				// The added final argument corresponds to the variadic part.
				// Validate against the type of the elements of the variadic slice.
				t = t.Elem()
			}
		}
		argv[i] = s.validateType(final, t)
	}
	result := fun.Call(argv)
	// If we have an error that is not nil, stop execution and return that error to the caller.
	if len(result) == 2 && !result[1].IsNil() {
		s.at(node)
		s.errorf("error calling %s: %s", name, result[1].Interface().(error))
	}
	return result[0]
}
*/

// goodFunc checks that the function or method has the right result signature.
func goodFunc(typ reflect.Type) bool {
	_ = "STUB: not implemented"
	// We allow functions with 1 result or 2 results where the second is an error.
	return false
}
