package asserts_test

import (
	"testing"

	"github.com/allanborba/go-utilitaries/asserts"
)

func TestGetFieldNamesOfStruct(t *testing.T) {
	expected := []string{"A", "b", "c"}
	result := asserts.GetFieldNames(TestStruct{})

	asserts.Slices(t, expected, result)
}

func TestGetFieldNamesOfPointerStruct(t *testing.T) {
	expected := []string{"A", "b", "c"}
	result := asserts.GetFieldNames(&TestStruct{})

	asserts.Slices(t, expected, result)
}
