package utils

import (
	"math/rand"
	"time"
)

// .config dir use os.UserConfigDir

// GetFormattedTime returns a time/date string in time.ANSIC format
func GetFormattedTime() string {
	return time.Now().Format(time.ANSIC)
}

// GetRandomSecond returns a random number between 10 and 60
func GetRandomSecond() int {
	rand.Seed(time.Now().UnixNano())
	min := 10
	max := 60

	return rand.Intn(max-min+1) + min
}
