package asserts_test

import (
	"reflect"
	"testing"

	"github.com/allanborba/utilitaries/asserts"
)

func TestGetFieldNamesOfStruct(t *testing.T) {
	type TestStruct struct {
		A int
		B string
		c bool
	}

	expected := []string{"A", "B", "c"}
	result := asserts.GetFieldNames(TestStruct{})

	asserts.Slices(t, expected, result)
}

func TestConvertStructToMap(t *testing.T) {
	testStruct := struct {
		A int
		b string
	}{
		A: 0,
		b: "string",
	}

	expected := map[string]interface{}{
		"A": 0,
		"b": "string",
	}
	result := asserts.StructToMap(testStruct)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestCovertStructToMap_WhenStructIsPointer(t *testing.T) {
	testStruct := struct {
		A int
		b string
	}{
		A: 0,
		b: "string",
	}

	expected := map[string]interface{}{
		"A": 0,
		"b": "string",
	}
	result := asserts.StructToMap(&testStruct)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
