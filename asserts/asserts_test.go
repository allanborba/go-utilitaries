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
