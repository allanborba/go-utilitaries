package asserts_test

import (
	"testing"

	"github.com/allanborba/go-utilitaries/asserts"
)

func TestStringifyedStruct_WhenIsNotStruct_ThenShowValue(t *testing.T) {
	expected := "1"
	result := asserts.StringifyedStruct(expected)

	asserts.Equal(t, expected, result)
}

func TestStringifyedStruct_WhenIsStruct_ThenShowValue(t *testing.T) {
	struct1 := TestStruct{A: 1, b: "a"}
	result := asserts.StringifyedStruct(struct1)

	expected := "{ A: 1 b: a }"

	asserts.Equal(t, expected, result)
}

func TestStringifyedStruct_WhenIsStructWithStructField_ThenShowValue(t *testing.T) {
	struct1 := TestStruct{A: 1, b: "a", c: &TestStruct{A: 2, b: "b", c: &TestStruct{A: 3, b: "c"}}}
	result := asserts.StringifyedStruct(struct1)

	expected := "{ A: 1 b: a c: { A: 2 b: b c: { A: 3 b: c } } }"

	asserts.Equal(t, expected, result)
}
