package conditionals

func Coalesce[T any](val ...*T) *T {
	if val[0] != nil {
		return val[0]
	} else {
		return val[1]
	}
}
