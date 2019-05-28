/*
 This sample code goes with chapter 2 of the book: Go in action, Manning publications
 All rights reserved.
 Code sample from: github.com/goinaction/code (chapter 2)
*/

package main

import (
	"log"
	"os"

	_ "github.com/goinaction/code/chapter2/sample/matchers"
	"github.com/goinaction/code/chapter2/sample/search"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {
	// Perform the search for the specified term.
	// default word to search
	var searchTerm string = "president"
	// has the program been called with an argument to search ?
	if len(os.Args) > 1{
		// if yes (see documentation for "os" package and Args variable)
		searchTerm = os.Args[1]
	}
	// Now we search for "president" or whatever the user gave as an argument
	search.Run(searchTerm)
}
