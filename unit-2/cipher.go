/*
To send ciphered messages, write a program that ciphers plain text using a keyword:
plainText := "your message goes here"
keyword := "GOLANG"
*/

package main

import (
	"fmt"
)

func main()  {
	message := "hello from golang"
	keyword := "golang"
	keyIndex := 0
	cipherText := ""

	message = strings.ToUpper(strings.Replace(message, " ", "", -1))
	keyword = strings.ToUpper(strings.Replace(keyword, " ", "", -1))
	
	for i := 0; i < len(message); i++ {
		c := message[i]

		if c >= 'A' && c <= 'Z' {
			// A=0, B=1, ... Z=25
			c -= 'A'
			k := keyword[keyIndex] - 'A'
			// cipher letter + key letter
			c = (c+k)%26 + 'A'
			// increment keyIndex
			keyIndex++
			keyIndex %= len(keyword)
			}
			cipherText += string(c)
		}

		fmt.Println(cipherText)
}
