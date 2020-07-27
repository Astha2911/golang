package go_errors

import (
	"fmt"
)

func main() {
	const fn = "temp.cpp"
	var ln = 19
	err := fmt.Errorf("compile problem with %q: %d", fn, ln)
	if err != nil {
		fmt.Println(err)
	}
}
