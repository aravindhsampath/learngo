package main

import (
	"fmt"
)

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 { // initialise and test
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
	// Use a Map for test data
	statePopulations := map[string]int{
		"CA": 1000000,
		"SC": 50000,
		"WA": 20000,
	}
	if pop, ok := statePopulations["CA"]; ok {
		fmt.Println(pop) // pop variable is available inside the block
	}

	if 2 < 10 && 5 < 10 {
		fmt.Println("AND statement works")
	}

	if 2 < 10 || 11 < 10 {
		fmt.Println("OR statement works")
	}
}
