// Package internal provides internal functions and variables.
package internal

import (
	"fmt"
	"os"
)

var version = "0.0.0-dev"

// CheckVersionFlag checks if len(os.Args) is 2, and the second
// arg is "--version". If so, it prints the current version and exits.
func CheckVersionFlag() {
	if len(os.Args) == 2 && os.Args[1] == "--version" {
		fmt.Println(version)
		os.Exit(0)
	}
}
