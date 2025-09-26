package asserts

import (
	"reflect"
	"unsafe"
)

func IsStruct[T any](expected T) bool {
	value := reflect.ValueOf(expected)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	return value.Kind() == reflect.Struct
}

func IsInterfaceNil(value interface{}) bool {
	return value == nil || unsafe.Pointer((*[2]uintptr)(unsafe.Pointer(&value))[1]) == nil
}
