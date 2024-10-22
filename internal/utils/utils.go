package utils

import (
	"fmt"
	"io"
	"os"
)

// HandleError print error and exit
func HandleError(err error) {
	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	os.Exit(1)
}

// OpenFileOrStdin open or return os.Stdin if file name is "-"
func OpenFileOrStdin(filename string) (io.ReadCloser, error) {
	if filename == "-" {
		return os.Stdin, nil
	}
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("can not open file: %v", err)
	}
	return file, nil
}

func CheckFileExists(filename string) error {
	if filename == "-" {
		return nil
	}
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return fmt.Errorf(" No such file: '%s'", filename)
	}
	if info.IsDir() {
		return fmt.Errorf(" Expected file got directory: '%s'", filename)
	}
	return nil
}
