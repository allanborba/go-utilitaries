package conditionals

func Ternary[T any](conditional bool, ifTrue T, ifFalse T) T {
	if conditional {
		return ifTrue
	} else {
		return ifFalse
	}
}
