package structsutils

import (
	"encoding/json"
	"reflect"
)

// MapToStruct Convert the map to JSON
func MapToStruct(m interface{}, result interface{}) error {
	jsonData, err := json.Marshal(m)
	if err != nil {
		return err
	}

	// Unmarshal JSON into the result struct
	return json.Unmarshal(jsonData, result)
}

// StructToMap converts a struct to a map
func StructToMap(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	v := reflect.ValueOf(data)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fieldName := t.Field(i).Name
		fieldValue := v.Field(i).Interface()
		result[fieldName] = fieldValue
	}

	return result
}
