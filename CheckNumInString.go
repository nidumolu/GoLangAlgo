package main

import (
	"fmt"
	"regexp"
)

func main() {

	fmt.Printf("Test...")
	str1 := "Hello, This is Tom and Tom is 6 feet tall and is 26 years old."

	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)

	fmt.Printf("Pattern: %v\n", re.String()) // Print Pattern

	fmt.Printf("String contains any 'number' match: %v\n", re.MatchString(str1)) // True

	/**
	func (re *Regexp) FindAllStringIndex(s string, n int) [][]int
	FindAllStringIndex is the 'All' version of FindStringIndex; it returns a slice of all successive matches of the expression
	, as defined by the 'All' description in the package comment. A return value of nil indicates no match.
		**/
	submatchall := re.FindAllString(str1, -1)
	for index, element := range submatchall {
		fmt.Println("Match found, #", index+1, " ", element)
	}
}
