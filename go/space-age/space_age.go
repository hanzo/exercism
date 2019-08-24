// Space package contains functions for calculating ones age on various planets.
package space

import "fmt"

type Planet string

const SecondsInEarthYear = 31557600

var RatioToEarthYear = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age converts the given number of seconds to the number of years on the given planet.
func Age(seconds float64, planet Planet) float64 {
	earthYears := seconds / SecondsInEarthYear
	ratio, knownPlanet := RatioToEarthYear[planet]
	if !knownPlanet {
		fmt.Println(fmt.Sprintf("Unknown planet %s", planet))
		return -1
	} else {
		return earthYears / ratio
	}
}
