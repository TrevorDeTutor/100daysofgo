/*
- Reuse the kelvinToCelsius function to convert 233° K to Celsius.
- Write and use a celsiusToFahrenheit temperature conversion
function. Hint: the formula for converting to Fahrenheit is: (c * 9.0 / 5.0) +32.0.
- Write a kelvinToFahrenheit function and verify that it converts 0° K
to approximately –459.67° F.
*/
package main

import (
	"fmt"
)

func kelvinToCelsius(k float64) float64  {
	k -= 273.15
	return k
}

func celsiusToFahrenheit(c float64) float64 {
	return  (c * 9.0/5.0) + 32.0
}

func kelvinToFahrenheit (k float64) float64 {
	return celsiusToFahrenheit(kelvinToCelsius(k))
}

func main() {
	fmt.Printf("233° K is %.2f° C\n", kelvinToCelsius(233))
	fmt.Printf("0° K is %.2f° F\n", kelvinToFahrenheit(0))
}