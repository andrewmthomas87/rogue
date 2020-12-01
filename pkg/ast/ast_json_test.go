package ast

import (
	"fmt"
	"reflect"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func TestNewExpressionFromJSON(t *testing.T) {
	expressions := []Expression{
		&Module{
			Name:        "test",
			Definitions: nil,
		},
		&Nil{},
		&Boolean{Value: false},
		&Int32{Value: 23},
		&Float64{Value: 0.0000582},
		&String{Value: "Hello, world!"},
		&ID{Value: "id"},
		&Definition{
			ID:         &ID{Value: "definitionID"},
			Expression: &Int32{Value: -42},
		},
		&Lambda{
			Parameters: []*ID{{Value: "x"}},
			Expression: &ID{Value: "x"},
		},
		&Call{
			ID:        &ID{Value: "test"},
			Arguments: []Expression{&Float64{Value: 123.456}},
		},
		&AnonymousCall{
			Lambda: &Lambda{
				Parameters: []*ID{{Value: "a"}},
				Expression: &ID{Value: "a"},
			},
			Arguments: []Expression{&String{Value: "a's value"}},
		},
	}

	for _, e := range expressions {
		t.Run(fmt.Sprintf("%T", e), func(t *testing.T) {
			b, err := jsoniter.Marshal(e.JSON())
			if err != nil {
				t.Error(err)
			}
			t.Log(string(b))

			parsedExpression, err := NewExpressionFromJSON(jsoniter.Get(b))
			if err != nil {
				t.Error(err)
			}

			expectedTypeOf, actualTypeOf := reflect.TypeOf(e), reflect.TypeOf(parsedExpression)
			if actualTypeOf != expectedTypeOf {
				t.Errorf("expected %v, got %v", expectedTypeOf, actualTypeOf)
			}

			expectedType, actualType := e.Type(), parsedExpression.Type()
			if actualType != expectedType {
				t.Errorf("expected %v, got %v", expectedType, actualType)
			}
		})
	}
	t.Run("invalid type", func(t *testing.T) {
		e := &Boolean{Value: true}
		json := e.JSON()
		json["type"] = 0
		b, err := jsoniter.Marshal(json)
		if err != nil {
			t.Error(err)
		}
		t.Log(string(b))

		_, err = NewExpressionFromJSON(jsoniter.Get(b))
		if err != ErrInvalidJSON {
			t.Errorf("expected %v, got %v", ErrInvalidJSON, err)
		}
	})
}
