package utils

import (
	"bytes"
	"encoding/json"
)

//Mapping from Map[string]interface{} to interface{}/struct{}
func Mapping(in, out interface{}) error {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(in)
	if err != nil {
		return err
	}
	err = json.NewDecoder(buf).Decode(out)
	if err != nil {
		return err
	}
	return nil
}
