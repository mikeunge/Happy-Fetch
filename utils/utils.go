package utils

import (
	"math/rand"
	"os"
	"time"
)

// FileExists :: check if a file exists or not
func FileExists(f string) bool {
	_, err := os.Stat(f)
	return err == nil
}

// RandInt :: create a random integer.
func RandInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

// CheckModificationTime :: check if the modification time is smaller than n (-> refresh rate).
func CheckModificationTime(config Config, method string) bool {
	var path string
	var refreshRate int

	if method == "motd" {
		path = config.MotdPath
		refreshRate = 1
	} else {
		path = config.CachePath
		refreshRate = config.CacheRefresh
	}

	fi, err := os.Stat(path)
	if err != nil {
		Error(config, err)
		return false
	}

	// make sure the "ModTime" is not bigger then the refresh rate.
	tNow := time.Now()
	days := fi.ModTime().Sub(tNow).Hours() / 24

	// if modification time lesser then 1 (day) => true
	return int(days) < refreshRate
}
