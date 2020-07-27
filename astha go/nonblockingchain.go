package main

import "fmt"

func main() {
	message := make(chan string)
	signals := make(chan bool)
	select {
	case msg := <-message:
		fmt.Println("recived message", msg)
	default:
		fmt.Println("no message recived")
	}
	msg := "hi"
	select {
	case message <- msg:
		fmt.Println("sent msz", msg)
	default:
		fmt.Println("no msg sent")
	}
	select {
	case msg := <-message:
		fmt.Println("recived message", msg)
	case sig := <-signals:
		fmt.Println("recived signals", sig)
	default:
		fmt.Println("no activity")
	}

}
