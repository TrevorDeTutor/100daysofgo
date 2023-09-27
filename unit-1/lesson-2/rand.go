package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var num = rand.Intn(10) + 1
	fmt.Println(num)

	num = rand.Intn(10) + 1
	fmt.Println(num)
	fmt.Println(" ")

	//Quick check 2.6
	// Write a program that generates a random distance from 56,000,000 to 401,000,000 km.

	var distance = rand.Intn(345000001) + 56000000
    fmt.Println("random distance to Mars is:", distance)
}