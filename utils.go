package main

import (
	"math/rand"
	"time"
)

// .config dir use os.UserConfigDir

func getFormattedTime() string {
	return time.Now().Format(time.ANSIC)
}

func getRandomSecond() int {
	rand.Seed(time.Now().UnixNano())
	min := 10
	max := 60

	return rand.Intn(max-min+1) + min
}
