package main

import "fmt"

// create a struct
type MyError struct{}

// struct implements `Error` method
func (myErr *MyError) Error() string {
	return "Something unexpected happend!"
}

func main() {

	// create error
	myErr := &MyError{}

	// print error message
	fmt.Println(myErr)
}
