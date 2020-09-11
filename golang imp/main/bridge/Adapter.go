package main

import "fmt"

type computer interface {
	insertInSquarePort()
}

type mac struct {
}

func (m *mac) insertInSquarePort() {
	fmt.Println("Insert square port into mac machine")
}

type windowsAdapter struct {
	windowMachine *windows
}

func (w *windowsAdapter) insertInSquarePort() {
	w.windowMachine.insertInCirclePort()
}

type windows struct{}

func (w *windows) insertInCirclePort() {
	fmt.Println("Insert circle port into windows machine")
}

type client struct {
}

func (c *client) insertSquareUsbInComputer(com computer) {
	com.insertInSquarePort()
}
func main() {
	client := &client{}
	mac := &mac{}
	client.insertSquareUsbInComputer(mac)
	windowsMachine := &windows{}
	windowsMachineAdapter := &windowsAdapter{
		windowMachine: windowsMachine,
	}
	client.insertSquareUsbInComputer(windowsMachineAdapter)
}
