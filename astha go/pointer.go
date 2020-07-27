package main

import "fmt"

func main() {
	val := 20
	fmt.Println(val)
	ptr := &val //it will give the add where 20 is stored
	fmt.Println(ptr)
	//print value stored at the address
	fmt.Println(*ptr)
	*ptr = 10
	fmt.Println(*ptr)
	fmt.Println(val) // it will update the value and return 10
	val = 50
	fmt.Println(*ptr) //it will update 50
}
