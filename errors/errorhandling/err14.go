package main

import "fmt"

// simple user unauthorized error
type UnauthorizedError struct {
	UserId int
	error
}

func validateUser(userId int) error {
	err := fmt.Errorf("Session invalid for user id %d", userId)
	return &UnauthorizedError{userId, err}
}

func main() {

	// validate user with id 1
	err := validateUser(10)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("User is allowed to perform this action!")
	}

}
