package asserts_test

import (
	"fmt"
	"testing"

	"github.com/allanborba/utilitaries/asserts"
)

type FakeT struct {
	ErrorMsg string
}

func NewFakeT() *FakeT {
	return &FakeT{}
}

func (f *FakeT) Errorf(format string, args ...interface{}) {
	f.ErrorMsg = fmt.Sprintf(format, args...)
}

type TestStruct struct {
	A int
	b string
}

func TestEqual_WhenValueAreEqual_ThenNotShowError(t *testing.T) {
	mokingT := NewFakeT()
	asserts.Equal(mokingT, 0, 0)

	assertNoError(t, mokingT)
}

func TestEqual_WhenValueAreNotEqual_ThenShowNotEqualErrorMsg(t *testing.T) {
	mokingT := NewFakeT()
	asserts.Equal(mokingT, 1, 0)

	assertErrorMsg(t, mokingT, "expected 1, got 0")
}

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

/*
testes faltantes
- [] campo com atributo sendo uma struct
- [] campo com atributo sendo um ponteiro
- [] um elemento é nulo e o outro não
*/

func assertNoError(t *testing.T, mokingT *FakeT) {
	if mokingT.ErrorMsg != "" {
		t.Errorf("expected no error msg, got: %v", mokingT.ErrorMsg)
	}
}

func assertErrorMsg(t *testing.T, mokingT *FakeT, errorMsgExpected string) {
	if mokingT.ErrorMsg != errorMsgExpected {
		t.Errorf("wrong error msg: %v", mokingT.ErrorMsg)
	}
}
