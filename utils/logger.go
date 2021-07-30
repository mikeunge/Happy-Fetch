package utils

import (
	"log"
	"os"
)

func StartLogger() {
	f, err := os.OpenFile("/var/log/hpy.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
		os.Exit(-1)
	}
	log.SetOutput(f)
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
