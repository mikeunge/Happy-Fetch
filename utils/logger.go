package utils

import (
	"log"
	"os"
)

func StartLogger(path string) bool {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return false
	}
	log.SetOutput(f)
	return true
}

func Debug(c Config, m string) {
	if c.Logging && c.Debug {
		log.Printf("DEBUG: %s\n", m)
	}
}

func Error(c Config, err error) {
	if c.Logging {
		log.Printf("ERROR: %s\n", err.Error())
	}
}

func Warn(c Config, m string) {
	if c.Logging {
		log.Printf("WARN: %s\n", m)
	}
}

func Info(c Config, m string) {
	if c.Logging {
		log.Printf("INFO: %s\n", m)
	}
}
