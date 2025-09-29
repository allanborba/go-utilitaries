package collections

func Contains[T comparable](slice []T, element T) bool {
	for _, item := range slice {
		if item == element {
			return true
		}
	}
	return false
}

func IndexOf[T comparable](slice []T, element T) int {
	for i, item := range slice {
		if item == element {
			return i
		}
	}
	return -1
}

func Remove[T comparable](slice []T, element T) []T {
	newSlice := make([]T, 0)
	for _, item := range slice {
		if item != element {
			newSlice = append(newSlice, item)
		}
	}
	return newSlice
}
