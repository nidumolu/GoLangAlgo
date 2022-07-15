package main

// importing fmt and linkedlist packages
import (
	"container/list"
	"fmt"
)

// main method
func main() {
	var intList list.List
	intList.PushBack(11)
	intList.PushBack(23)
	intList.PushBack(34)

	printListElements(&intList)

	intList.Remove(intList.Front())

	printListElements(&intList)

}

func printListElements(intList *list.List) {
	fmt.Printf("*****----------------*****")
	for element := intList.Front(); element != nil; element = element.Next() {
		fmt.Println("\n Element :", element.Value.(int))
	}
}
