package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

/*Suppose we wanted to create a file, write to it,
and then close when we’re done. Here’s how we could do that with defer.
*/
func main() {
	/*Immediately after getting a file object with createFile,
	  we defer the closing of that file with closeFile. This will be executed at the end of the enclosing function (main), after writeFile has finished.
	*/
	f := createFile("/tmp/deferExample.txt")
	defer closeFile(f)
	writeFile(f)
	size, _ := printFileSize(f)
	fmt.Println("File size is : " + string(size))

	size2 := fileSize(f.Name())
	fmt.Println("File2 size is : " + string(size2))

}
func createFile(p string) *os.File {
	fmt.Println("creating file...")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}
func writeFile(f *os.File) {
	fmt.Println("writing to file :" + f.Name())
	fmt.Fprintln(f, "A quick fox jumped over the lazy dog ")

}

/*It’s important to check for errors when closing a file
, even in a deferred function.*/

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func printFileSize(f *os.File) (i int64, err error) {

	fmt.Println("Dir(p):", filepath.Dir(f.Name()))
	fmt.Println("Base(p):", filepath.Base(f.Name()))

	fi, err := os.Stat(filepath.Join(filepath.Dir(f.Name()), "/", filepath.Base(f.Name())))

	fmt.Println("FullPath of file:", fi.Name())
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return 0, err
	}
	// get the size
	size := fi.Size()
	return size, nil
}

func fileSize(path string) int64 {
	fi, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}

	return fi.Size()
}
