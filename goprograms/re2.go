package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := "If I am 22 years 10 months and 14 days old as of july291,2020 then my DOB would be 1997-11-29"

	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)

	fmt.Printf("Pattern: %v\n", re.String()) // print pattern

	fmt.Println(re.MatchString(str1)) // true

	submatchall := re.FindAllString(str1, -1)
	for _, element := range submatchall {
		fmt.Println(element)
	}
}
