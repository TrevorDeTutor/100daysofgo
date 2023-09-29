// The rules for determining a leap year are as follows:
// - Any year that is evenly divisible by 4 but not evenly divisible by 100
// - Or any year that is evenly divisible by 400

package main

import "fmt"

func main() {
	var year = 2009  //eg. 2008 is a leap year

	//if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		fmt.Println(year, " is a leap year")
	} else {
		fmt.Println(year, "is not a leap year")
	}
}