//as we know that we use func keyword
package main

import "fmt"

func add(a int, b int) int {
	return a + b
}
func sub(a int, b int) int {
	return a - b
}

//we can also use
func addn(a, b, c int) int {
	return a + b + c
}
func main() {
	a := add(1, 2)
	fmt.Println(a)
	ans := sub(20, 10)
	fmt.Println(ans)
	addz := addn(10, 20, 30)
	fmt.Println(addz)
}
