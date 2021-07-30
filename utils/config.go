package utils

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Cache         bool   `toml:"cache"`
	CacheRefresh  int    `toml:"cache_refresh"`
	CachePath     string `toml:"cache_path"`
	Motd          bool   `toml:"motd"`
	MotdPath      string `toml:"motd_path"`
	Logging       bool   `toml:"logging"`
	LoggingPath   string `toml:"logging_path"`
	IgnoreLogging bool   `toml:"ignore_logging_if_fail"`
	Debug         bool   `toml:"debug"`
}

// ConfigReader :: reads the init.toml file and maps the data with the Config struct
func ConfigReader(filePath string) (Config, error) {
	var c Config

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return c, err
	}

	err = toml.Unmarshal(file, &c)
	if err != nil {
		return c, err
	}
	return c, nil
}

// ConfigDefault :: returns Config struct with the "original" data
func ConfigDefault() Config {
	c := Config{
		Cache:         true,
		CacheRefresh:  3,
		CachePath:     "./.hpy",
		Motd:          true,
		MotdPath:      "/tmp/hpy.json",
		Logging:       true,
		LoggingPath:   "/var/log/hpy.log",
		IgnoreLogging: false,
		Debug:         false,
	}
	return c
}
