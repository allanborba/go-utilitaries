package collections

type Set[T comparable] struct {
	elemements map[T]struct{}
}

func NewSet[T comparable](elements []T) *Set[T] {
	set := &Set[T]{}
	return set
}

func (this *Set[T]) Len() int {
	return 0
}
