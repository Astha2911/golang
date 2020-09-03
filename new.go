package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "I'm girl, and I'm 25."
	str = strings.Replace(str, "'", "\\'", -1)
	//str = strings.Replace(str, "'", `\'`, -1)
	//str = strings.Replace(str, "'", "E'\'", -1)
	fmt.Println(str)
}
