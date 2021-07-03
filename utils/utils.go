package utils

import (
	"time"
)

// GetFormattedTime returns a time/date string in time.ANSIC format
func GetFormattedTime() string {
	return time.Now().Format(time.ANSIC)
}
