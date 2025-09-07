package collections

type Set[T comparable] struct {
	elemements map[T]struct{}
}

func NewSet[T comparable](elements []T) *Set[T] {
	set := &Set[T]{make(map[T]struct{})}
	for _, el := range elements {
		set.Add(el)
	}
	return set
}

func (this *Set[T]) Add(el T) {
	this.elemements[el] = struct{}{}
}

func (this *Set[T]) Len() int {
	return len(this.elemements)
}
func (this *Set[T]) ToSlice() []T {
	sl := []T{}
	for el := range this.elemements {
		sl = append(sl, el)
	}
	return sl
}
