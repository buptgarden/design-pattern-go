package chain

import "fmt"

type Cashier struct {
	next Department
}

func (c *Cashier) Execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
	}

	fmt.Println("Cashier getting money from patient")
}

func (c *Cashier) SetNext(next Department) {
	c.next = next
}
