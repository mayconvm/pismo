package main

import (
	"testing"
	"time"
)

func Test_updateLimitsAccount(t *testing.T) {
	// func updateLimitsAccount(account Accounts, transaction Transactions) (Accounts, error) {
	account := Accounts{
		AccountID:                1,
		AvailableCreditLimit:     100,
		AvailableWithdrawalLimit: 100,
	}

	_, err := registreAccount(account)
	if err != nil {
		t.Errorf("Trasanction -> updateLimitsAccount() error %v", err)
	}

	instanceBank.transaction[account.AccountID] = make(map[int]Transactions)
	transactions := Transactions{
		TransactionID:   1,
		OperationTypeID: 2,
		Amount:          100,
		EventDate:       time.Now(),
	}

	account, err = updateLimitsAccount(account, transactions)

	if err != nil {
		t.Errorf("Trasanction -> updateLimitsAccount() error %v", err)
	}

	if account.AvailableCreditLimit != 0 {
		t.Errorf("Trasanction -> updateLimitsAccount() Expect %v, got %v", 0, account.AvailableCreditLimit)
	}

	if account.AvailableWithdrawalLimit != 100 {
		t.Errorf("Trasanction -> updateLimitsAccount() Expect %v, got %v", 100, account.AvailableWithdrawalLimit)
	}
}
func Test_checkLimitsUser(t *testing.T) {
	// func checkLimitsUser(accounts Accounts, transactionAdapter TransactionAdapter) bool {
	account := Accounts{
		AccountID:                1,
		AvailableCreditLimit:     100,
		AvailableWithdrawalLimit: 100,
	}

	transactions := TransactionAdapter{
		OperationTypeID: 2,
		Amount:          110,
	}

	if checkLimitsUser(account, transactions) {
		t.Errorf("Trasanction -> checkLimitsUser() User not possibilty to do this operation. Limit %v, amount %v", account.AvailableCreditLimit, transactions.Amount)
	}
}
func Test_checkLimitsUserCredit(t *testing.T) {
	// func checkLimitsUserCredit(accounts Accounts, transactionAdapter TransactionAdapter) bool {
	account := Accounts{
		AccountID:                1,
		AvailableCreditLimit:     90,
		AvailableWithdrawalLimit: 100,
	}

	transactions := TransactionAdapter{
		OperationTypeID: 2,
		Amount:          90,
	}

	if !checkLimitsUserCredit(account, transactions) {
		t.Errorf("Trasanction -> checkLimitsUserCredit() User not possibilty to do this operation. Limit %v, amount %v", account.AvailableCreditLimit, transactions.Amount)
	}
}
func Test_checkLimitsUserWithdrawal(t *testing.T) {
	// func checkLimitsUserWithdrawal(accounts Accounts, transactionAdapter TransactionAdapter) bool {
	account := Accounts{
		AccountID:                1,
		AvailableWithdrawalLimit: 100,
	}

	transactions := TransactionAdapter{
		OperationTypeID: 1,
		Amount:          90,
	}

	if !checkLimitsUserWithdrawal(account, transactions) {
		t.Errorf("Trasanction -> checkLimitsUserWithdrawal() User not possibilty to do this operation. Limit %v, amount %v", account.AvailableWithdrawalLimit, transactions.Amount)
	}
}
