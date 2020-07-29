package main

import (
	"fmt"
	"time"
)

func getChars(s string) {
	for _, c := range s {
		fmt.Printf("%c ", c)
	}
}

func getDigits(s []int) {
	for _, d := range s {
		fmt.Printf("%d ", d)
	}
}

func main() {
	fmt.Println("main execution started")

	// getChars goroutine
	go getChars("Hello")
	// getDigits goroutine
	go getDigits([]int{1, 2, 3, 4, 5})

	// schedule another goroutine
	time.Sleep(time.Millisecond)

	fmt.Println("\nmain execution stopped")
}
