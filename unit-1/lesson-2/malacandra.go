package main

import (
	"fmt"
)

func main() {
	// Travel distance of 56,000,000 km in 28 days
	const hoursPerDay = 24 //hrs
	var distance = 56000000 //km 
	var days = 28 //days

	fmt.Println("How fast is:", distance/(days*hoursPerDay), "km/h")
}