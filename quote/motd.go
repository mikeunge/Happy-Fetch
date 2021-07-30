package quote

import (
	"encoding/json"
	"io/ioutil"
	"os"

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
func CheckMotd(config u.Config, writeAfterCheck bool) bool {
	var createMotd = false

	if !config.Motd {
		u.Debug(config, "MOTD is not enabled")
		return false
	}

	if !u.FileExists(config.MotdPath) {
		if !writeAfterCheck {
			u.Info(config, "MOTD does not exist")
			return false
		}
		createMotd = true
	}

	if !createMotd {
		mTime := u.CheckModificationTime(config, "motd")
		// Check if we want to write the file again or nah, if we are trying to write, flip the statement.
		if writeAfterCheck {
			return !mTime
		}
		return mTime
	}

	return true
}
