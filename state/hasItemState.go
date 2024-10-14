package state

import "fmt"

type HasItemState struct {
	VendingMachine *VendingMachine
}

func (i *HasItemState) RequestItem() error {
	if i.VendingMachine.itemCount == 0 {
		i.VendingMachine.SetState(i.VendingMachine.noItem)
		return fmt.Errorf("no item present")
	}
	fmt.Printf("Item requested\n")
	i.VendingMachine.SetState(i.VendingMachine.itemRequest)
	return nil
}

func (i *HasItemState) AddItem(count int) error {
	fmt.Printf("%d items added\n", count)
	i.VendingMachine.IncrementItem(count)
	return nil
}

func (i *HasItemState) InsertMoney(money int) error {
	return fmt.Errorf("Please select item first")
}

func (i *HasItemState) DispeneseItem() error {
	return fmt.Errorf("Please select item first")
}
