package structsutils

import (
	"fmt"
	"reflect"
)

// Print Struct With JSON tag
func Print(s interface{}) {
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	if v.Kind() != reflect.Struct {
		fmt.Println("Provided interface is not a struct")
		return
	}

	fmt.Println("Struct fields with JSON tags:")
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		// Get the JSON tag value
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			jsonTag = field.Name
		}

		fmt.Printf("%s: %v\n", jsonTag, value.Interface())
	}
}
