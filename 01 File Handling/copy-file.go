package main

import (
	"fmt"
	"io"
	"os"
)

// CopyFile copies the contents from one file to another
func CopyFile(from, to string) {
	// Open the source file in read-only mode
	f1, err := os.Open(from)
	if err != nil {
		panic(err)
	}
	defer f1.Close()

	// Create the destination file
	f2, err := os.Create(to)
	if err != nil {
		panic(err)
	}
	defer f2.Close()

	// Copy the contents from the source file to the destination
	n, err := io.Copy(f2, f1)
	if err != nil {
		panic(err)
	}

	fmt.Println("Bytes written :", n)
}

func main() {
	CopyFile("1.txt", "1 copy.txt")
}
