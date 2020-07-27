package main

import "fmt"

func main() {
	i := 2
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	j := 5
	switch j {
	case 1, 2:
		fmt.Println("one or two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("none")
	}
	switch {
	case j == 5:
		fmt.Println("j equal to 5")
	case j > 5:
		fmt.Println("j is greater than 5")
	}
}
