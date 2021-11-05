package main

import (
	"fmt"
	"unicode/utf8"
)

/**
Stings are always defined using characters or bytes. In GoLang, strings are always made of bytes.
Go uses UTF-8 encoding, so any valid character can be represented in Unicode code points.

What is GoLang rune?
A character is defined using “code points” in Unicode. Go language introduced a new term for this code point called rune.

Go rune is also an alias of type int32 because Go uses UTF-8 encoding. Some interesting points about rune and strings.


s := 'a'

    s_rune := rune(s)

    fmt.Println(s_rune)

	output
	97
**/
func main() {
	// check one
	city := "Kraków"
	fmt.Println(utf8.RuneCountInString(city))

	// check two
	const goodMorning = "शुभ प्रभात"
	for index, runeValue := range goodMorning {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}

	//cehck three
	s := "Learn go, నా పేరు శ్రీని"

	s_rune := []rune(s)
	s_byte := []byte(s)

	fmt.Println(" rune :", s_rune) // [71 214]
	fmt.Println(" byte :", s_byte) // [71 195 150]

	// see difference in output when character is Ascii vs unicde char
}
