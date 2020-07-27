//errors are just a value , conferming to an interface
package main

import (
	"errors"
	"fmt"
)

type StudentError struct {
	Name string
	Err  error
}

func (se *StudentError) Error() string {
	return se.Name + se.Err.Error()
}

var (
	ErrTooOld         = errors.New("too old")
	ErrDuplicateEntry = errors.New("duplicate entry")
)

type Student struct {
	Name string
	Age  int
}

var db map[string]Student

func AddStudent(s Student) error {
	if s.Age > 30 {
		return errors.New("student too old")
		//return &StudentError{s.Name, ErrTooOld}
	}

	// if student with a given name is already stored
	// error indicated duplicate entry
	if _, ok := db[s.Name]; ok {
		return errors.New("duplicate entry")
		//return &StudentError{s.Name, ErrDuplicateEntry}
	}

	// store the student
	db[s.Name] = s
	return nil
}

func main() {
	db = make(map[string]Student)

	tom := Student{"Tom", 31}
	err := AddStudent(tom)
	fmt.Println("we got an error:" + err.Error())

	jerry := Student{"jerry", 20}
	err = AddStudent(jerry)
	if err == nil {
		fmt.Println("adding jerry-should got an empty error")

	}
	err = AddStudent(jerry)
	fmt.Println("adding jerry twise an error" + err.Error())

}
