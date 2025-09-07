package conditionals

func Coalesce[T any](values ...*T) *T {
	for _, val := range values {
		if val != nil {
			return val
		}
	}

	return nil
}
