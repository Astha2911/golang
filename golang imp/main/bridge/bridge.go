package main

import "fmt"

type computer interface {
	print()
	setPrinter(printer)
}

type mac struct {
	printer printer
}

func (m *mac) print() {
	fmt.Println("Print request for mac")
	m.printer.printFile()
}

func (m *mac) setPrinter(p printer) {
	m.printer = p
}

type windows struct {
	printer printer
}

func (w *windows) print() {
	fmt.Println("Print request for windows")
	w.printer.printFile()
}

func (w *windows) setPrinter(p printer) {
	w.printer = p
}

type printer interface {
	printFile()
}

type epson struct {
}

func (p *epson) printFile() {
	fmt.Println("Printing by a EPSON Printer")
}

type hp struct {
}

func (p *hp) printFile() {
	fmt.Println("Printing by a HP Printer")
}

func main() {
	hpPrinter := &hp{}
	epsonPrinter := &epson{}
	macComputer := &mac{}
	macComputer.setPrinter(hpPrinter)
	macComputer.print()
	fmt.Println()
	macComputer.setPrinter(epsonPrinter)
	macComputer.print()
	fmt.Println()
	winComputer := &windows{}
	winComputer.setPrinter(hpPrinter)
	winComputer.print()
	fmt.Println()
	winComputer.setPrinter(epsonPrinter)
	winComputer.print()
	fmt.Println()
}
