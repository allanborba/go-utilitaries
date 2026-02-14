package asserts_test

import (
	"testing"

	"github.com/allanborba/go-utilitaries/asserts"
)

func TestSlices_WhenSlicesAreEqual_ThenNoError(t *testing.T) {
	assertSlicesNewNoError(t, []int{1, 2, 3}, []int{1, 2, 3})
}

func TestSlices_WhenSlicesAreEqualWithDiffOrder_ThenNoError(t *testing.T) {
	assertSlicesNewNoError(t, []int{1, 2, 3}, []int{3, 1, 2})
}

func TestSlices_WhenSlicesOfStringsAreEqual_ThenNoError(t *testing.T) {
	assertSlicesNewNoError(t, []string{"a", "b"}, []string{"b", "a"})
}

func TestSlices_WhenMissingElement_ThenShowMissingError(t *testing.T) {
	assertSlicesNew(t, []int{1, 2, 3}, []int{1, 2}, "\nmissing elements:\n  - 3")
}

func TestSlices_WhenExtraElement_ThenShowExtraError(t *testing.T) {
	assertSlicesNew(t, []int{1, 2}, []int{1, 2, 4}, "\nextra elements:\n  - 4")
}

func TestSlices_WhenMissingAndExtraElements_ThenShowBothErrors(t *testing.T) {
	assertSlicesNew(t, []int{1, 2, 3}, []int{1, 2, 4}, "\nmissing elements:\n  - 3\n\nextra elements:\n  - 4")
}

func TestSlices_WhenDuplicateElements_ThenHandleCorrectly(t *testing.T) {
	assertSlicesNewNoError(t, []int{1, 1, 2}, []int{2, 1, 1})
	assertSlicesNew(t, []int{1, 1, 2}, []int{1, 2, 2}, "\nmissing elements:\n  - 1\n\nextra elements:\n  - 2")
}

func TestSlices_WhenBothEmpty_ThenNoError(t *testing.T) {
	assertSlicesNewNoError(t, []int{}, []int{})
}

func TestSlices_WhenStructsAreEqual_ThenNoError(t *testing.T) {
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

func TestSlices_WhenStructMissing_ThenShowMissingError(t *testing.T) {
	expected := []TestStruct{
		{A: 1, b: "a"},
		{A: 2, b: "b"},
	}
	result := []TestStruct{
		{A: 1, b: "a"},
	}
	assertSlicesNew(t, expected, result, "\nmissing elements:\n  - { A: 2 b: b }")
}

func TestSlices_WhenStructExtra_ThenShowExtraError(t *testing.T) {
	expected := []TestStruct{
		{A: 1, b: "a"},
	}
	result := []TestStruct{
		{A: 1, b: "a"},
		{A: 3, b: "c"},
	}
	assertSlicesNew(t, expected, result, "\nextra elements:\n  - { A: 3 b: c }")
}

func TestSlices_WhenSliceOfStructsHasDiffElement_ThenShowError(t *testing.T) {
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
	assertSlicesNew(t, expected, result, "\nmissing elements:\n  - { A: 4 b: b }\n\nextra elements:\n  - { A: 4 b: Z }")
}

func TestSlices_WhenDuplicateStructs_ThenHandleCorrectly(t *testing.T) {
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
	assertSlicesNew(t, expected, result, "\nmissing elements:\n  - { A: 2 b: c c: { A: 3 b: 2 } }\n\nextra elements:\n  - { A: 4 b: a }")
}

func TestSlices_WhenStructWithNestedStruct_ThenCompareRecursively(t *testing.T) {
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

func TestSlices_WhenStructWithDiffNestedStruct_ThenShowError(t *testing.T) {
	nested1 := TestStruct{A: 10, b: "x"}
	nested2 := TestStruct{A: 10, b: "y"}
	expected := []TestStruct{
		{A: 1, b: "a", c: &nested1},
	}
	result := []TestStruct{
		{A: 1, b: "a", c: &nested2},
	}
	assertSlicesNew(t, expected, result, "\nmissing elements:\n  - { A: 1 b: a c: { A: 10 b: x } }\n\nextra elements:\n  - { A: 1 b: a c: { A: 10 b: y } }")
}

func TestSlices_WhenStructWithSliceField_ThenCompareRecursively(t *testing.T) {
	expected := []TestStructSlice{
		{TestStruct: TestStruct{A: 1}, d: []TestStruct{{A: 10}, {A: 20}}},
	}
	result := []TestStructSlice{
		{TestStruct: TestStruct{A: 1}, d: []TestStruct{{A: 20}, {A: 10}}},
	}
	assertSlicesNewNoError(t, expected, result)
}

func TestSlices_WhenStructWithDiffSliceField_ThenShowError(t *testing.T) {
	expected := []TestStructSlice{
		{TestStruct: TestStruct{A: 1}, d: []TestStruct{{A: 10}, {A: 20}}},
	}
	result := []TestStructSlice{
		{TestStruct: TestStruct{A: 1}, d: []TestStruct{{A: 10}, {A: 30}}},
	}
	mockT := NewFakeT()
	asserts.Slices(mockT, expected, result)
	if mockT.ErrorMsg == "" {
		t.Errorf("expected error but got none")
	}
}

func TestSlices_WhenStructWithNilPointer_ThenNoError(t *testing.T) {
	expected := []TestStruct{
		{A: 1, c: nil},
	}
	result := []TestStruct{
		{A: 1, c: nil},
	}
	assertSlicesNewNoError(t, expected, result)
}

func TestSlices_WhenStructWithSliceFieldEqualDiffOrder_ThenNoError(t *testing.T) {
	expected := []TestStructSlice{
		{TestStruct: TestStruct{A: 1}, d: []TestStruct{{A: 3}, {A: 1}, {A: 2}}},
	}
	result := []TestStructSlice{
		{TestStruct: TestStruct{A: 1}, d: []TestStruct{{A: 1}, {A: 2}, {A: 3}}},
	}
	assertSlicesNewNoError(t, expected, result)
}

func TestSlices_WhenStructWithDiffSliceFieldElements_ThenShowError(t *testing.T) {
	expected := []TestStructSlice{
		{TestStruct: TestStruct{A: 1}, d: []TestStruct{{A: 1}, {A: 2}}},
	}
	result := []TestStructSlice{
		{TestStruct: TestStruct{A: 1}, d: []TestStruct{{A: 1}, {A: 3}}},
	}
	mockT := NewFakeT()
	asserts.Slices(mockT, expected, result)
	if mockT.ErrorMsg == "" {
		t.Errorf("expected error but got none")
	}
}

func assertSlicesNewNoError[T any](t *testing.T, expected, result []T) {
	mockT := NewFakeT()
	asserts.Slices(mockT, expected, result)
	assertNoError(t, mockT)
}

func assertSlicesNew[T any](t *testing.T, expected, result []T, errorMsgExpected string) {
	mockT := NewFakeT()
	asserts.Slices(mockT, expected, result)
	assertErrorMsg(t, mockT, errorMsgExpected)
}
