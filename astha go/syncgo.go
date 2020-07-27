//goroutine syncronization
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func say(s string) {
	//here-defer wg.Done()
	for i := 0; i < 4; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done()

}
func main() {
	wg.Add(1)
	go say("hello")
	wg.Add(1)
	go say("how  are you")
	wg.Wait()
}

//we can also use in sync when we used wg.Done
//we can also use defer wg.Done
