package main

import "fmt"

func add_sub(a, b int) (int, int) {
	return a + b, a - b
}
func main() {
	add_res, sub_res := add_sub(10, 2)
	fmt.Println("add_res:", add_res, "sub_res:", sub_res)

}
