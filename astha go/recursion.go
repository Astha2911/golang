package main

import "fmt"

//n!=n*n-1 * n-2 * n-3.....*1
func fact(n int) int {
	if n <= 1 {
		return 1
	}
	return n * fact(n-1)
}
func main() {
	fmt.Println(fact(5))
	fmt.Println(fact(10))
}
