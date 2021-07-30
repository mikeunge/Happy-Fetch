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
