package main

import (
	"fmt"
	"time"
)

func main() {

	// Simple switch statement
	i := 3
	switch i {
	case 1, 10: //multiple options split by ,
		fmt.Println("one or ten")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// switch with default
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
}
