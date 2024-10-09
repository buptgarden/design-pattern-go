package bridge

import "fmt"

type HP struct {
}

func (h *HP) PrintFile() {
	fmt.Println("Printing by a HP printer.")
}
