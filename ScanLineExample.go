package main

import (
	"fmt"
	"time"
)

func callMe(animal string) {

	for i := 0; true; i++ {
		fmt.Println("Counting %s %d", animal, i)
		time.Sleep(time.Millisecond * 1000)
	}
}
func main() {
	go callMe(" Sheep ")
	go callMe(" Goat ")

	// Scanln is similar to Scan, but stops scanning at a newline and after the final item there must be a newline or EOF.
	fmt.Scanln()
}
