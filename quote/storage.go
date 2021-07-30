package quote

import (
	"fmt"

	u "github.com/mikeunge/Happy-Fetch/utils"
)

// CheckStorage :: make sure if storage option is enabled and the file exists as well as there is text written in it
func CheckStorage(config u.Config, writeAfterCheck bool, storage string) bool {
	var path string
	var enabled bool
	var create = false

	if storage == "motd" {
		enabled = config.Motd
		path = config.MotdPath
	} else {
		enabled = config.Cache
		path = config.CachePath
	}

	if !enabled {
		u.Debug(config, fmt.Sprintf("%s is not enabled", storage))
		return false
	}

	if !u.FileExists(path) {
		if !writeAfterCheck {
			u.Info(config, fmt.Sprintf("%s does not exist", storage))
			return false
		}
		create = true
	}

	if !create {
		mTime := u.CheckModificationTime(config, storage)
		// Check if we want to write the file again or nah, if we are trying to write, flip the statement.
		if writeAfterCheck {
			return !mTime
		}
		return mTime
	}

	return true
}
