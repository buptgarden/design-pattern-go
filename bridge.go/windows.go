package bridge

import "fmt"

type Windows struct {
	Printer Printer
}

func (w *Windows) Print() {
	fmt.Println("Print request for windows.")
	w.Printer.PrintFile()
}

func (w *Windows) SetPrinter(p Printer) {
	w.Printer = p
}
