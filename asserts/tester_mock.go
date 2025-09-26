package asserts

type Tester interface {
	Errorf(format string, args ...interface{})
}
