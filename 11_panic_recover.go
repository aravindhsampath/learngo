package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Panic - unrecoverable error or programmer error

	fmt.Println("Hello")
	defer fmt.Println("Defer 1")
	// The foll line will trigger a panic and show where it comes from.
	// panic("Something bad happened")
	defer fmt.Println("Defer 2")
	// Only defer calls before the panic are executed at panic.
	panicker()
	// Eventhough panic occured in panicker(), it was recovered.
	// Therefore execution continues as normal in main()
	// Anonymous function..
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error:", err)
			// The foll line can re-throw the panic
			// panic(err)
		}
	}()
	panic("Ooops..")

	// Simple web server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})
	// For e.g, if 9090 is already in use, the application panics
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		panic(err.Error)
	}

}

func panicker() {
	fmt.Println("About to panic")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error:", err)
		}
	}()
	panic("Panicker panicked..")
	fmt.Println("Last line in panicker") // Never gets printed
}
