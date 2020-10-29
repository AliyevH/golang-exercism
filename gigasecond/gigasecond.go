package gigasecond

// import path for the time package from the standard library
import (
	"time"
)

// AddGigasecond add seconds to the given time
func AddGigasecond(t time.Time) time.Time {
	gigasecond := 1e9 * time.Second
	return t.Add(gigasecond)
}
