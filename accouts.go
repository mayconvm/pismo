package main

import (
	"encoding/json"
	"errors"
)

type Accounts struct {
	AccountID                int8    `json:"AccountID"`
	AvailableCreditLimit     float32 `json:"AvailableCreditLimit"`
	AvailableWithdrawalLimit float32 `json:"AvailableWithdrawalLimit"`
}

// parse json
func (account Accounts) parseJson(body Reader) (Accounts, error) {
	var newAccount Accounts
	err := json.NewDecoder(body).Decode(&newAccount)

	return newAccount, err
}

// register new account
func registreAccount(account Accounts) (Accounts, error) {
	id := len(instanceBank.accounts) + 1
	instanceBank.accounts[id] = account

	return account, nil
}

// check account exists
func hasAccount(id int) bool {
	_, account := instanceBank.accounts[id]
	return account
}

// retriever account by id
func retrieverAccount(id int) (Accounts, error) {
	var account Accounts
	if hasAccount(id) {
		account = instanceBank.accounts[id]
		return account, nil
	}

	return account, errors.New("Account not found")
}

// update limits to user
func updateAccount(id int, availableCreditLimit float32, availableWithdrawalLimit float32) (Accounts, error) {
	account, err := retrieverAccount(id)
	if err != nil {
		return account, err
	}

	accountAvailableCreditLimit := account.AvailableCreditLimit - availableCreditLimit
	accountAvailableWithdrawalLimit := account.AvailableWithdrawalLimit - availableWithdrawalLimit

	if accountAvailableCreditLimit < 0 {
		accountAvailableCreditLimit = 0
	}

	if accountAvailableWithdrawalLimit < 0 {
		accountAvailableWithdrawalLimit = 0
	}

	account.AvailableCreditLimit = accountAvailableCreditLimit
	account.AvailableWithdrawalLimit = accountAvailableWithdrawalLimit

	instanceBank.accounts[id] = account

	return account, err
}

// check limits to transactions
func checkLimitsUser(accounts Accounts, transactionAdapter TransactionAdapter) bool {

	switch transactionAdapter.OperationTypeID {
	case OperationTypeIDWithdrawal:
		return checkLimitsUserWithdrawal(accounts, transactionAdapter)
	case OperationTypeIDCredit:
		return checkLimitsUserCredit(accounts, transactionAdapter)
	}

	return true
}

// check credit limit available to transaction
func checkLimitsUserCredit(accounts Accounts, transactionAdapter TransactionAdapter) bool {
	if accounts.AvailableCreditLimit >= transactionAdapter.Amount {
		return true
	}

	return false
}

// check withdrawal limit available to transaction
func checkLimitsUserWithdrawal(accounts Accounts, transactionAdapter TransactionAdapter) bool {
	if accounts.AvailableWithdrawalLimit >= transactionAdapter.Amount {
		return true
	}

	return false
}
