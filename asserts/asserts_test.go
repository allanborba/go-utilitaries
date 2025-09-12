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

func TestAssertEqual_WhenValueAreEqual_ThenNotShowError(t *testing.T) {
	mokingT := NewFakeT()
	asserts.Equal(mokingT, 0, 0)

	if mokingT.ErrorMsg != "" {
		t.Errorf("error message should be empty")
	}
}

func TestAssertEqual_WhenValueAreNotEqual_ThenShowNotEqualErrorMsg(t *testing.T) {
	mokingT := NewFakeT()
	asserts.Equal(mokingT, 1, 0)

	if mokingT.ErrorMsg != "expected %v, got %v" {
		t.Errorf("expected error message, but got nothing")
	}
}
