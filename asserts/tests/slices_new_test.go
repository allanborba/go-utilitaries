package asserts_test

import (
	"testing"

	"github.com/allanborba/go-utilitaries/asserts"
)

func TestSlicesNew_WhenSlicesAreEqual_ThenNoError(t *testing.T) {
	assertSlicesNewNoError(t, []int{1, 2, 3}, []int{1, 2, 3})
}

func TestSlicesNew_WhenSlicesAreEqualWithDiffOrder_ThenNoError(t *testing.T) {
	assertSlicesNewNoError(t, []int{1, 2, 3}, []int{3, 1, 2})
}

func TestSlicesNew_WhenSlicesOfStringsAreEqual_ThenNoError(t *testing.T) {
	assertSlicesNewNoError(t, []string{"a", "b"}, []string{"b", "a"})
}

func TestSlicesNew_WhenMissingElement_ThenShowMissingError(t *testing.T) {
	assertSlicesNew(t, []int{1, 2, 3}, []int{1, 2}, "missing elements: [3]")
}

func TestSlicesNew_WhenExtraElement_ThenShowExtraError(t *testing.T) {
	assertSlicesNew(t, []int{1, 2}, []int{1, 2, 4}, "extra elements: [4]")
}

func TestSlicesNew_WhenMissingAndExtraElements_ThenShowBothErrors(t *testing.T) {
	assertSlicesNew(t, []int{1, 2, 3}, []int{1, 2, 4}, "missing elements: [3]; extra elements: [4]")
}

func TestSlicesNew_WhenDuplicateElements_ThenHandleCorrectly(t *testing.T) {
	assertSlicesNewNoError(t, []int{1, 1, 2}, []int{2, 1, 1})
	assertSlicesNew(t, []int{1, 1, 2}, []int{1, 2, 2}, "missing elements: [1]; extra elements: [2]")
}

func TestSlicesNew_WhenBothEmpty_ThenNoError(t *testing.T) {
	assertSlicesNewNoError(t, []int{}, []int{})
}

func TestSlicesNew_WhenStructsAreEqual_ThenNoError(t *testing.T) {
	expected := []TestStruct{
		{A: 1, b: "a"},
		{A: 2, b: "b"},
	}
	result := []TestStruct{
		{A: 2, b: "b"},
		{A: 1, b: "a"},
	}
	assertSlicesNewNoError(t, expected, result)
}

func TestSlicesNew_WhenStructMissing_ThenShowMissingError(t *testing.T) {
	expected := []TestStruct{
		{A: 1, b: "a"},
		{A: 2, b: "b"},
	}
	result := []TestStruct{
		{A: 1, b: "a"},
	}
	assertSlicesNew(t, expected, result, "missing elements: [{ A: 2 b: b }]")
}

func TestSlicesNew_WhenStructExtra_ThenShowExtraError(t *testing.T) {
	expected := []TestStruct{
		{A: 1, b: "a"},
	}
	result := []TestStruct{
		{A: 1, b: "a"},
		{A: 3, b: "c"},
	}
	assertSlicesNew(t, expected, result, "extra elements: [{ A: 3 b: c }]")
}

func TestSlicesNew_WhenSliceOfStructsHasDiffElement_ThenShowError(t *testing.T) {
	expected := []TestStruct{
		{A: 1, b: "a"},
		{A: 2, b: "b"},
		{A: 3, b: "b"},
		{A: 4, b: "b"},
		{A: 5, b: "b"},
		{A: 6, b: "b"},
	}
	result := []TestStruct{
		{A: 2, b: "b"},
		{A: 4, b: "Z"},
		{A: 1, b: "a"},
		{A: 3, b: "b"},
		{A: 6, b: "b"},
		{A: 5, b: "b"},
	}
	assertSlicesNew(t, expected, result, "missing elements: [{ A: 4 b: b }]; extra elements: [{ A: 4 b: Z }]")
}

func TestSlicesNew_WhenDuplicateStructs_ThenHandleCorrectly(t *testing.T) {
	expected := []TestStruct{
		{A: 2, b: "c", c: &TestStruct{A: 3, b: "2"}},
		{A: 3, b: "a"},
		{A: 4, b: "a"},
		{A: 2, b: "c", c: &TestStruct{A: 3, b: "2"}},
		{A: 1, b: "a"},
	}
	result := []TestStruct{
		{A: 1, b: "a"},
		{A: 2, b: "c", c: &TestStruct{A: 3, b: "2"}},
		{A: 3, b: "a"},
		{A: 4, b: "a"},
		{A: 4, b: "a"},
	}
	assertSlicesNew(t, expected, result, "missing elements: [{ A: 2 b: c c: { A: 3 b: 2 } }]; extra elements: [{ A: 4 b: a }]")
}

func TestSlicesNew_WhenStructWithNestedStruct_ThenCompareRecursively(t *testing.T) {
	nested1 := TestStruct{A: 10, b: "x"}
	nested2 := TestStruct{A: 10, b: "x"}
	expected := []TestStruct{
		{A: 1, b: "a", c: &nested1},
	}
	result := []TestStruct{
		{A: 1, b: "a", c: &nested2},
	}
	assertSlicesNewNoError(t, expected, result)
}

func TestSlicesNew_WhenStructWithDiffNestedStruct_ThenShowError(t *testing.T) {
	nested1 := TestStruct{A: 10, b: "x"}
	nested2 := TestStruct{A: 10, b: "y"}
	expected := []TestStruct{
		{A: 1, b: "a", c: &nested1},
	}
	result := []TestStruct{
		{A: 1, b: "a", c: &nested2},
	}
	assertSlicesNew(t, expected, result, "missing elements: [{ A: 1 b: a c: { A: 10 b: x } }]; extra elements: [{ A: 1 b: a c: { A: 10 b: y } }]")
}

func TestSlicesNew_WhenStructWithSliceField_ThenCompareRecursively(t *testing.T) {
	expected := []TestStructSlice{
		{TestStruct: TestStruct{A: 1}, d: []TestStruct{{A: 10}, {A: 20}}},
	}
	result := []TestStructSlice{
		{TestStruct: TestStruct{A: 1}, d: []TestStruct{{A: 20}, {A: 10}}},
	}
	assertSlicesNewNoError(t, expected, result)
}

func TestSlicesNew_WhenStructWithDiffSliceField_ThenShowError(t *testing.T) {
	expected := []TestStructSlice{
		{TestStruct: TestStruct{A: 1}, d: []TestStruct{{A: 10}, {A: 20}}},
	}
	result := []TestStructSlice{
		{TestStruct: TestStruct{A: 1}, d: []TestStruct{{A: 10}, {A: 30}}},
	}
	mockT := NewFakeT()
	asserts.SlicesNew(mockT, expected, result)
	if mockT.ErrorMsg == "" {
		t.Errorf("expected error but got none")
	}
}

func TestSlicesNew_WhenStructWithNilPointer_ThenNoError(t *testing.T) {
	expected := []TestStruct{
		{A: 1, c: nil},
	}
	result := []TestStruct{
		{A: 1, c: nil},
	}
	assertSlicesNewNoError(t, expected, result)
}

func TestSlicesNew_WhenStructWithSliceFieldEqualDiffOrder_ThenNoError(t *testing.T) {
	expected := []TestStructSlice{
		{TestStruct: TestStruct{A: 1}, d: []TestStruct{{A: 3}, {A: 1}, {A: 2}}},
	}
	result := []TestStructSlice{
		{TestStruct: TestStruct{A: 1}, d: []TestStruct{{A: 1}, {A: 2}, {A: 3}}},
	}
	assertSlicesNewNoError(t, expected, result)
}

func TestSlicesNew_WhenStructWithDiffSliceFieldElements_ThenShowError(t *testing.T) {
	expected := []TestStructSlice{
		{TestStruct: TestStruct{A: 1}, d: []TestStruct{{A: 1}, {A: 2}}},
	}
	result := []TestStructSlice{
		{TestStruct: TestStruct{A: 1}, d: []TestStruct{{A: 1}, {A: 3}}},
	}
	mockT := NewFakeT()
	asserts.SlicesNew(mockT, expected, result)
	if mockT.ErrorMsg == "" {
		t.Errorf("expected error but got none")
	}
}

func assertSlicesNewNoError[T any](t *testing.T, expected, result []T) {
	mockT := NewFakeT()
	asserts.SlicesNew(mockT, expected, result)
	assertNoError(t, mockT)
}

func assertSlicesNew[T any](t *testing.T, expected, result []T, errorMsgExpected string) {
	mockT := NewFakeT()
	asserts.SlicesNew(mockT, expected, result)
	assertErrorMsg(t, mockT, errorMsgExpected)
}
