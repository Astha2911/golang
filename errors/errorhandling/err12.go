package main

import "fmt"

// network error
type NetworkError struct {
}

func (e *NetworkError) Error() string {
	return "A network connection was aborted."
}

// file save failed error
type FileSaveFailedError struct {
}

func (e *FileSaveFailedError) Error() string {
	return "The requested file could not be saved"
}

// a function that can return either of the above errors
func saveFileToRemote() error {
	//result := 1 // mock result of save oprtn.

	if result == 1 {
		return &NetworkError{}
	} else if result == 2 {
		return &FileSaveFailedError{}
	} else {
		return nil
	}
}

func main() {
	switch err := saveFileToRemote(); err.(type) {
	case nil:
		fmt.Println("File successfully saved.")
	case *NetworkError:
		fmt.Println("Network Error:", err)
	case *FileSaveFailedError:
		fmt.Println("File save Error:", err)
	}

}
