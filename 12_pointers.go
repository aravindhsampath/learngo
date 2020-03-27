package main

import (
	"fmt"
)

type myStruct struct {
	foo int
}

func main() {
	a := 42
	fmt.Println(a)
	b := a
	fmt.Println(b)  // value of a was copied into b
	var c *int = &a // c is a pointer to integer which carries address of a
	fmt.Println(c)  // O/P = 0xc0000a0010
	fmt.Println(*c) // c is now dereferenced. O/P = 42
	a = 20
	fmt.Println(a, *c) // 20 20
	*c = 25            // dereferenced pointer can also be assigned the value directly.
	fmt.Println(a, *c) // 25 25
	arr := [3]int{1, 2, 3}
	d := &arr[0]
	e := &arr[1]
	fmt.Printf("%v, %p, %p \n", arr, d, e) // [1 2 3], 0xc000018320, 0xc000018328
	// address locations are 8 bytes apart. Because Go used 8 bytes for those integers.
	// Pointer arithmetic is not allowed in Go.
	// Zero or initial value for pointer is nil
	var ms *myStruct
	fmt.Println(ms) // <nil>
	ms = new(myStruct)
	fmt.Println(ms) // &{0}
	(*ms).foo = 42
	fmt.Println((*ms).foo) // prints 42

	// pointer derefence can be taken for granted with Go runtime
	ms.foo = 42
	fmt.Println(ms.foo)

	// Remember that arrays are passed by value(copied), and slices are passed by reference
	// Maps are also passed by reference

}
