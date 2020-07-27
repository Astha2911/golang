//buffering and itterating over channal
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func foo(c chan int, someVal int) {
	defer wg.Done()
	c <- someVal * 7
}

func main() {
	fooVal := make(chan int, 10)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go foo(fooVal, i)
	}
	wg.Wait()
	close(fooVal)
	for item := range fooVal {
		fmt.Println(item)
	}

}
