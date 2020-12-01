package ast

import (
	"fmt"
)

// Type determines the type of an expression.
type Type uint

// Type constants.
const (
	TypeModule = iota + 1
	TypeNil
	TypeBoolean
	TypeInt32
	TypeFloat64
	TypeString
	TypeID
	TypeDefinition
	TypeLambda
	TypeCall
	TypeAnonymousCall
)

// Expression describes expressions.
type Expression interface {
	Type() Type
	JSON() map[string]interface{}
	fmt.Stringer
}

// A Module is a module.
type Module struct {
	Name        string
	Definitions []*Definition
}

// A Nil is a nil literal value.
type Nil struct{}

// A Boolean is a boolean literal value.
type Boolean struct {
	Value bool
}

// An Int32 is a 32-bit integer literal value.
type Int32 struct {
	Value int32
}

// A Float64 is a 64-bit floating point literal value.
type Float64 struct {
	Value float64
}

// A String is a string literal value.
type String struct {
	Value string
}

// An ID is an identifier.
type ID struct {
	Value string
}

// A Definition is a variable definition.
type Definition struct {
	ID         *ID
	Expression Expression
}

// A Lambda is a lambda procedure definition.
type Lambda struct {
	Parameters []*ID
	Expression Expression
}

// A Call is a procedure call.
type Call struct {
	ID        *ID
	Arguments []Expression
}

// An AnonymousCall is an anonymous procedure call.
type AnonymousCall struct {
	Lambda    *Lambda
	Arguments []Expression
}

// Type returns the module type.
func (m *Module) Type() Type {
	return TypeModule
}

// Type returns the nil type.
func (n *Nil) Type() Type {
	return TypeNil
}

// Type returns the boolean type.
func (b *Boolean) Type() Type {
	return TypeBoolean
}

// Type returns the int32 type.
func (i *Int32) Type() Type {
	return TypeInt32
}

// Type returns the float64 type.
func (f *Float64) Type() Type {
	return TypeFloat64
}

// Type returns the string type.
func (s *String) Type() Type {
	return TypeString
}

// Type returns the id type.
func (i *ID) Type() Type {
	return TypeID
}

// Type returns the definition type.
func (d *Definition) Type() Type {
	return TypeDefinition
}

// Type returns the lambda type.
func (l *Lambda) Type() Type {
	return TypeLambda
}

// Type returns the call type.
func (c *Call) Type() Type {
	return TypeCall
}

// Type returns the anonymous call type.
func (c *AnonymousCall) Type() Type {
	return TypeAnonymousCall
}
