/*Slices
Slices are important in Go. A slice is basically an array, but without a fixed length. Go achieves this by doing some clever things under the hood. A ‘Slice’ is really a pointer to pointer to part of an array (hence the name ‘Slice’) and offers more flexibility than a regular array.

Here’s the syntax for creating an array and a slice. As you can see, the slice is the same as an array, but without the length value.
*/

package main

import (
	"fmt"
)

func main() {
	array := [3]bool{true, true, false}
	fmt.Println(array)

	slice := []bool{true, true, false}
	fmt.Println(slice)
}
