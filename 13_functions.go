package main

import (
	"fmt"
)

func main() {
	fmt.Printf("hello, functions!\n")
	// Func to show multiple arguments
	sayMessage("Hello fn", 5)
	// Func to show multiple args of same type
	sayMessages("Hello", "message1", "message2")

	// Func to show pass by value
	account := 100
	deposit := 50
	doAdd(account, deposit)
	fmt.Println("At main after doAdd :", account, deposit) // 100 50

	// Func uses pointers to pass by reference
	doAddPtr(&account, &deposit)
	fmt.Println("At main after doAddPtr :", account, deposit) // 150 50

	// Passing pointers yield a great perf benefit for large datastructures.
	// passing slices, and maps are already by reference.

	// Variadic functions
	msg := "The sum is :"
	sum(msg, 1, 2, 3, 4, 5)
	//Function with a return value
	res := getSum(msg, 1, 2, 3, 4, 5) // res is a ptr
	fmt.Println(*res)

	// Func to show error handling
	d, err := divide(5.0, 0.1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)

	// Anonymous function
	func() {
		fmt.Println("Hello from instant anonymous fn")
	}()

	// Anonymous function as a type
	f := func() {
		fmt.Println("Hello from anonymous fn as a type")
	}
	// Execute the function wherever like..
	f()
}

func sayMessage(msg string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(msg)
	}
}

// Multiple argumets of same type can be written simpler.
func sayMessages(msg1, msg2, msg3 string) {
	fmt.Println(msg1, msg2, msg3)
}

// Func to demonstrate pass by value
func doAdd(account, deposit int) {
	account += deposit
	fmt.Println("Within doAdd :", account, deposit) // 150 50
}

// Func to demonstrate pass by ref using pointers
func doAddPtr(account, deposit *int) {
	*account += *deposit
	fmt.Println("Within doAdd :", *account, *deposit) // 150 50
}

// Variadic function
func sum(msg string, values ...int) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println(msg, result)
}

// Fn with a return value
func getSum(msg string, values ...int) *int {
	fmt.Println(values)
	result := 0 // local stack variable
	for _, v := range values {
		result += v
	}
	return &result // result variable is promoted to the heap and ptr is valid.
}

// Fn to show error handling
func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by Zero")
	}
	return a / b, nil
}
