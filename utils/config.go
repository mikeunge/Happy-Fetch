package utils

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Cache        bool   `toml:"cache"`
	CacheRefresh int    `toml:"cache_refresh"`
	CachePath    string `toml:"cache_path"`
	Motd         bool   `toml:"motd"`
	Logging      bool   `toml:"logging"`
	Debug        bool   `toml:"debug"`
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
		Cache:        false,
		CacheRefresh: 0,
		CachePath:    "",
		Motd:         true,
		Logging:      true,
		Debug:        false,
	}
	return c
}
