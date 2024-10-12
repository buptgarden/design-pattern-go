package facade

import "fmt"

type Account struct {
	name string
}

func newAccount(accoutName string) *Account {
	return &Account{
		name: accoutName,
	}
}

func (a *Account) checkAccout(accoutName string) error {
	if a.name != accoutName {
		return fmt.Errorf("Account Name is incorret")
	}

	fmt.Println("Account Verified")
	return nil
}
