package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {

	// original error
	orignalError := errors.New("I am original error message.")

	err := errors.Wrap(orignalError, "Context Info using Wrap")
	fmt.Printf("normal => %v \n\n", err)

	// print with stack trace
	fmt.Printf("with stack trace => %+v \n\n", err)

	// get original error using Cause function
	orignalErrorUnWrapped := errors.Cause(err)
	fmt.Println(orignalErrorUnWrapped)

}
