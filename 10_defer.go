package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	secret := "ON"
	fmt.Printf("Start\n")
	defer fmt.Println("Middle first", secret)
	defer fmt.Println("Middle last", secret)
	secret = "OFF"
	fmt.Println("End", secret)
	// O/P will be
	// Start
	// End OFF
	// Middle last ON
	// Middle first ON
	//  Because deferred fns are executed before exiting the calling fn.

	// Order of execution of deferred fns is LIFO

	// Fn arguments are populated at calling time.

	// Defer functions are used to clean up resources/connections.
	res, err := http.Get("https://aravindh.net")
	if err != nil {
		log.Fatal(err)
	}
	// Call clean up right alongside, but it execs at end of main()
	defer res.Body.Close()
	text, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(text))

}
