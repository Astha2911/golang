package main

import (
	"errors"
	"fmt"
)

func main() {

	// create error
	myErr := errors.New("Something unexpected happend!")

	// print error type
	fmt.Printf("Type of myErr is %T \n", myErr)
	fmt.Printf("Value of myErr is %#v \n", myErr)
}

//type of myErr is *errors.errorstring which is pointer to a errors.errorString
