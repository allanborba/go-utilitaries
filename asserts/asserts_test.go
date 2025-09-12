package asserts_test

import (
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
	f.ErrorMsg = format
}

func TestEqual_WhenValueAreEqual_ThenNotShowError(t *testing.T) {
	mokingT := NewFakeT()
	asserts.Equal(mokingT, 0, 0)

	if mokingT.ErrorMsg != "" {
		t.Errorf("error message should be empty")
	}
}

func TestEqual_WhenValueAreNotEqual_ThenShowNotEqualErrorMsg(t *testing.T) {
	mokingT := NewFakeT()
	asserts.Equal(mokingT, 1, 0)

	if mokingT.ErrorMsg != asserts.ERROR_MSG {
		t.Errorf("expected error message, but got nothing")
	}
}

// func TestDeepEqual_WhenOneAttributeIsNotEqual_ThenShowTheAttributeNotEqualOnMsg(t *testing.T) {
// 	mokingT := NewFakeT()
// 	asserts.DeepEqual(mokingT, struct{ A int }{1}, struct{ A int }{2})

// 	if mokingT.ErrorMsg != "expected {A:1}, got {A:2}" {
// 		t.Errorf("wrong error msg")
// 	}
// }

func TestGetFieldNamesOfStruct(t *testing.T) {
	type TestStruct struct {
		A int
	}

	expected := []string{"A"}
	result := asserts.GetFieldNames(TestStruct{})

	asserts.Slices(t, expected, result)
}
