// Write a guess-the-number program. 
//Make the computer pick random numbers between 1–100 until it guesses your number, 
//which you declare at the top of the program. 
//Display each guess and whether it was too big or too small.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var mynum = 39
	
	for {
		var randomnum = rand.Intn(100) + 1

		fmt.Println("Random number is: ", randomnum)
		time.Sleep(time.Second)
		if randomnum == mynum {
			fmt.Println("Got the number")
			break
		} else {
			fmt.Println("Let me try again!!")
		}
	}	
	
}