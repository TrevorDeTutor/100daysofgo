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

		time.Sleep(time.Second)

		if randomnum < mynum {
			fmt.Printf("%v is too small.\n", randomnum)
			fmt.Println("Let me try again!")
		} else if randomnum > mynum {
			fmt.Printf("%v is too big.\n", randomnum)
			fmt.Println("Let me try again!")
		} else {
			fmt.Printf("Got the number! %v\n", randomnum)
			break
		}
	}	
	
}