package time

import (
	"time"
)

// StartOfDay returns a the first time of a day.
func StartOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

// EndOfDay returns a the last time of a day.
func EndOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 23, 59, 59, 999999999, t.Location())
}

// Yesterday returns yesterday time.
func Yesterday() time.Time {
	return time.Now().AddDate(0, 0, -1)
}

// SetInterval executes a function that is repeated at regular intervals
// and returns a channel to clear the interval.
// Usage:
// interval := setInterval(myFunction, 1000, false)
// To clear interval:
// interval <- true
func SetInterval(fct func(), milliseconds int, async bool) chan<- bool {
	interval := time.Duration(milliseconds) * time.Millisecond
	ticker := time.NewTicker(interval)
	clear := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C:
				if async {
					// This won't block
					go fct()
				} else {
					// This will block
					fct()
				}
			case <-clear:
				ticker.Stop()
				return
			}
		}
	}()

	// We return the channel so we can pass in a value to it to clear the interval
	return clear
}

// SetTimeout runs a function after a given period of time.
func SetTimeout(fct func(), milliseconds int) {
	timeout := time.Duration(milliseconds) * time.Millisecond

	// This spawns a goroutine and therefore does not block
	time.AfterFunc(timeout, fct)
}