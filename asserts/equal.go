package asserts

const ERROR_MSG = "expected %v, got %v"

func Equal[T comparable](t Tester, expected T, result T) {
	if result != expected {
		t.Errorf(ERROR_MSG, expected, result)
	}
}
