package clock

import "fmt"

// Clock keeps track of hours and minutes
type Clock struct {
	Hour   int
	Minute int
}

// New returns a new Clock struct
func New(hours, minutes int) Clock {
	hours = (hours + (minutes / 60)) % 24
	minutes = minutes % 60

	if hours < 0 {
		hours += 24
	}

	if minutes < 0 {
		hours--
		minutes += 60
	}

	return Clock{hours, minutes}
}

// String returns a human readable string
// satifies the "Stringer" interface
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hour, c.Minute)
}

// Add adds minutes
func (c Clock) Add(minutes int) Clock {
	hours := (c.Hour + ((c.Minute + minutes) / 60)) % 24
	minutes = (c.Minute + minutes) % 60

	return Clock{hours, minutes}
}

// Subtract subtracts minutes
func (c Clock) Subtract(minutes int) Clock {
	hours := (c.Hour - (minutes / 60)) % 24
	minutes = c.Minute - (minutes % 60)

	if minutes < 0 {
		hours--
		minutes += 60
	}

	if hours < 0 {
		hours += 24
	}

	return Clock{hours, minutes}
}
