package main

import (
	"fmt"
)

// A simple struct
type Doctor struct {
	number     int
	name       string
	companions []string
}

// struct that is composed of another struct
type Specialist struct {
	Doctor
	specialty string
	salary    int
}

func main() {
	// Structs are a heterogenous collection of elements
	aDoctor := Doctor{number: 3, name: "Jon", companions: []string{"Liz", "Jo"}}
	fmt.Println(aDoctor)            //{3 Jon [Liz Jo]}
	fmt.Println(aDoctor.companions) // [Liz Jo]

	// anonymous structs are temporary structures used to organise
	//                                elements within the function.
	newDoctor := struct{ name string }{name: "Aravindh"}
	fmt.Println(newDoctor)

	// Structs are passed by value!
	anotherDoc := aDoctor
	anotherDoc.name = "Tom"
	fmt.Println(aDoctor)    // {3 Jon [Liz Jo]}
	fmt.Println(anotherDoc) // {3 Tom [Liz Jo]}

	// To pass a struct by reference, use the & operator
	yetAnotherDoc := &aDoctor
	yetAnotherDoc.name = "Ron"
	fmt.Println(yetAnotherDoc) // &{3 Ron [Liz Jo]}
	fmt.Println(aDoctor)       // {3 Ron [Liz Jo]}

	// Structs can be used to achieve composition in Go
	//   In other words struct of a struct
	splDoc := Specialist{}
	splDoc.number = 1
	splDoc.specialty = "neurology"
	fmt.Println(splDoc) // {{1  []} neurology 0}

}
