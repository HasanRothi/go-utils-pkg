package structsutils

import "encoding/json"

func MapToStruct(m map[string]interface{}, result interface{}) error {
	// Convert the map to JSON
	jsonData, err := json.Marshal(m)
	if err != nil {
		return err
	}

	// Unmarshal JSON into the result struct
	return json.Unmarshal(jsonData, result)
}
