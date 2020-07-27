package main

import "fmt"

func mult(nums ...int) int {
	total := 1
	for _, num := range nums {
		//total=total * num
		total *= num
	}
	return total
}
func main() { //println is an variadic function
	fmt.Println("this", "is", "variadic", "function")
	//call mult function with as many values as necessary
	fmt.Println(mult(1, 2, 3, 4))

}
