//also use Errorf function from fmt package to create interpolated error messages
package main

import "fmt"

func main() {
	const name, id = "astha", 15
	err := fmt.Errorf("name %q (id %d)is not found", name, id)
	if err != nil {
		fmt.Print(err)

	}

}
