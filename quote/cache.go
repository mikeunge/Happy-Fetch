package quote

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	u "github.com/mikeunge/Happy-Fetch/utils"
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

func CheckCache(isEnabled bool, filePath string, refRate int) bool {
	if !isEnabled {
		return false
	}
	if !u.FileExists(filePath) {
		return false
	}
	// Check for errors and check the file size.
	if fi, err := os.Stat(filePath); err != nil || fi.Size() <= 0 {
		return false
	} else {
		// make sure the "ModTime" is not bigger then the refresh rate.
		// if so, we will scrap the cache and pull the information new.
		tNow := time.Now()
		days := fi.ModTime().Sub(tNow).Hours() / 24
		if int(days) > refRate {
			return false
		}
	}
	return true
}
