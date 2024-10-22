package utils

import (
	"fmt"
	"os"
)

// HandleError print error and exit
func HandleError(err error) {
	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	os.Exit(1)
}
