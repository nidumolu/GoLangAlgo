package main

import (
	"fmt"
	"time"
)

func strlen(s string, c chan int) {
	fmt.Println("String :", s)
	time.Sleep(1 * time.Second)
	c <- len(s)
}

func main() {
	c := make(chan int)
	go strlen("Dictionary", c)
	go strlen("Onomatopoeia", c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
