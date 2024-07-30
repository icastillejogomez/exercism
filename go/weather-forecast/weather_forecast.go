// Package weather provides weather forecasts.
package weather

// CurrentCondition is the current condition.
var CurrentCondition string

// CurrentLocation is the current location.
var CurrentLocation string

// Forecast returns the forecast for the given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
