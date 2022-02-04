package gojson

import "encoding/json"

func Serialize(i interface{}) []byte {
	jsonData, err := json.Marshal(i)
	if err != nil {
		return nil
	}
	return jsonData
}

func Deserialize(data []byte, i interface{}) error {
	err := json.Unmarshal(data, &i)
	return err
}
