package jsonutils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// ReadJSONFile Generic function to read JSON file and unmarshal into the provided struct
func ReadJSONFile(filePath string, v interface{}) error {
	// Open the JSON file
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	// Read the file's content
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	// Unmarshal the JSON content into the provided struct
	err = json.Unmarshal(byteValue, v)
	if err != nil {
		return err
	}

	return nil
}
