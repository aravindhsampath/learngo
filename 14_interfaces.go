package main

import (
	"fmt"
)

type Writer interface {
	Write([]byte) (int, error)
}
type ConsoleWriter struct {
	// Empty struct that implements the Write interface
}

// Method in the ConsoleWriter Struct
func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

// Since the ConsoleWriter struct implemented _all_ the methods in Writer interface,
//    we can claim that ConsoleWriter implements Writer interface.

// Imagine there could be a FileWriter struct and LogWriter Struct all implementing the writer interface.

func main() {
	// Reference : https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
	// Interface is
	//     * A set of methods
	//     * A type
	// Design our abstractions in terms of what actions our types can execute.
	var w Writer = ConsoleWriter{}
	// A var of type interface is asigned an empty struct that implements that interface.
	w.Write([]byte("A generic write from interface! "))
	// var w of type interface now allows access to its functions that are eventually implemented
	//                           by the struct.

	// The above is only useful in case of using interfaces with multiple implementations.
	// One can also use the struct directly using,
	wr := ConsoleWriter{}
	wr.Write([]byte("Hello directly from ConsoleWriter"))
}
