package asserts_test

import (
	"testing"

	"github.com/allanborba/go-utilitaries/asserts"
)

func TestObject_PrimitiveValues(t *testing.T) {
	assertObjectNoError(t, 1, 1)
	assertObjectNoError(t, "1", "1")
	assertObject(t, 1, 2, "expected 1, got 2")
}

func TestObject_StructValues(t *testing.T) {
	assertObjectNoError(t, testStruct{A: "a", b: 1}, testStruct{A: "a", b: 1})
	assertObjectNoError(t, testStruct{A: "a"}, testStruct{A: "a"})
	assertObject(t, testStruct{A: "1"}, testStruct{A: "2"}, "expected {A: 1}, got {A: 2}")
	assertObject(t, testStruct{A: "1", b: 1}, testStruct{A: "2", b: 2}, "expected {A: 1, b: 1}, got {A: 2, b: 2}")

	struct1 := testStruct{A: "10", b: 10}
	struct2 := testStruct{A: "10", b: 20}
	struct3 := testStruct{A: "20", b: 20}
	assertObjectNoError(t, testStruct{A: "1", b: 1, c: &struct1}, testStruct{A: "1", b: 1, c: &struct1})
	assertObject(t, testStruct{A: "1", b: 1, c: &struct1}, testStruct{A: "1", b: 1, c: &struct2}, "expected {c: {b: 10}}, got {c: {b: 20}}")
	assertObject(t, testStruct{A: "1", b: 1, c: &struct1}, testStruct{A: "1", b: 2, c: &struct3}, "expected {b: 1, c: {A: 10, b: 10}}, got {b: 2, c: {A: 20, b: 20}}")

	struct4 := testStruct{A: "10", b: 10, c: &struct1}
	struct5 := testStruct{A: "10", b: 20, c: &struct2}
	assertObject(t, testStruct{A: "1", b: 1, c: &struct4}, testStruct{A: "1", b: 1, c: &struct5}, "expected {c: {b: 10, c: {b: 10}}}, got {c: {b: 20, c: {b: 20}}}")

	assertObject(t, testStruct{d: []int{1, 2}}, testStruct{d: []int{1, 3}}, "expected {d: [1 2]}, got {d: [1 3]}")
}

func TestObjectIgnoringKeys(t *testing.T) {
	assertObjectIgnoringFieldsNoError(t, testStruct{A: "1"}, testStruct{A: "2"}, []string{"A"})
	assertObjectIgnoringFields(t, testStruct{A: "1", b: 1}, testStruct{A: "2", b: 2}, []string{"A"}, "expected {b: 1}, got {b: 2}")
}

type testStruct struct {
	A string
	b int
	c *testStruct
	d []int
}

func assertObjectNoError[T any](t *testing.T, expected, actual T) {
	mockT := NewFakeT()
	asserts.Object(mockT, expected, actual)
	assertNoError(t, mockT)
}

func assertObject[T any](t *testing.T, expected, actual T, errorMsgExpected string) {
	mockT := NewFakeT()
	asserts.Object(mockT, expected, actual)
	assertErrorMsg(t, mockT, errorMsgExpected)
}

func assertObjectIgnoringFieldsNoError[T any](t *testing.T, expected, actual T, keys []string) {
	mockT := NewFakeT()
	asserts.ObjectIgnoringFields(mockT, expected, actual, keys)
	assertNoError(t, mockT)
}

func assertObjectIgnoringFields[T any](t *testing.T, expected, actual T, fieldsToIgnore []string, errorMsgExpected string) {
	mockT := NewFakeT()
	asserts.ObjectIgnoringFields(mockT, expected, actual, fieldsToIgnore)
	assertErrorMsg(t, mockT, errorMsgExpected)
}
