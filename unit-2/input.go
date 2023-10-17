/*
Write a program that converts strings to Booleans:

The strings “true”, “yes”, or “1” are true.
The strings “false”, “no”, or “0” are false.
Display an error message for any other values.
*/

package main

import (
	"fmt"
)

func main()  {
	yesNo := "1"

	var launch bool

	switch yesNo {
	case "true", "yes", "1":
		launch = true
	case "false", "no", "0":
		launch = false
	default:
		fmt.Println(yesNo, "is not valid")
	}

	fmt.Println("Ready for lauch:", launch)
}