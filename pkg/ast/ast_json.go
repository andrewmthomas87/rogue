package ast

import (
	"errors"
	"strconv"

	jsoniter "github.com/json-iterator/go"
)

// ErrInvalidJSON signals invalid JSON.
var ErrInvalidJSON = errors.New("invalid JSON")

// NewExpressionFromJSON constructs an Expression from JSON.
func NewExpressionFromJSON(json jsoniter.Any) (Expression, error) {
	switch json.Get("type").ToUint() {
	case TypeModule:
		return NewModuleFromJSON(json)
	case TypeNil:
		return NewNilFromJSON(json)
	case TypeBoolean:
		return NewBooleanFromJSON(json)
	case TypeInt32:
		return NewInt32FromJSON(json)
	case TypeFloat64:
		return NewFloat64FromJSON(json)
	case TypeString:
		return NewStringFromJSON(json)
	case TypeID:
		return NewIDFromJSON(json)
	case TypeDefinition:
		return NewDefinitionFromJSON(json)
	case TypeLambda:
		return NewLambdaFromJSON(json)
	case TypeCall:
		return NewCallFromJSON(json)
	case TypeAnonymousCall:
		return NewAnonymousCallFromJSON(json)
	default:
		return nil, ErrInvalidJSON
	}
}

// NewModuleFromJSON constructs a Module from JSON.
func NewModuleFromJSON(json jsoniter.Any) (*Module, error) {
	if json.Get("type").ToUint() != TypeModule {
		return nil, ErrInvalidJSON
	}

	name := json.Get("name").ToString()

	var jsonDefinitions []jsoniter.Any
	json.Get("definitions").ToVal(&jsonDefinitions)

	definitions := make([]*Definition, len(jsonDefinitions))
	for i, json := range jsonDefinitions {
		d, err := NewDefinitionFromJSON(json)
		if err != nil {
			return nil, err
		}
		definitions[i] = d
	}

	return &Module{
		Name:        name,
		Definitions: definitions,
	}, nil
}

// NewNilFromJSON constructs a Nil from JSON.
func NewNilFromJSON(json jsoniter.Any) (*Nil, error) {
	if json.Get("type").ToUint() != TypeNil {
		return nil, ErrInvalidJSON
	}

	return &Nil{}, nil
}

// NewBooleanFromJSON constructs a Boolean from JSON.
func NewBooleanFromJSON(json jsoniter.Any) (*Boolean, error) {
	if json.Get("type").ToUint() != TypeBoolean {
		return nil, ErrInvalidJSON
	}

	value := json.Get("value").ToBool()
	return &Boolean{Value: value}, nil
}

// NewInt32FromJSON constructs an Int32 from JSON.
func NewInt32FromJSON(json jsoniter.Any) (*Int32, error) {
	if json.Get("type").ToUint() != TypeInt32 {
		return nil, ErrInvalidJSON
	}

	value, err := strconv.ParseInt(json.Get("value").ToString(), 10, 32)
	if err != nil {
		return nil, err
	}
	return &Int32{Value: int32(value)}, nil
}

// NewFloat64FromJSON constructs a Float64 from JSON.
func NewFloat64FromJSON(json jsoniter.Any) (*Float64, error) {
	if json.Get("type").ToUint() != TypeFloat64 {
		return nil, ErrInvalidJSON
	}

	value, err := strconv.ParseFloat(json.Get("value").ToString(), 64)
	if err != nil {
		return nil, err
	}
	return &Float64{Value: float64(value)}, nil
}

// NewStringFromJSON constructs a String from JSON.
func NewStringFromJSON(json jsoniter.Any) (*String, error) {
	if json.Get("type").ToUint() != TypeString {
		return nil, ErrInvalidJSON
	}

	value := json.Get("value").ToString()
	return &String{Value: value}, nil
}

// NewIDFromJSON constructs an ID from JSON.
func NewIDFromJSON(json jsoniter.Any) (*ID, error) {
	if json.Get("type").ToUint() != TypeID {
		return nil, ErrInvalidJSON
	}

	value := json.Get("value").ToString()
	return &ID{Value: value}, nil
}

// NewDefinitionFromJSON constructs a Definition from JSON.
func NewDefinitionFromJSON(json jsoniter.Any) (*Definition, error) {
	if json.Get("type").ToUint() != TypeDefinition {
		return nil, ErrInvalidJSON
	}

	id, err := NewIDFromJSON(json.Get("id"))
	if err != nil {
		return nil, err
	}

	expression, err := NewExpressionFromJSON(json.Get("expression"))
	if err != nil {
		return nil, err
	}

	return &Definition{
		ID:         id,
		Expression: expression,
	}, nil
}

// NewLambdaFromJSON constructs a Lambda from JSON.
func NewLambdaFromJSON(json jsoniter.Any) (*Lambda, error) {
	if json.Get("type").ToUint() != TypeLambda {
		return nil, ErrInvalidJSON
	}

	var jsonParameters []jsoniter.Any
	json.Get("parameters").ToVal(&jsonParameters)

	parameters := make([]*ID, len(jsonParameters))
	for i, json := range jsonParameters {
		p, err := NewIDFromJSON(json)
		if err != nil {
			return nil, err
		}
		parameters[i] = p
	}

	expression, err := NewExpressionFromJSON(json.Get("expression"))
	if err != nil {
		return nil, err
	}

	return &Lambda{
		Parameters: parameters,
		Expression: expression,
	}, nil
}

// NewCallFromJSON constructs a Call from JSON.
func NewCallFromJSON(json jsoniter.Any) (*Call, error) {
	if json.Get("type").ToUint() != TypeCall {
		return nil, ErrInvalidJSON
	}

	id, err := NewIDFromJSON(json.Get("id"))
	if err != nil {
		return nil, err
	}

	var jsonArguments []jsoniter.Any
	json.Get("arguments").ToVal(&jsonArguments)

	arguments := make([]Expression, len(jsonArguments))
	for i, json := range jsonArguments {
		a, err := NewExpressionFromJSON(json)
		if err != nil {
			return nil, err
		}
		arguments[i] = a
	}

	return &Call{
		ID:        id,
		Arguments: arguments,
	}, nil
}

// NewAnonymousCallFromJSON constructs an AnonymousCall from JSON.
func NewAnonymousCallFromJSON(json jsoniter.Any) (*AnonymousCall, error) {
	if json.Get("type").ToUint() != TypeAnonymousCall {
		return nil, ErrInvalidJSON
	}

	lambda, err := NewLambdaFromJSON(json.Get("lambda"))
	if err != nil {
		return nil, err
	}

	var jsonArguments []jsoniter.Any
	json.Get("arguments").ToVal(&jsonArguments)

	arguments := make([]Expression, len(jsonArguments))
	for i, json := range jsonArguments {
		a, err := NewExpressionFromJSON(json)
		if err != nil {
			return nil, err
		}
		arguments[i] = a
	}

	return &AnonymousCall{
		Lambda:    lambda,
		Arguments: arguments,
	}, nil
}

// JSON returns a JSON representation of the Module.
func (m *Module) JSON() map[string]interface{} {
	definitions := make([]map[string]interface{}, len(m.Definitions))
	for i, d := range m.Definitions {
		definitions[i] = d.JSON()
	}

	return map[string]interface{}{
		"type":        TypeModule,
		"definitions": definitions,
	}
}

// JSON returns a JSON representation of the Nil.
func (n *Nil) JSON() map[string]interface{} {
	return map[string]interface{}{"type": TypeNil}
}

// JSON returns a JSON representation of the Boolean.
func (b *Boolean) JSON() map[string]interface{} {
	return map[string]interface{}{
		"type":  TypeBoolean,
		"value": b.Value,
	}
}

// JSON returns a JSON representation of the Int32.
func (i *Int32) JSON() map[string]interface{} {
	return map[string]interface{}{
		"type":  TypeInt32,
		"value": strconv.FormatInt(int64(i.Value), 10),
	}
}

// JSON returns a JSON representation of the Float64.
func (f *Float64) JSON() map[string]interface{} {
	return map[string]interface{}{
		"type":  TypeFloat64,
		"value": strconv.FormatFloat(f.Value, 'f', -1, 64),
	}
}

// JSON returns a JSON representation of the String.
func (s *String) JSON() map[string]interface{} {
	return map[string]interface{}{
		"type":  TypeString,
		"value": s.Value,
	}
}

// JSON returns a JSON representation of the ID.
func (i *ID) JSON() map[string]interface{} {
	return map[string]interface{}{
		"type":  TypeID,
		"value": i.Value,
	}
}

// JSON returns a JSON representation of the Definition.
func (d *Definition) JSON() map[string]interface{} {
	return map[string]interface{}{
		"type":       TypeDefinition,
		"id":         d.ID.JSON(),
		"expression": d.Expression.JSON(),
	}
}

// JSON returns a JSON representation of the Lambda.
func (l *Lambda) JSON() map[string]interface{} {
	parameters := make([]map[string]interface{}, len(l.Parameters))
	for i, p := range l.Parameters {
		parameters[i] = p.JSON()
	}

	return map[string]interface{}{
		"type":       TypeLambda,
		"parameters": parameters,
		"expression": l.Expression.JSON(),
	}
}

// JSON returns a JSON representation of the Call.
func (c *Call) JSON() map[string]interface{} {
	arguments := make([]map[string]interface{}, len(c.Arguments))
	for i, a := range c.Arguments {
		arguments[i] = a.JSON()
	}

	return map[string]interface{}{
		"type":      TypeCall,
		"id":        c.ID.JSON(),
		"arguments": arguments,
	}
}

// JSON returns a JSON representation of the AnonymousCall.
func (c *AnonymousCall) JSON() map[string]interface{} {
	arguments := make([]map[string]interface{}, len(c.Arguments))
	for i, a := range c.Arguments {
		arguments[i] = a.JSON()
	}

	return map[string]interface{}{
		"type":      TypeAnonymousCall,
		"lambda":    c.Lambda.JSON(),
		"arguments": arguments,
	}
}

func (m *Module) String() string {
	b, err := jsoniter.Marshal(m.JSON())
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func (n *Nil) String() string {
	b, err := jsoniter.Marshal(n.JSON())
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func (b *Boolean) String() string {
	bytes, err := jsoniter.Marshal(b.JSON())
	if err != nil {
		return err.Error()
	}
	return string(bytes)
}

func (i *Int32) String() string {
	b, err := jsoniter.Marshal(i.JSON())
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func (f *Float64) String() string {
	b, err := jsoniter.Marshal(f.JSON())
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func (s *String) String() string {
	b, err := jsoniter.Marshal(s.JSON())
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func (i *ID) String() string {
	b, err := jsoniter.Marshal(i.JSON())
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func (d *Definition) String() string {
	b, err := jsoniter.Marshal(d.JSON())
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func (l *Lambda) String() string {
	b, err := jsoniter.Marshal(l.JSON())
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func (c *Call) String() string {
	b, err := jsoniter.Marshal(c.JSON())
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func (c *AnonymousCall) String() string {
	b, err := jsoniter.Marshal(c.JSON())
	if err != nil {
		return err.Error()
	}
	return string(b)
}
