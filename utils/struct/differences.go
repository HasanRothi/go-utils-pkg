package structsutils

import (
	"fmt"
	arrayutils "github.com/HasanRothi/go-utils/utils/array"
	"reflect"
	"strings"
)

// FindDifferencesBetweenStruct Function to compare two structs and find the differences
func FindDifferencesBetweenStruct(a, b interface{}, excludedFields ...string) map[string][2]interface{} {

	differences := make(map[string][2]interface{})

	va := reflect.ValueOf(a)
	vb := reflect.ValueOf(b)

	// Ensure both inputs are structs
	if va.Kind() != reflect.Struct || vb.Kind() != reflect.Struct {
		fmt.Println("Inputs must be structs")
		return differences
	}

	ta := va.Type()

	for i := 0; i < va.NumField(); i++ {
		field := ta.Field(i)
		fieldName := field.Name

		jsonTag := field.Tag.Get("json")
		if jsonTag != "" && jsonTag != "-" {
			fieldName = strings.Split(jsonTag, ",")[0]
		}

		// Check if the field is in the excludedFields
		if excluded := arrayutils.SliceContains(excludedFields, fieldName); excluded {
			continue
		}

		fieldValueA := va.Field(i).Interface()
		fieldValueB := vb.Field(i).Interface()

		// Custom comparison for slices
		if va.Field(i).Kind() == reflect.Slice {
			if !reflect.DeepEqual(fieldValueA, fieldValueB) {
				differences[fieldName] = [2]interface{}{fieldValueA, fieldValueB}
			}
		} else if fieldValueA != fieldValueB {
			differences[fieldName] = [2]interface{}{fieldValueA, fieldValueB}
		}
	}

	return differences
}
