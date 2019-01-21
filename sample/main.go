package main

import (
	_ "./matchers"
	"./search"
	"fmt"
	"log"
	"os"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	fmt.Println("Main init")
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {
	// Perform the search for the specified term.
	search.Run("president")
}
