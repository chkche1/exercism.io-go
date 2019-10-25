/*
Package gigasecond implements a simple method for adding a gigasecond to Time
 */
package gigasecond

// import path for the time package from the standard library
import "time"

const Gigasecond = 1e9

func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * Gigasecond)
}
