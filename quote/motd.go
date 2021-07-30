package quote

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// GetMotd :: get the message of the day (motd) from the provided file
func GetMotd(filePath string) (Quote, error) {
	var q Quote

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

// WriteMotd :: write the given quote to motd
func WriteMotd(filePath string, q Quote) error {
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
