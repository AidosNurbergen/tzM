package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Data struct {
	A int64 `json:"a"`
	B int64 `json:"b"`
}

func ParseJSON(fileName string) ([]Data, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var data []Data
	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}

	return data, nil
}
