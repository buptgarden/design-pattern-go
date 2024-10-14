package state

import "fmt"

type NoItemState struct {
	VendingMachine *VendingMachine
}

func (i *NoItemState) RequestItem() error {
	return fmt.Errorf("Item out of stock")
}

func (i *NoItemState) AddItem(count int) error {
	i.VendingMachine.IncrementItem(count)
	i.VendingMachine.SetState(i.VendingMachine.hasItem)
	return nil
}

func (i *NoItemState) InsertMoney(monye int) error {
	return fmt.Errorf("Item out of sock")
}

func (i *NoItemState) DispeneseItem() error {
	return fmt.Errorf("Item out of stock")
}
