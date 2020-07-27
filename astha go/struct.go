package main

import "fmt"

type employee struct {
	first_name string
	last_name  string
	id         int
}

func main() {
	fmt.Println(employee{first_name: " astha", last_name: "mishra", id: 15})
	//also
	fmt.Println(employee{"astha", "mishra", 2})
	//also
	emp := employee{first_name: "neha", last_name: "mishra", id: 4}
	fmt.Println(emp.first_name)
	fmt.Println(emp.last_name)
	fmt.Println(emp.id)
	emp_ptr := &emp
	fmt.Println(emp_ptr.first_name)

}

//if uh will not initialize id it will take bydefault 0
