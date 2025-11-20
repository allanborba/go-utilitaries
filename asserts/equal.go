package asserts

import "reflect"

const ERROR_MSG = "expected %v, got %v"

func Equal[T any](t Tester, expected T, result T) {
	if !reflect.DeepEqual(expected, result) {
		t.Errorf(ERROR_MSG, expected, result)
	}
}
