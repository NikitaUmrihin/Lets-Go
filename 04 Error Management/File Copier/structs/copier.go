package structs

import (
	"errors"
	"filecopier/errorhandling"
	"fmt"
	"io"
	"os"
	"strings"
)

// FileCopier manages the file copying process with error tracking.
type FileCopier struct {
	From *os.File // Source file
	To   *os.File // Destination file
	// err is private so.....
	err *FileCopierError // Tracks any error encountered
}

// NewFileCopier initializes a FileCopier instance by opening the source file.
func NewFileCopier(fileName string) *FileCopier {
	// Open the source file, and handle any error that occurs
	f, err := os.Open(fileName)
	if err != nil {
		errorhandling.HandleOpenFileError(err)
		return &FileCopier{
			From: nil,
			To:   nil,
			err:  NewFileCopierError("NewFileCopier() -> os.Open()", err),
		}
	}

	return &FileCopier{
		From: f,
		To:   nil,
		err:  nil,
	}
}

// ------------------------------------------------------------------------------------
// 	WRAPPER FUNCTIONS
// ------------------------------------------------------------------------------------

// CreateNewFile creates a new destination file for copying.
func (fc *FileCopier) CreateNewFile() {

	// If there's a previous error, don't do anything
	if fc.err != nil {
		return
	}

	// Create the name for the copied file
	newFile := strings.TrimSuffix(fc.From.Name(), ".txt") + "-copy.txt"

	// Create the destination file and gandle error
	f2, err := os.Create(newFile)
	if err != nil {
		fc.err = NewFileCopierError("CreateNewFile() -> os.Create()", err)
		errorhandling.HandleCreateFileError(err)
	}
	// Assign the newly created file to the 'To' field
	fc.To = f2
}

// Copy copies data from the source file to a new destination file.
func (fc *FileCopier) Copy() int {

	// If there's a previous error, don't do anything
	if fc.err != nil {
		return -1
	}

	// Copy data and handle error
	n, err := io.Copy(fc.To, fc.From)
	if err != nil {
		fc.err = NewFileCopierError("Copy() -> io.Copy()", err)
		errorhandling.HandleCopyError(err)
	}

	fmt.Println("File copied successfully.\nBytes written:", n)
	return int(n)
}

// Close ensures both files are closed properly.
func (fc *FileCopier) Close() {
	// If there's a previous error, don't do anything
	if fc.err != nil {
		return
	}
	// Close the source file
	if err := fc.From.Close(); err != nil {
		fc.err = NewFileCopierError("Close()", err)
		errorhandling.HandleCloseFileError(err)
	}
	// Close the destination file
	if err := fc.To.Close(); err != nil {
		fc.err = NewFileCopierError("Close()", err)
		errorhandling.HandleCloseFileError(err)
	}
}

// Err returns any error encountered during file operations.
func (fc *FileCopier) Err() error {
	return fc.err
}

// ------------------------------------------------------------------------------------

// CopyFile is the main function that initializes and executes the file copy operation.
func (fc *FileCopier) CopyFile() {

	// Create the destination file for copying
	fc.CreateNewFile()
	fc.Copy()
	defer fc.Close()

	var fcErr *FileCopierError

	//
	if fc.err != nil && errors.As(fc.err, &fcErr) {
		fmt.Printf("encountered a %T in CopyFile() while: %s\n", fcErr, fcErr.Op)
	}
}
