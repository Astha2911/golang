//make is use for slices
package main

import "fmt"

func main() {
	s := make([]int, 3)
	s[0] = 1
	s[1] = 2
	s[2] = 3
	fmt.Println(s)
	fmt.Println(len(s))
	//append function
	s = append(s, 4, 5, 6, 7)
	fmt.Println(s)
	//slice syntax
	fmt.Println(s[1:3])
	fmt.Println(s[:3])
	fmt.Println(s[1:])
	//x := s
	//fmt.Println(x)
	//x[0] = 500
	//fmt.Println(x)
	//fmt.Println(s) //if we are changing x 0 index than changes occue also in s
	//use copy to prevent from changing both x and s
	x := make([]int, len(s))
	copy(x, s)
	x[0] = 500
	fmt.Println(x)
	fmt.Println(s)

}
