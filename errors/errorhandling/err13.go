//adding context to an error
package main

import "fmt"

// simple user unauthorized error
type UnauthorizedError struct {
	UserId        int
	OriginalError error
}

// add some context to the original error
func (a *UnauthorizedError) Error() string {
	return fmt.Sprintf("User unauthorized Error: %v", a.OriginalError)
}

// mock function call to validate user, returns error
func validateUser(userId int) error {

	// mock general error from a function call
	err := fmt.Errorf("Session invalid for user id %d", userId)

	// return UnauthorizedError with original error
	return &UnauthorizedError{userId, err}
}

func main() {

	// validate user with id `1`
	err := validateUser(1)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("User is allowed to perform this action!")
	}

}
