package main

import (
	"fmt"
	"strconv"
)

// Variables can also be declared at package level. Variables declared here
//    may be shadowed inside functions.
var (
	i int     = 42
	s string  = "Hello"
	k float32 = 62.45
)

// Multiple var blocks can exist - used to group variable decls by purpose.
var (
	z int32 = 70
)

func main() {
	var i int // i of type integer.
	i = 42
	fmt.Println(i)

	j := 44 // declare and initialize at once, let
	//             compiler infer the type.
	fmt.Printf("%v , %T \n", j, j) // %v implies value, and %T implies type.

	fmt.Printf("%v,%v,%v,%v \n", i, s, k, z)
	println(" >> String conversion")
	// Use the strconv package to convert nums to string
	l := strconv.Itoa(i)
	fmt.Println("String value is ", l)

	// BOOLEAN Variables
	// -------------------
	println(" >> Boolean")
	var n bool = true
	fmt.Printf("%v , %T \n", n, n)
	p := j == 44 // Bool result of an equality test
	fmt.Printf("%v , %T \n", p, p)

	// Numerics
	println(" >> Numerics")
	a := 10
	b := 3
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b) // Modulo operator to compute the remainder
	var c int32 = 32
	// fmt.Println(a + c)  // invalid operation: a + c (mismatched types int and int32)
	fmt.Println(a + int(c)) // works fine with integer casting
	println(" >> Strings")
	fmt.Println(string(s))
	fmt.Println(string(s[2:])) // index based slicing works
	t := " World!"
	fmt.Println(string(s + t))

	by := []byte(s) // byte array - useful in data transfers
	fmt.Println(string(by))

}
