package main

import "fmt"

// user session state interface
type UserSessionState interface {
	error
	isLoggedIn() bool
	getSessionId() int
}

// simple user unauthorized error
type UnauthorizedError struct {
	UserId    int
	SessionId int
}

// check if user is logged in
func (err *UnauthorizedError) isLoggedIn() bool {
	return err.SessionId != 0 // SessionId is 0 for non-logged in users
}

// get user session id
func (err *UnauthorizedError) getSessionId() int {
	return err.SessionId
}

// return error message
func (httpErr *UnauthorizedError) Error() string {
	return fmt.Sprintf("User with id %v unauthorized to perform this action", httpErr.UserId)
}

// mock function call to validate user, returns error
func validateUser(userId int) error {

	// return an error
	return &UnauthorizedError{userId, 1234}
}

func main() {

	// validate user with id `1`
	err := validateUser(1)

	if err != nil {

		// print error
		fmt.Println(err)

		// check if error is of `UserSessionState` interface type
		if errVal, ok := err.(UserSessionState); ok {
			if errVal.isLoggedIn() {
				fmt.Printf("Cleaning user session with session id %v", errVal.getSessionId())

				// to get nested struct
				// errValNested, ok := err.(*UnauthorizedError)
			}
		}
	} else {
		fmt.Println("User is allowed to perform this action!")
	}
}
