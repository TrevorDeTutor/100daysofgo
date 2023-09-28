package main

import (
	"fmt"
)

func main() {
	// Travel distance of 56,000,000 km in 28 days
	const distanceTravelled = 56000000 //km 
	const days = 28 //days

	fmt.Println("How fast is:", distanceTravelled/(days*24), "km/h")
}