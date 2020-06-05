// Package space has an Age function
package space

// Planet represents the name of one of the 8 planets
// in our solar system.
type Planet string

// secondsPerDay is the total number of seconds in
// one Earth day
const secondsPerDay = 86400

// PlanetToEarthYears stores how "earth days" a given planet
// takes to go around the Sun one time.
// A map is used instead of an if/else or switch, for
// performance.
var PlanetToEarthYears = map[Planet]float64{
	"Mercury": 87.97,
	"Venus":   224.7,
	"Earth":   365.25,
	"Mars":    687,
	"Jupiter": 4332.59,
	"Saturn":  10759,
	"Uranus":  30688.5,
	"Neptune": 60182,
}

// Age returns how many years old a person is on a planet
// param age is a persons age in seconds
// param planet is the name of the planet to convert to
func Age(age float64, planet Planet) float64 {
	// Identify how many seconds in a given planets year.
	planetSecondsPerYear :=
		secondsPerDay * PlanetToEarthYears[planet]
	// Divide the age by a planets seconds per year to see
	// how many planet years old they are.
	return age / planetSecondsPerYear
}
