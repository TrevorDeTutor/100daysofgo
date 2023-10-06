package main

import (
	"fmt"
)

func main() { 
	year := 2018
	fmt.Printf("Type %T for %v\n", year, year)

	days := 365.2425
	fmt.Printf("Type %T for %[1]v\n", days)

	//hexadecimal
	var red, green, blue uint8 = 0, 141, 213
	//var red, green, blue uint8 = 0x00, 0x8d, 0xd5
	fmt.Printf("%x %x %x\n", red, green, blue)

	//css
	fmt.Printf("color: #%02x%02x%02x\n;", red, green, blue)
}
