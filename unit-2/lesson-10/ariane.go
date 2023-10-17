package main

import (
	"fmt"
)

func main()  {
	var bh float64 = 32767
	var h = int16(bh)
	fmt.Println(h)

	// convert number to string
	countdown := 9
	str := fmt.Sprintf("Launch in T minus %v seconds.", countdown)
	fmt.Println(str)
}