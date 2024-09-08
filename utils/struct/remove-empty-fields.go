package structsutils

import "reflect"

func RemoveEmptyFields(dto interface{}) map[string]interface{} {
	fields := make(map[string]interface{})

	val := reflect.ValueOf(dto)
	if val.Kind() == reflect.Ptr {
		val = val.Elem() // Use Elem to get the value if it's a pointer
	}

	if val.Kind() != reflect.Struct {
		panic("expected a struct or pointer to struct")
	}

	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		// Check if the field is a pointer and is nil
		if field.Kind() == reflect.Ptr && field.IsNil() {
			continue
		}

		// Check if the field is a zero value
		if isZeroValue(field) {
			continue
		}

		// Get the JSON tag for the field
		jsonTag := fieldType.Tag.Get("json")
		if jsonTag == "" {
			jsonTag = fieldType.Name // Default to field name if no JSON tag is present
		}

		// Handle nested structs
		if field.Kind() == reflect.Struct {
			nestedFields := RemoveEmptyFields(field.Interface())
			if len(nestedFields) > 0 {
				fields[jsonTag] = nestedFields
			}
		} else {
			// Add the field to the map with JSON tag as the key
			fields[jsonTag] = field.Interface()
		}
	}

	return fields
}

func isZeroValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.String:
		return v.String() == ""
	case reflect.Slice, reflect.Map, reflect.Chan:
		return v.Len() == 0
	case reflect.Ptr:
		return v.IsNil()
	}
	return false
}
