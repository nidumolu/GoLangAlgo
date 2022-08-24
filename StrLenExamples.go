package main

import "fmt"

/*
A string is also an array  of bytes, but you have to be careful:
only for the ASCII subset is a  byte equal to a character.
All other characters occupy 2, 3, or 4 bytes.
This means that the length of a string in characters (runes) is  generally not the same as the length of its byte array.
 They are equal  only when the string consists of ASCII characters only.
*/

func main() {
	str1 := "శుభోదయం"

	str2 := "Good Morning"
	fmt.Println("Str1: String length: ", len([]rune(str1)))
	fmt.Println("Str1: Byte length: ", len(str1))

	fmt.Println("Str2: String length: ", len([]rune(str2)))
	fmt.Println("Str2: Byte length: ", len(str2))
}
