// Package weather provides tools to forecast conditions in the city.
package weather

// CurrentCondition represents a current condition.
var CurrentCondition string
// CurrentLocation represents a current location.
var CurrentLocation string

// Forecast returns a string value equal to the forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
