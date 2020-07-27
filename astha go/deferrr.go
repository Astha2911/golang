package main

import "fmt"

func foo() {
	defer fmt.Println("Done!")
	defer fmt.Println("are we done!")
	fmt.Println("welcome here") // defer will print from bottom to up
	for i := 1; i < 5; i++ {
		defer fmt.Println(i)
	}
}
func main() {
	foo()
}
