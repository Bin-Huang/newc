package unittest

import (
	"reflect"
	"testing"
)

func TestRefMode(t *testing.T) {
	value := NewStructRef(false)
	typeName := reflect.TypeOf(value).String()
	if typeName != "*unittest.StructRef" {
		t.Errorf("expected *unittest.StructRef, but got %v", typeName)
	}
}

func TestValueMode(t *testing.T) {
	value := NewStructValue(false)
	typeName := reflect.TypeOf(value).String()
	if typeName != "unittest.StructValue" {
		t.Errorf("expected unittest.StructValue, but got %v", typeName)
	}
}

func TestInitMode(t *testing.T) {
	value := NewStructWithInit(false)
	if value.Debug != true {
		t.Errorf("NewStructWithInit should calling init method")
	}
}

func TestValueInitMode(t *testing.T) {
	value := NewStructValueWithInit(false)
	if value.Debug != true {
		t.Errorf("NewStructValueWithInit should calling init method")
	}
	typeName := reflect.TypeOf(value).String()
	if typeName != "unittest.StructValueWithInit" {
		t.Errorf("expected unittest.StructValueWithInit, but got %v", typeName)
	}
}
