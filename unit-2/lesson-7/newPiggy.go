/*
Write a new piggy bank program that uses integers to track the number of cents
rather than dollars. Randomly place nickels (5¢), dimes (10¢), and quarters (25¢)
into an empty piggy bank until it contains at least $20.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	const nickels int = 5
	const dimes int = 10
	const quarters int = 25
	piggyBank := 0
	
	for piggyBank < 2000 {
		time.Sleep(time.Second)
		
		switch rand.Intn(3) {
		case 0:
			piggyBank += nickels
			fmt.Println("Added nickels ")

		case 1:
			piggyBank += dimes
			fmt.Println("Added dimes ")
			
		case 2:
			piggyBank += quarters
			fmt.Println("Added quarters ")
		}
		
		dollars := piggyBank /100
		cents := piggyBank % 100

		fmt.Printf("Piggy balance is $%d.%02d\n", dollars, cents)
	}
}