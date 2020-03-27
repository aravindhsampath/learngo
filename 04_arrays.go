package main

import (
	"fmt"
)

func main() {
	grades := [3]int{1, 2, 3}
	fmt.Printf("Grades : %v\n", grades)
	var students [3]string                               // 3 element empty array
	students[0] = "Aravindh"                             // Address by index
	fmt.Printf("There are %v students\n", len(students)) // 3 students
	// 2D Array
	var i [3][3]int
	i[0] = [3]int{1, 0, 0}
	i[1] = [3]int{0, 1, 0}
	i[2] = [3]int{0, 0, 1}
	fmt.Println(i)

	// VERY IMPORTANT :
	//     Arrays are passed by value by default and NOT by reference.

	agrades := grades // pass by value
	agrades[1] = 10
	fmt.Println(grades)
	fmt.Println(agrades)

	// To explicitly pass by reference
	bgrades := &grades // pass by reference
	bgrades[1] = 10
	fmt.Println(grades)
	fmt.Println(bgrades)

	// Slices are like arrays but dynamic in nature. Copying happens in bg.
	sl := []int{4, 5, 6}
	fmt.Printf("Length = %v and Capacity = %v\n", len(sl), cap(sl))
	sl1 := sl[:] // sl1 is a slice of all elements of sl
	fmt.Printf("Length = %v and Capacity = %v\n", len(sl1), cap(sl1))
	sl2 := sl[1:] // sl1 is a slice of elements from 1 onwards in sl
	fmt.Printf("Length = %v and Capacity = %v\n", len(sl2), cap(sl2))
	// VERY IMPORTANT :
	//     Slices are passed by reference by default and NOT by value.
	sl2[1] = 10
	fmt.Println(sl)
	fmt.Println(sl2)

	// Expanding a slice means the runtime copies the contents of the array over
	//     to a new bigger location in the background - inefficient.
	// Always try to reserve a reasonable size for an array to avoid copying penalty.

	list := make([]int, 3, 100) // length =3 and Capacity =100
	list = append(list, 1)
	fmt.Println(list) // prints  [0 0 0 1] as make initialised first 3 elements with 0.
	list[0] = 1
	fmt.Println(list)

}
