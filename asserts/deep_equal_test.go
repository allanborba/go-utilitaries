package asserts_test

import (
	"testing"

	"github.com/allanborba/utilitaries/asserts"
)

func TestDeepEqual_WhenAreEqualPrimitivesType_ThenNotShowError(t *testing.T) {
	mokingT := NewFakeT()
	asserts.DeepEqual(mokingT, 1, 1)

	assertNoError(t, mokingT)

}

func TestDeepEqual_WhenAreBotEqualPrimitivesType_ThenNotShowError(t *testing.T) {
	mokingT := NewFakeT()
	asserts.DeepEqual(mokingT, 1, 2)

	assertErrorMsg(t, mokingT, "expected 1, got 2")

}

func TestDeepEqual_WhenAreEqualStructsWithSameAttrs_ThenNotShowError(t *testing.T) {
	mokingT := NewFakeT()
	struct1 := TestStruct{A: 1, b: "string"}
	struct2 := TestStruct{A: 1, b: "string"}

	asserts.DeepEqual(mokingT, struct1, struct2)

	assertNoError(t, mokingT)
}

func TestDeepEqual_WhenAreEqualStructsWithOneDiffAttr_ThenShowBothAttrsValues(t *testing.T) {
	mokingT := NewFakeT()
	struct1 := TestStruct{A: 3, b: "1"}
	struct2 := TestStruct{A: 3, b: "2"}

	asserts.DeepEqual(mokingT, struct1, struct2)

	assertErrorMsg(t, mokingT, "expected { b: 1 }, got { b: 2 }")
}

func TestDeepEqual_WhenAreEqualStructsPointersWithOneDiffAttr_ThenShowBothAttrsValues(t *testing.T) {
	mokingT := NewFakeT()
	struct1 := &TestStruct{A: 3, b: "1"}
	struct2 := &TestStruct{A: 3, b: "2"}

	asserts.DeepEqual(mokingT, struct1, struct2)

	assertErrorMsg(t, mokingT, "expected { b: 1 }, got { b: 2 }")
}

func TestDeepEqual_WhenAreEqualStructsWithTwoDiffAttr_ThenShowBothAttrsValues(t *testing.T) {
	mokingT := NewFakeT()
	struct1 := TestStruct{A: 2, b: "1"}
	struct2 := TestStruct{A: 3, b: "2"}

	asserts.DeepEqual(mokingT, struct1, struct2)

	assertErrorMsg(t, mokingT, "expected { A: 2 b: 1 }, got { A: 3 b: 2 }")
}

func TestDeepEqual_WhenStructFieldHasInteralDiff_ThenShowOnlyDiffAttributes(t *testing.T) {
	mokingT := NewFakeT()
	struct1 := TestStruct{A: 2, b: "1", c: &TestStruct{A: 3, b: "2"}}
	struct2 := TestStruct{A: 2, b: "1", c: &TestStruct{A: 3, b: "3"}}

	asserts.DeepEqual(mokingT, struct1, struct2)

	assertErrorMsg(t, mokingT, "expected { c: { b: 2 } }, got { c: { b: 3 } }")
}

func TestDeepEqual_WhenStructFieldIsEqual_ThenShowNoError(t *testing.T) {
	mokingT := NewFakeT()
	struct1 := TestStruct{A: 2, b: "1", c: &TestStruct{A: 3, b: "3"}}
	struct2 := TestStruct{A: 2, b: "1", c: &TestStruct{A: 3, b: "3"}}

	asserts.DeepEqual(mokingT, struct1, struct2)

	assertNoError(t, mokingT)
}

func TestDeepEqual_WhenStructFieldIsEqualAndOtherFieldIsDiff_ThenShowOnlyDiffAttributes(t *testing.T) {
	mokingT := NewFakeT()
	struct1 := TestStruct{A: 2, b: "1", c: &TestStruct{A: 3, b: "3"}}
	struct2 := TestStruct{A: 2, b: "2", c: &TestStruct{A: 3, b: "3"}}

	asserts.DeepEqual(mokingT, struct1, struct2)

	assertErrorMsg(t, mokingT, "expected { b: 1 }, got { b: 2 }")
}
