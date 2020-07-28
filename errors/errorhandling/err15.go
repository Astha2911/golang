package main

import "fmt"

type UnauthorizedError struct {
	UserId    int
	SessionId int
	error
}

// check if user is logged in
func (err *UnauthorizedError) isLoggedIn() bool {
	return err.SessionId != 0 // SessionId is 0 for non-logged in users session id is userid
}

// mock function call to validate user, returns error
func validateUser(userId int) error {

	// mock general error from a function call
	err := fmt.Errorf("Session invalid for user id %d", userId)

	// return UnauthorizedError with original error
	return &UnauthorizedError{userId, 1234, err}
}
func main() {

	// validate user with id `1`
	err := validateUser(1)

	if err != nil {

		// extract error object from `err` interface
		if errVal, ok := err.(*UnauthorizedError); ok {
			fmt.Println("Is user logged in:", errVal.isLoggedIn())
		} else {
			fmt.Println(err)
		}

	} else {
		fmt.Println("User is allowed to perform this action!")
	}
}
