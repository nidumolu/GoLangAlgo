package main

import (
	"fmt"
	"math/rand"
	"time"
)

func blabber(msgs ...string) <-chan string {
	c := make(chan string)
	go func() {
		/* in Go language, defer statements delay the execution of the
		function or method or an anonymous method until the nearby functions returns. In other words, defer function or method call arguments evaluate instantly
		, but they don't execute until the nearby functions returns.
		*/
		defer close(c)
		//infinit eloop untif channel closed
		for {
			// random blabber msg
			c <- fmt.Sprintf("%s", msgs[rand.Intn(len(msgs))])

			// sleep for a while
			time.Sleep(time.Duration(rand.Intn(100) * int(time.Millisecond)))

		}
	}()
	return c
}

func main() {
	babbler := blabber("hi", "howz the weather", "how r u doing", "lol", "wtf")
	timeout := time.After(1 + time.Second)

	for {
		select {
		case s := <-babbler:
			fmt.Println(s)
			// if timeout exit
		case <-timeout:
			fmt.Print("Enough of our babbling for today, bye :-)")
			return
		}
	}
}
