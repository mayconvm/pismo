package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Accounts struct {
	AccountID                int8    `json:"AccountID"`
	AvailableCreditLimit     float32 `json:"AvailableCreditLimit"`
	AvailableWithdrawalLimit float32 `json:"AvailableWithdrawalLimit"`
}

func (account Accounts) String() string {
	result := "AccountID: " + strconv.Itoa(int(account.AccountID))
	result += " | AvailableCreditLimit: " + fmt.Sprintf("%f", account.AvailableCreditLimit)
	result += " | AvailableWithdrawalLimit: " + fmt.Sprintf("%f", account.AvailableWithdrawalLimit)

	return result
}

// register new account
func registreAccount(account Accounts) (Accounts, error) {
	id := len(instanceBank.accounts) + 1
	account.AccountID = int8(id)
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

	account.AvailableCreditLimit = account.AvailableCreditLimit + availableCreditLimit
	account.AvailableWithdrawalLimit = account.AvailableWithdrawalLimit + availableWithdrawalLimit

	instanceBank.accounts[id] = account

	return account, err
}

// update limits to user
func updateLimitsAccount(account Accounts, transaction Transactions) (Accounts, error) {
	availableCreditLimit := float32(0)
	availableWithdrawalLimit := float32(0)

	switch transaction.OperationTypeID {
	case OperationTypeIDWithdrawal:
		availableWithdrawalLimit = transaction.Amount
	case OperationTypeIDCredit:
		availableCreditLimit = transaction.Amount
	}

	return updateAccount(int(account.AccountID), availableCreditLimit, availableWithdrawalLimit)
}

// check limits to transactions
func checkLimitsUser(accounts Accounts, transactionAdapter TransactionAdapter) bool {
	switch transactionAdapter.OperationTypeID {
	case OperationTypeIDWithdrawal:
		return checkLimitsUserWithdrawal(accounts, transactionAdapter)
	case OperationTypeIDCredit:
		return checkLimitsUserCredit(accounts, transactionAdapter)
	default:
		return true
	}
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
