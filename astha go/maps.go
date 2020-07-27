//map is lyk dictionary in python,hashtable in java
//specify key value pair
package main

import "fmt"

func main() {
	m := make(map[string]int) //string is key and int is value
	m["a"] = 0
	m["b"] = 1
	fmt.Println(m)
	fmt.Println(m["a"])
	//remove key/pair from map by using delete keyword
	delete(m, "a")
	fmt.Println(m)
	//another way to initialize map
	m2 := map[string]int{"val 1": 100, "val 2": 200}
	fmt.Println(m2)
}
