package asserts_test

import (
	"fmt"
	"testing"
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
	c *TestStruct
}

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
