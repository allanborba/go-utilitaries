package asserts

import (
	"fmt"
	"slices"
)

func StringifyedStruct[T any](expected T) string {
	if !IsStruct(expected) {
		return fmt.Sprintf("%v", expected)
	}

	mapped := StructToMap(expected)
	keys := GetFieldNames(expected)
	slices.Sort(keys)
	str := ""

	for _, k := range keys {
		v := mapped[k]
		if v == nil || IsInterfaceNil(v) {
			continue
		}

		if IsStruct(v) {
			str += fmt.Sprintf("%v: %v ", k, StringifyedStruct(v))
		} else {
			str += fmt.Sprintf("%v: %v ", k, v)
		}
	}

	return fmt.Sprint("{ ", str, "}")
}
