//Panic and Recover

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func cleanup() {
	if r := recover(); r != nil {
		fmt.Println("Recovered in cleanup", r)
	}
}
func say(s string) {
	defer wg.Done()
	defer cleanup()
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
		if i == 2 {
			panic("oh dear, a 2")
		}
	}
}
func main() {
	wg.Add(1)
	go say("hello")
	wg.Add(1)
	go say("how are you")
	wg.Wait()
}
