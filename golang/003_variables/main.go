package main

import "fmt"

// Types of variables (Details: https://go.dev/src/builtin/builtin.go)
// bool
// string
// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64
// byte (alias for uint8)
// rune (alias for int32, represents Unicode code point)
// float32 float64
// complex64 complex128

// Long form
var int_one int

// Long form with declaration
var string_one string = "empty"

// Constants can't use := or be a complex type (slices, maps, structs)
const bool_const = true

func main() {
	// Short form - can only be done inside a function
	// Will infer the type based on value provided
	float_one := 3.1415926535 // https://www.youtube.com/watch?v=vFfkvGhzv_4
	complex_one := 3.1 + 0.5i

	// Same line declaration (multi-declaration)
	int_three, string_two := 12345, "onetwothreefourfive"

	// Printf verbs https://pkg.go.dev/fmt
	fmt.Printf("%d", int_one)
	fmt.Printf("%s", string_one)
	fmt.Printf("%t", bool_const)
	fmt.Printf("%b", float_one)
	fmt.Printf("%v", complex_one) // %v for when not sure what to use
	fmt.Printf("%d", int_three)
	fmt.Printf("%s", string_two)
}
