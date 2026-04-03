package main

import "fmt"

func main() {
	var i interface{} = "hello"

	// Assert that i holds a string
	s, ok := i.(string)
	if ok {
		fmt.Printf("Assertion successful: s = '%s'\n", s)
	} else {
		fmt.Println("Assertion to string failed.")
	}

	// Assert that i holds a float64 (this will fail)
	f, ok := i.(float64)
	if ok {
		fmt.Printf("Assertion successful: f = %f\n", f)
	} else {
		// f will be the zero value for float64, which is 0
		fmt.Printf("Assertion to float64 failed. f = %v\n", f)
	}
}
