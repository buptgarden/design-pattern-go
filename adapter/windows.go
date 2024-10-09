package adapter

import "fmt"

type Window struct {
}

func (w *Window) InsertIntoUsbPort() {
	fmt.Println("USB connector is plugged into window machine")
}
