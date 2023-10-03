// program makes use of variables, constants, switch, if, and for. 
// it should also draw on the fmt and math/rand packages

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Printf("Spaceline      Days Trip type  Price\n")
	fmt.Println("====================================")

	for count := 0; count < 10; count++ {
		time.Sleep(time.Second)

		year := 2018 + rand.Intn(10)
		leap := year%400 == 0 || (year%4 == 0 && year%100 != 0)
		month := rand.Intn(12) + 1
		daysInMonth := 31
		//speed := 16 + rand.Intn(14) //km/s
		spaceline := rand.Intn(3) + 1
		selectedSL := ""
		trip := rand.Intn(2) + 1
		tripType := ""
		price := 30 + rand.Intn(70)

		switch spaceline {
		case 1:
			selectedSL = "Space Adventures"
		case 2:
			selectedSL = "SpaceX"
		case 3:
			selectedSL = "Virgin Galactic"
		}

		switch trip {
		case 1: 
			tripType = "Round-trip" 
		case 2:
			tripType = "One-way"
		}


		switch month {
		case 2:
			daysInMonth = 28
			if leap {
				daysInMonth = 29
			}
		case 4, 6, 9, 11:
			daysInMonth = 30
		}

		day := rand.Intn(daysInMonth) + 1

		fmt.Println(selectedSL, "        ", day, tripType, "$", price )
	}
}