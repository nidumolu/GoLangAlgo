package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	var b bytes.Buffer

	for i := 0; i < 10; i++ {
		b.WriteString(randString(i))
	}

	fmt.Println(b.String())
}

func randString(param int) string {
	// Pretend to return a random string
	return (strconv.Itoa(param) + " -")
}
