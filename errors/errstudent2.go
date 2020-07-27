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
		return &StudentError{s.Name, ErrTooOld}
	}

	// if student with a given name is already stored
	// error indicated duplicate entry
	if _, ok := db[s.Name]; ok {
		return &StudentError{s.Name, ErrDuplicateEntry}
	}

	// store the student
	db[s.Name] = s
	return nil
}

func errorHandler(err error) string {
	switch cerr := err.(type) {
	case *StudentError:
		return cerr.Name + " is " + errorHandler(cerr.Err)
	default:
		switch err {
		case ErrTooOld:
			return "too old"
		case ErrDuplicateEntry:
			return "already in DB"
		default:
			return err.Error()
		}
	}
}

func main() {
	db = make(map[string]Student)

	tom := Student{"Tom", 31}
	err := AddStudent(tom)

	fmt.Println(errorHandler(err))

	jerry := Student{"jerry", 20}
	AddStudent(jerry)

	err = AddStudent(jerry)

}
