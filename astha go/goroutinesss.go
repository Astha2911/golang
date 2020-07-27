package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}
}
func main() {
	go say("hey") //if we use go in both statement it will print nothing bcz of cuncurentn operation
	// for using go in both statement we have to use  time.sleep()
	go say("there")
	time.Sleep(time.Second)
}
