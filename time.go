package goutils

import (
	"time"
)

// IsFutureTime - returns true if supplied time is a future time
func IsFutureTime(t time.Time) bool {
	return t.After(time.Now())
}

// IsPastTime - returns true if supplied time is a past time
func IsPastTime(t time.Time) bool {
	return t.Before(time.Now())
}
