package asserts_test

import (
	"reflect"
	"testing"

	"github.com/allanborba/utilitaries/asserts"
)

func TestGetFieldNamesOfStruct(t *testing.T) {
	expected := []string{"A", "b"}
	result := asserts.GetFieldNames(TestStruct{})

	asserts.Slices(t, expected, result)
}

func TestGetFieldNamesOfPointerStruct(t *testing.T) {
	expected := []string{"A", "b"}
	result := asserts.GetFieldNames(&TestStruct{})

	asserts.Slices(t, expected, result)
}

func TestConvertStructToMap(t *testing.T) {
	expected := map[string]interface{}{"A": 0, "b": "string"}
	result := asserts.StructToMap(TestStruct{A: 0, b: "string"})

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestCovertStructToMap_WhenStructIsPointer(t *testing.T) {
	expected := map[string]interface{}{"A": 0, "b": "string"}
	result := asserts.StructToMap(&TestStruct{A: 0, b: "string"})

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
