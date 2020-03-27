package main

import (
	"fmt"
)

func main() {
	// Emulate a while loop using for
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// Full for loop example
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// multiple element loop
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}

	// infinite loop and breaking out of loo[]
	for {
		fmt.Println("loop")
		break
	}

	// Using for loop to iterate over slices
	sl := []int{4, 5, 6}
	for k, v := range sl {
		fmt.Println(k, v)
	}

	//using for loop to iterate over strings
	s := "Hello"
	for k, v := range s {
		fmt.Println(k, string(v))
	}

	// Using for loop to iterate over Maps
	statePopulations := map[string]int{
		"CA": 1000000,
		"SC": 50000,
		"WA": 20000,
	}
	for k, v := range statePopulations {
		fmt.Println(k, v)
	}
	// You can also print just the values
	for _, v := range statePopulations {
		fmt.Println(v)
	}
	// Print just the keys
	for k := range statePopulations {
		fmt.Println(k)
	}
}
