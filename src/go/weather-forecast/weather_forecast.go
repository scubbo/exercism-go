// Package weather contains a function Forecast for generating a weather forecase.
package weather

// CurrentCondition stores the latest condition passed to function Forecast.
var CurrentCondition string

// CurrentLocation stores the latest city passed to function Forecast.
var CurrentLocation string

/*
Forecast receives a city and a condition, stores those values in package variables, then returns
a string describing the Forecast for that city.
*/
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
