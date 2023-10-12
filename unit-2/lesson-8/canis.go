/*
Canis Major Dwarf is the closest known galaxy to Earth at
236,000,000,000,000,000 km from our Sun (though some dispute that it is a
galaxy). Use constants to convert this distance to light years.
*/

package main

import (
	"fmt"
)

func main()  {
	const distance = 236000000000000000 //km
	const lightSpeed = 299792
	const secondsPerDay = 86400
	const daysPerYear = 365

	// days := distance/lightspeed/secondsPerDay

	days := 236000000000000000 / 299792 / 86400
	years := days / daysPerYear

	fmt.Println("Canis Major Dwarf is", years, "light years away.")
}