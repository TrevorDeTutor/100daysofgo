// My weight loss program
package main

import "fmt"

func main() {
	fmt.Println("My weight on the surface of Mars is ")
	fmt.Println(149.0 * 0.3783)
	fmt.Println(" lbs, and I would be ")
	fmt.Println(26 *365 / 687)
	fmt.Println(" years old.")
	fmt.Println(" ")

	//use Printf
	fmt.Printf("My weight on the surface of Mars is %v lbs,", 139.0*0.3783)
	fmt.Printf(" and I would be %v years old.\n", 41*365/687)
}