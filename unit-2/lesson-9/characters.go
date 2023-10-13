package main

import (
	"fmt"
)

func main()  {
	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33
	fmt.Printf("%c%c%c%c\n", pi, alpha, omega, bang)

	fmt.Printf("%c\n", 128515)

	fmt.Printf("%c\n\n", 65)

	//one character per line
	name := "trevor"
	for i := 0; i < 6; i++ {
		c := name[i]
		fmt.Printf("%c\n",c)
	}
}