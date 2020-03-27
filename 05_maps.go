package main

import (
	"fmt"
)

func main() {
	// Maps are the equivalent of dicts in Python.
	// Store key:value pairs that can be queries at O(1)

	statePopulations := map[string]int{
		"CA": 1000000,
		"SC": 50000,
		"WA": 20000,
	}
	fmt.Println(statePopulations)

	// A slice cannot be a key in  a map
	sl := []int{4, 5, 6}
	fmt.Println(sl)
	// tstMap := map[[]int]string{}   ----> //invalid map key type []int
	tstMap := make(map[string]int) // initialise an empty map
	tstMap["CA"] = 10
	fmt.Println(tstMap)

	tstMaparr := make(map[string][]int) // map where keys are strings and val are slices
	tstMaparr["CA"] = []int{1, 2, 3}
	tstMaparr["GA"] = []int{1, 2, 3, 5, 6, 7, 8, 9}
	fmt.Println(tstMaparr)

	fmt.Println(statePopulations["CA"])
	fmt.Println(statePopulations["MA"]) // non-existing lookup returns 0!!! Beware!

	_, ok := statePopulations["MA"]
	fmt.Println(ok) // false

	// Idiomatic way to check for presence of an element.
	attended := map[string]bool{
		"Ann": true,
		"Joe": true,
	}
	person := "Ann"
	if attended[person] { // will be false if person is not in the map
		fmt.Println(person, "was at the meeting")
	}
	fmt.Println(len(attended))
	// Maps are passed by reference by default!
	newAttended := attended
	fmt.Println("attended before modification :", attended[person])
	newAttended["Ann"] = false
	fmt.Println("attended after modification :", attended[person])
	fmt.Println("newAttended after modification :", newAttended[person])

}
