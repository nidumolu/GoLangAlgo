// Program in GO language to demonstrates how to use base log package.
package main

import (
	"log"
)

/*init() function is just like the main function, does not take any argument nor return anything.
This function is present in every package and this function is called when the package is initialized.
This function is declared implicitly, so you cannot reference it from anywhere and you are allowed
to create multiple init() function in the same program and they execute in the order they are created. */
func init() {
	log.SetPrefix("LOGGING: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("init() started")
}

/*The main() function is a special type of function and it is the entry point of the executable programs.
It does not take any argument nor return anything*/

func main() {
	// Println writes to the standard logger.
	log.Println("main started")

	// Fatalln is Println() followed by a call to os.Exit(1)
	log.Fatalln("fatal message")

	// Panicln is Println() followed by a call to panic()
	log.Panicln("panic message")
}
