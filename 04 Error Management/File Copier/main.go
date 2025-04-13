package main

import (
	"filecopier/structs"
)

// This program demonstrates robust, two-layered error handling using:

// 		- errors.Is to detect standard I/O errors
// 		- errors.As to extract and handle custom errors
// 		- A FileCopier struct that encapsulates file operations and tracks errors
// 		- Centralized, readable, and reusable error logging functions

// LETS GO
func main() {
	fc := structs.NewFileCopier("1.txt")
	fc.CopyFile()
}
