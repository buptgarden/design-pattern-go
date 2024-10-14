package state

import (
	"fmt"
)

type ItemRequestedState struct {
	VendingMachine *VendingMachine
}

func (i *ItemRequestedState) RequestItem() error {
	return fmt.Errorf("Item Already requested")
}

func (i *ItemRequestedState) AddItem(count int) error {
	return fmt.Errorf("Item Dispense in progress")
}

func (i *ItemRequestedState) InsertMoney(money int) error {
	if money < i.VendingMachine.itemPrice {
		return fmt.Errorf("Inserted money is less, Please insert %d", i.VendingMachine.itemPrice)
	}

	fmt.Println("Money entered is ok")
	i.VendingMachine.SetState(i.VendingMachine.hasMoney)
	return nil
}

func (i *ItemRequestedState) DispeneseItem() error {
	return fmt.Errorf("Please insert money first")
}
