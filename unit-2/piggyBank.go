//check 6.6
//If you add 11 dimes ($0.10 each) to an empty piggyBank of type float64, what is the final balance?


// Assignment:
// Write a program that randomly places nickels ($0.05), dimes ($0.10), and 
// quarters ($0.25) into an empty piggy bank until it contains at least $20.00.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	//check
	pb := 0.0

	for i := 0; i < 11; i++ {
		pb += 0.10
	}
	fmt.Println("Final balance is ", pb)

	fmt.Println("")
	
	// Assignment
	const nickels = 0.05
	const dimes = 0.10
	const quarters = 0.25
	piggyBank := 0.0

	for piggyBank < 20.00 {
		time.Sleep(time.Second)

		switch rand.Intn(3) {
		case 0:
			piggyBank += nickels
			fmt.Printf("Added nickels, balance is %.2f\n", piggyBank )
		case 1:
			piggyBank += dimes
			fmt.Printf("Added dimes, balance is %.2f\n", piggyBank )
		case 2:
			piggyBank += quarters
			fmt.Printf("Added quarters, balance is %.2f\n", piggyBank )
		}
	}
}