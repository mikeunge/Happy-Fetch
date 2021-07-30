package quote

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	u "github.com/mikeunge/Happy-Fetch/utils"
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

// CheckMotd :: make sure if motd is enabled and the file exists as well as there is text written in it
func CheckMotd(isEnabled bool, filePath string) bool {
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
		// if so, we will scrap the motd and pull the information new.
		tNow := time.Now()
		days := fi.ModTime().Sub(tNow).Hours() / 24
		if int(days) > 1 {
			return false
		}
	}
	return true
}
