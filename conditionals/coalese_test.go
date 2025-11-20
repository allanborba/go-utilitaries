package conditionals_test

import (
	"testing"

	"github.com/allanborba/go-utilitaries/asserts"
	"github.com/allanborba/go-utilitaries/conditionals"
)

type fakeStruct struct {
	Val int
}

func TestCoalesce(t *testing.T) {
	val1 := &fakeStruct{1}
	val2 := &fakeStruct{2}
	val3 := &fakeStruct{3}
	var val4 *fakeStruct

	asserts.Object(t, &fakeStruct{1}, conditionals.Coalesce(val1))
	asserts.Object(t, &fakeStruct{2}, conditionals.Coalesce(nil, val2))
	asserts.Object(t, &fakeStruct{3}, conditionals.Coalesce(nil, nil, val3))
	asserts.Object(t, nil, conditionals.Coalesce(nil, nil, nil, val4))
}
