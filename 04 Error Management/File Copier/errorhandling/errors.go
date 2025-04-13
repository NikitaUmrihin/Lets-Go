package errorhandling

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
)

// HandleOpenFileError logs errors encountered while opening a file.
func HandleOpenFileError(err error) {
	if errors.Is(err, os.ErrPermission) {
		err = fmt.Errorf("you don't have permission to open file: %w", err)
		log.Println(err)
	} else if errors.Is(err, os.ErrNotExist) {
		err = fmt.Errorf("the file doesn't exist: %w", err)
		log.Println(err)
	} else {
		err = fmt.Errorf("file couldn't be opened: %w", err)
		log.Println(err)
	}
}

// HandleCreateFileError logs errors encountered while creating a file.
func HandleCreateFileError(err error) {
	switch {
	case errors.Is(err, os.ErrPermission):
		err = fmt.Errorf("you don't have permission to open file: %w", err)
		log.Println(err)
	case errors.Is(err, os.ErrExist):
		err = fmt.Errorf("the file already exists: %w", err)
		log.Println(err)
	default:
		err = fmt.Errorf("file couldn't be created: %w", err)
		log.Println(err)
	}
}

// HandleCopyError logs errors encountered while copying data from one file to another.
func HandleCopyError(err error) {
	if errors.Is(err, os.ErrNotExist) {
		err = fmt.Errorf("source file does not exist: %w", err)
		log.Println(err)
	} else if errors.Is(err, io.EOF) {
		err = fmt.Errorf("unexpectedly reached end of file: %w", err)
		log.Println(err)
	} else {
		err = fmt.Errorf("couldn't copy file: %w", err)
		log.Println(err)
	}
}

// HandleCloseFileError logs errors encountered while closing a file.
func HandleCloseFileError(err error) {
	if errors.Is(err, fs.ErrClosed) {
		err = fmt.Errorf("file is already closed: %w", err)
		log.Println(err)
	} else if errors.Is(err, os.ErrPermission) {
		err = fmt.Errorf("you don't have permission to close the file: %w", err)
		log.Println(err)
	} else {
		err = fmt.Errorf("an unexpected error occurred while closing the file: %w", err)
		log.Println(err)
	}
}
