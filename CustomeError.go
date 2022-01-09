package main

import (
	"errors"
	"fmt"
)

func divvyUp(arg float64) (float64, error) {
	if arg == 0 {

		// `errors.New` constructs a basic `error` value
		// with the given error message.
		return -1, errors.New(`can’t divide by Zero `)

	}

	// error value of nil means no error
	return float64(100 / arg), nil
}

func main() {
	for _, i := range []int{1, 7, 0} {
		if r, e := divvyUp(float64(i)); e != nil {
			fmt.Println(`divvyUp failed:`, e)
		} else {
			fmt.Println(`divvyUp worked:`, r)
		}
	}
}
