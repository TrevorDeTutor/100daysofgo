package main

import "fmt"

func main() {
	third := 1.0 / 3


	fmt.Println(third)
	fmt.Printf("%v\n", third)
	fmt.Printf("%f\n", third)
	fmt.Printf("%.3f\n", third)
	fmt.Printf("%4.2f\n", third)
	fmt.Println("")

	// Zero padding
	fmt.Printf("%05.2f\n", third)
}

// The %f verb formats the value of third with a width and with precision
// 1. Precision specifies how many digits should appear after the decimal point
// 2. Width specifies the minimum number of characters to display, including the
//	  decimal point and the digits before and after the decimal