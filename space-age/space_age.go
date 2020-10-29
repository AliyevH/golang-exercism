package space

import (
	"math"
)

// Planet type is a custom string
type Planet string

var earthYearInSeconds float64 = 31557600.0

var planetYears = map[string]float64{
	"Earth":   1.0,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age calculates the age relative to a planet by a given seconds and planet
func Age(seconds float64, planet Planet) float64 {
	// Finding year in a given planet
	planetYear := seconds / earthYearInSeconds / planetYears[string(planet)]
	// Return float rounding year
	return math.Round(planetYear*100) / 100
}
