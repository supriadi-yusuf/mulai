package simhelper

import (
	"time"
)

// Wait is function to make a routine sleep.
func Wait(d time.Duration) {
	select {
	case <-time.After(d):
		break
	}
}
