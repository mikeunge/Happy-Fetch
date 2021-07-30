package quote

import (
	"encoding/json"
	"io/ioutil"
	"os"

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

// CheckMotd :: make sure if motd is enabled and the file exists as well as there is text written in it
func CheckCache(config u.Config, writeAfterCheck bool) bool {
	var createCache = false

	if !config.Cache {
		u.Debug(config, "Cache is not enabled")
		return false
	}

	if !u.FileExists(config.CachePath) {
		if !writeAfterCheck {
			u.Info(config, "Cache does not exist")
			return false
		}
		createCache = true
	}

	if !createCache {
		mTime := u.CheckModificationTime(config, "cache")
		// Check if we want to write the file again or nah, if we are trying to write, flip the statement.
		if writeAfterCheck {
			return !mTime
		}
		return mTime
	}

	return true
}

// func CheckCache(config u.Config) bool {
// 	if !config.Cache {
// 		return false
// 	}
// 	if !u.FileExists(config.CachePath) {
// 		return false
// 	}
// 	// Check for errors and check the file size.
// 	if fi, err := os.Stat(config.CachePath); err != nil || fi.Size() <= 0 {
// 		return false
// 	} else {
// 		// make sure the "ModTime" is not bigger then the refresh rate.
// 		// if so, we will scrap the cache and pull the information new.
// 		tNow := time.Now()
// 		days := fi.ModTime().Sub(tNow).Hours() / 24
// 		if int(days) > config.CacheRefresh {
// 			return false
// 		}
// 	}
// 	return true
// }
