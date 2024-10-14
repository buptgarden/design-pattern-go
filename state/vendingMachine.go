package state

import "fmt"

type VendingMachine struct {
	hasItem     State
	itemRequest State
	hasMoney    State
	noItem      State

	currentState State
	itemPrice    int
	itemCount    int
}

func NewVendingMachine(itemCount, itemPrice int) *VendingMachine {
	v := &VendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}

	hasItemState := &HasItemState{
		VendingMachine: v,
	}
	itemRequestedState := &ItemRequestedState{
		VendingMachine: v,
	}
	itemMoneyState := &HasMoneyState{
		VendingMachine: v,
	}
	noItemState := &NoItemState{
		VendingMachine: v,
	}
	v.SetState(hasItemState)
	v.hasItem = hasItemState
	v.itemRequest = itemRequestedState
	v.hasMoney = itemMoneyState
	v.noItem = noItemState
	return v
}

func (v *VendingMachine) RequestItem() error {
	return v.currentState.RequestItem()
}

func (v *VendingMachine) AddItem(count int) error {
	return v.currentState.AddItem(count)
}

func (v *VendingMachine) InsertMoney(money int) error {
	return v.currentState.InsertMoney(money)
}

func (v *VendingMachine) DispeneseItem() error {
	return v.currentState.DispeneseItem()
}

func (v *VendingMachine) SetState(s State) {
	v.currentState = s
}

func (v *VendingMachine) IncrementItem(count int) {
	fmt.Printf("Adding %d items\n", count)
	v.itemCount = v.itemCount + count
}
