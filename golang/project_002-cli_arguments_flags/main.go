package main

import (
	"flag"
	"fmt"
)

func main() {
	// variable for string flag contents
	wordPtr := flag.String("word", "foo", "a string")
	// variable for int flag contents
	numbPtr := flag.Int("numb", 42, "an int")
	// variable for boolean flag contents
	boolPtr := flag.Bool("bool", false, "a bool")
	// variable for string flag contents using previously defined variable
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// Execute command line parsing of configured flag options
	flag.Parse()

	// Print out provided string output
	// Need to deference the pointers (*wordPtr) to get any trailing values
	fmt.Printf("String word: %s\n", *wordPtr)
	fmt.Printf("Number value: %d\n", *numbPtr)
	fmt.Printf("Bool value: %t\n", *boolPtr)
	fmt.Printf("String var: %s\n", svar)
	fmt.Println("Trailing values:", flag.Args())
}
