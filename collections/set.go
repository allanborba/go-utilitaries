package collections

type Set[T comparable] struct {
	elemements map[T]struct{}
}

func NewSet[T comparable](elements []T) *Set[T] {
	set := &Set[T]{make(map[T]struct{})}
	for _, element := range elements {
		set.elemements[element] = struct{}{}
	}
	return set
}

func (this *Set[T]) Len() int {
	return len(this.elemements)
}
