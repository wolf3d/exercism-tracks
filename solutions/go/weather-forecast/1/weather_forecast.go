// Package weather provides a method to display current
// weather conditions in a particular city.
package weather

// CurrentCondition describes current weather conditions.
var CurrentCondition string
// CurrentLocation contains current location information.
var CurrentLocation string

// Forecast returns a string with info about current conditions in a current location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
