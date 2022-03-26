// Package weather provides the tools to forecast current weather conditions of your current location.
package weather

// CurrentCondition represent current condtion of your current location.
var CurrentCondition string
// CurrentLocation represent your current location.
var CurrentLocation string

//Forecast takes city and condtion as arguments, with the help of these arguments it gets to know your current location and current location condtions and returns current weather condtitions of your current location. 
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
