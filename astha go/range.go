package main

import "fmt"

func main() {
	str_arr := []string{"a", "b", "c", "d"}
	for i, c := range str_arr {
		fmt.Println("idx", i)
		fmt.Println("c", c)
	}
	//can also range over key/value pairs in map
	m := map[string]int{"a": 1, "b": 2}
	for k, v := range m {
		fmt.Println("key:", k, "value:", v)
	}
}
