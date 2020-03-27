package main

import (
	"fmt"
)

const (
	a = iota
	b
	c
)

// Human readable file sizes
const (
	_  = iota // ignore first value. _ is write only variable.
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

func main() {
	fmt.Printf("Value of a = %v \n", a)
	filesize := 400000000000. // bytes
	fmt.Printf("%.2f GB \n", filesize/GB)
	fmt.Printf("%.2f TB \n", filesize/TB)
}
