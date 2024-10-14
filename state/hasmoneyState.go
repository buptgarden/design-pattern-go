package state

import "fmt"

type HasMoneyState struct {
	VendingMachine *VendingMachine
}

func (i *HasMoneyState) RequestItem() error {
	return fmt.Errorf("Item dispense in progress")
}

func (i *HasMoneyState) AddItem(count int) error {
	return fmt.Errorf("Item dispense in progress")
}

func (i *HasMoneyState) InsertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}

func (i *HasMoneyState) DispeneseItem() error {
	fmt.Println("Dispensing Item")
	i.VendingMachine.itemCount = i.VendingMachine.itemCount - 1
	if i.VendingMachine.itemCount == 0 {
		i.VendingMachine.SetState(i.VendingMachine.noItem)
	} else {
		i.VendingMachine.SetState(i.VendingMachine.hasItem)
	}
	return nil
}
