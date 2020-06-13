package clock

import "fmt"

// Clock keeps track of hours and minutes
type Clock struct {
	Hour   int
	Minute int
}

// {8, 0, "08:00"},        // on the hour
// {11, 9, "11:09"},       // past the hour
// {24, 0, "00:00"},       // midnight is zero hours
// {25, 0, "01:00"},       // hour rolls over
// {100, 0, "04:00"},      // hour rolls over continuously
// {1, 60, "02:00"},       // sixty minutes is next hour
// {0, 160, "02:40"},      // minutes roll over
// {0, 1723, "04:43"},     // minutes roll over continuously
// {25, 160, "03:40"},     // hour and minutes roll over
// {201, 3001, "11:01"},   // hour and minutes roll over continuously
// {72, 8640, "00:00"},    // hour and minutes roll over to exactly midnight
// {-1, 15, "23:15"},      // negative hour
// {-25, 0, "23:00"},      // negative hour rolls over
// {-91, 0, "05:00"},      // negative hour rolls over continuously
// {1, -40, "00:20"},      // negative minutes
// {1, -160, "22:20"},     // negative minutes roll over
// {1, -4820, "16:40"},    // negative minutes roll over continuously
// {2, -60, "01:00"},      // negative sixty minutes is previous hour
// {-25, -160, "20:20"},   // negative hour and minutes both roll over
// {-121, -5810, "22:10"}, // negative hour and minutes both roll over continuously

// New returns a new Clock struct
func New(hour, minute int) Clock {
	hour += minute / 60

	hour = hour % 24
	if hour < 0 {
		hour += 24
	}

	minute = minute % 60
	if minute < 0 {
		hour--
		minute += 60
	}

	c := Clock{
		hour,
		minute,
	}

	return c
}

// String returns a human readable string
// satifies the "Stringer" interface
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hour, c.Minute)
}

// Add adds minutes
func (c Clock) Add(minutes int) Clock {

	return c
}

// Subtract subtracts minutes
func (c Clock) Subtract(minutes int) Clock {
	c.Hour -= (c.Minute + minutes) / 60
	c.Minute -= (c.Minute + minutes) % 60
	return c
}
