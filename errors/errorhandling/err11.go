//this is about to khow the method od request and status
package main

import (
	"fmt"
)

type HttpError struct { //simple http error
	status int
	method string
}

func (c *HttpError) Error() string { //HttpError implement error method
	return fmt.Sprintf("something is wrong with %v request,server returned the %v status",
		c.status, c.method)

}

// mock function: sends HTTP request
func GetUserEmail(userId int) (string, error) { //func design to send http req and return the response

	// reqst failed, return an error
	return "", &HttpError{403, "GET"}
}
func main() {

	// get user email add
	if email, err := GetUserEmail(1); err != nil {
		fmt.Println(err) // print error

		// if user is unauthorized, perform session cleaning
		if errVal := err.(*HttpError); errVal.status == 403 {
			fmt.Println("Performing session cleaning...")
		}

	} else {
		// do something with the `email`
		fmt.Println("User email is", email)
	}
}

//    errVal, ok := err.(*HttpError) where ok will be false when err does not hold the value of type *HttpError. Hence, you also need to check in the if condition if ok is true.
//when we r nt sure tht err interface hold the value of type(*HttpError)
