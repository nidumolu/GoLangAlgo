/*
Go supports anonymous functions, which can form closures. Anonymous functions are useful when you want to define a function inline without having to name it.
*/

package main

import "fmt"

/*This function intSeq returns another function, which we define anonymously in the body of intSeq. The returned function closes over the variable i to form a closure.
 */

func intSeq(k int) func() (int, int) {
	i := 0 // below func has closure on this variable
	k++    // no closure on this variable
	return func() (int, int) {
		i++
		return i, k
	}
}

/*
We call intSeq, assigning the result (a function) to nextInt. This function value captures its own i value, which will be updated each time we call nextInt.
*/

func main() {
	j := 2
	k := 3
	nextInt := intSeq(j)
	/* See the effect of the closure by calling nextInt a few times.*/

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	/*To confirm that the state is unique to that particular function, create and test a new one.
	 */
	newInts := intSeq(k)
	fmt.Println(newInts())
}
