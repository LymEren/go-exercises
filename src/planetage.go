// We calculate age on different Planets

package main

import "fmt"

type PlanetType string // I defined Planet data type (string based)

// Earth: orbital period 1.0 Earth years, 365.25 Earth days, or 31557600 seconds

const Earth = 31557600.0

var AllPlanets = map[PlanetType]float64{
	"Earth":   Earth,
	"Mercury": Earth * 0.2408467,
	"Venus":   Earth * 0.61519726,
	"Mars":    Earth * 1.8808158,
	"Jupiter": Earth * 11.862615,
	"Saturn":  Earth * 29.447498,
	"Uranus":  Earth * 84.016846,
	"Neptune": Earth * 164.79132}

// Both input and output must be seconds
func Age(seconds float64, planetname PlanetType) float64 {
	return seconds / AllPlanets[planetname]
}

func main() {
	// The result will be displayed in seconds type
	fmt.Println("You are in this age: ", Age(11500, "Venus"))
}
