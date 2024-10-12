package facade

import "fmt"

type Ledger struct {
}

func (s *Ledger) makeEntry(accoutID, txnType string, amount int) {
	fmt.Printf("Make ledger entry for accountID %s with txnType %s for amount %d\n", accoutID, txnType, amount)
}
