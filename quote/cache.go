package quote

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func WriteCache(filePath string, q []Quote) error {
	file, err := json.MarshalIndent(q, "", " ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filePath, file, 0644)
	if err != nil {
		return err
	}
	return nil
}

func GetCache(filePath string) ([]Quote, error) {
	var q []Quote

	file, err := os.Open(filePath)
	if err != nil {
		return q, err
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&q)
	if err != nil {
		return q, err
	}
	return q, nil
}
