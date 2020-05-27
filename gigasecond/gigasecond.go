package gigasecond

import "time"

const gigasecond = 1000000000 * time.Second

// AddGigasecond adds a billion seconds to the time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigasecond)
}
