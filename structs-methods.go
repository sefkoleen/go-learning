package main

import (
	"fmt"
	"math/rand"
)

// STRUCTS AND METHODS
type BankAccount struct {
	Owner         string
	Balance       float64
	AccountNumber int
}

type Bank struct {
	Accounts []*BankAccount
}

func (account *BankAccount) Deposit(amount float64) error {
	if amount == 0 {
		return fmt.Errorf("Cannot deposit 0 USD")
	} else {
		account.Balance += amount
		return nil
	}
}

func (account *BankAccount) Withdraw(amount float64) error {
	if amount > account.Balance {
		return fmt.Errorf("Insufficient balance")
	} else {
		account.Balance -= amount
		return nil
	}
}

func (account BankAccount) Info() string {
	return fmt.Sprintf("Bank Account Info\n Owner: %s\n Current Balance: %f\n Accout Number: %s\n", account.Owner, account.Balance, account.AccountNumber)
}

func CreateNewBank() (*Bank, error) {
	newBank := &Bank{
		Accounts: make([]*BankAccount, 0),
	}
	return newBank, nil
}

func (b *Bank) AddBankAccount(owner string, balance float64) (*Bank, *BankAccount) {
	newBankAccount := &BankAccount{
		Owner:         owner,
		Balance:       balance,
		AccountNumber: rand.Int(),
	}
	b.Accounts = append(b.Accounts, newBankAccount)
	return b, newBankAccount
}

func (b *Bank) FindBankAccount(accountNumber int) *BankAccount {
	for _, account := range b.Accounts {
		if account.AccountNumber == accountNumber {
			return account
		}
	}
	return nil
}
