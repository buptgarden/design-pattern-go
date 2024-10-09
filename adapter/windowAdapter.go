package adapter

import "fmt"

type WindowsAdaper struct {
	Windows *Window
}

func (w *WindowsAdaper) InsertIntoLightningPort() {
	fmt.Println("Inset Light port into Windows Adapter machine")
	w.Windows.InsertIntoUsbPort()
}
