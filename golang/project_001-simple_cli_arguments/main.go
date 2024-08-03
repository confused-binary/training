package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	argsWithProgram := os.Args
	argsWithoutProgram := os.Args[1:]

	// Print all arguments and the program
	fmt.Printf("All Arguments with Program: %+v\n", strings.Join(argsWithProgram, ", "))

	// Print all arguments without the program with a for loop
	fmt.Printf("All Arguments without Program: ")
	for _, value := range argsWithoutProgram {
		fmt.Printf("%v, ", value)
	}
	fmt.Printf("\n")

	fmt.Printf("Third Argument: %v\n", os.Args[3])
}
