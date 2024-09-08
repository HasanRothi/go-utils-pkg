package structsutils

import "reflect"

// IsEmptyStruct checks if all the fields of a struct (including nested structs) have their zero values
func IsEmptyStruct(v interface{}) bool {
	val := reflect.ValueOf(v)

	// Handle pointers
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// Ensure the input is a struct
	if val.Kind() != reflect.Struct {
		return false
	}

	// Iterate over all fields in the struct
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		// Handle nested structs
		if field.Kind() == reflect.Struct {
			if !IsEmptyStruct(field.Interface()) {
				return false
			}
		} else if !field.IsZero() {
			return false
		}
	}
	return true
}
