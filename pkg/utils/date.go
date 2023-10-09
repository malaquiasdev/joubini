package utils

import (
	"time"
)

func DateNowSubtractFormated(pattern string, days int) string {
	today := time.Now()
	before := today.AddDate(0, 0, -days)
	return before.Format(pattern)
}
