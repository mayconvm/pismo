package main

import (
	"testing"
	"time"
)

func Test_sortTrasanctions(t *testing.T) {
	transactions := make(map[int]Transactions)

	transactions[0] = Transactions{
		TransactionID:   1,
		OperationTypeID: 1,
		EventDate:       time.Now(),
	}

	transactions[1] = Transactions{
		TransactionID:   2,
		OperationTypeID: 2,
		EventDate:       time.Now(),
	}

	transactions[2] = Transactions{
		TransactionID:   3,
		OperationTypeID: 2,
		EventDate:       time.Now().Add(time.Duration(-10 * time.Second)),
	}

	transactions[3] = Transactions{
		TransactionID:   4,
		OperationTypeID: 3,
		EventDate:       time.Now(),
	}

	sortTransactions := sortTrasanctions(transactions)

	if sortTransactions[0].TransactionID != 4 {
		t.Errorf("Trasanction -> sortTrasanctions() position %v expect %v . got = %v", 0, 4, sortTransactions[1].TransactionID)
	}

	if sortTransactions[1].TransactionID != 3 {
		t.Errorf("Trasanction -> sortTrasanctions() position %v expect %v . got = %v", 1, 3, sortTransactions[1].TransactionID)
	}

	if sortTransactions[2].TransactionID != 2 {
		t.Errorf("Trasanction -> sortTrasanctions() position %v expect %v . got = %v", 2, 2, sortTransactions[1].TransactionID)
	}

	if sortTransactions[3].TransactionID != 1 {
		t.Errorf("Trasanction -> sortTrasanctions() position %v expect %v . got = %v", 3, 1, sortTransactions[1].TransactionID)
	}
}

func Test_registreTransaction(t *testing.T) {
	// func registreTransaction(transactionAdapter TransactionAdapter, account Accounts) (Transactions, error) {
	transactionAdapter := TransactionAdapter{
		AccountID:       1,
		OperationTypeID: 1,
		Amount:          10,
	}

	account := Accounts{
		AccountID:                1,
		AvailableCreditLimit:     20,
		AvailableWithdrawalLimit: 15,
	}

	// transaction
	transaction, err := registreTransaction(transactionAdapter, account)

	if err != nil {
		t.Errorf("Trasanction -> registreTransaction() error %v", err)
	}

	if transaction.Amount != transactionAdapter.Amount*-1 {
		t.Errorf("Trasanction -> registreTransaction() expect Amout  %v. got = %v", transactionAdapter.Amount*-1, transaction.Amount)
	}

	// check limits
	transactionAdapter.Amount = 30
	transaction, err = registreTransaction(transactionAdapter, account)
	// TODO
	// if err == nil {
	if err != nil {
		t.Errorf("Trasanction -> registreTransaction() error %v", err)
	}

	// payment
	transactionAdapter.OperationTypeID = OperationTypeIDPayment
	transaction, err = registreTransaction(transactionAdapter, account)

	if err != nil {
		t.Errorf("Trasanction -> registreTransaction() error %v", err)
	}

	if transaction.Amount != transactionAdapter.Amount {
		t.Errorf("Trasanction -> registreTransaction() expect Amout  %v. got = %v", transactionAdapter.Amount, transaction.Amount)
	}
}

func Test_retrieverOpenTransactions(t *testing.T) {
	// func retrieverOpenTransactions(date string, account Accounts) map[int]Transactions {
}

func Test_calculationTransactions(t *testing.T) {
	// func calculationTransactions(date string, account Accounts) error {
}

func Test_resolveBalanceTransaction(t *testing.T) {
	transactions := make(map[int]Transactions)
	var payment Transactions

	transactions[1] = Transactions{
		TransactionID:   1,
		AccountID:       1,
		OperationTypeID: 0,
		Amount:          -10,
		Balance:         -10,
	}

	transactions[2] = Transactions{
		TransactionID:   2,
		AccountID:       1,
		OperationTypeID: 1,
		Amount:          -5,
		Balance:         -5,
	}

	payment = Transactions{
		TransactionID:   1,
		AccountID:       1,
		OperationTypeID: 1,
		Amount:          10,
		Balance:         10,
	}

	got, got1 := resolveBalanceTransaction(transactions, payment)

	if got[1].Balance != 0 {
		t.Errorf("Trasanction -> resolveBalanceTransaction() got = %v, want %v", got[1].Balance, 0)
	}

	if got1.Balance != 0 {
		t.Errorf("payment -> resolveBalanceTransaction() got = %v, want %v", got1.Balance, 0)
	}

	if got[2].Balance != -5 {
		t.Errorf("Trasanction -> resolveBalanceTransaction() got = %v, want %v", got[2].Balance, -5)
	}

}

func Test_retrieverPaymentsTrasaction(t *testing.T) {
	// func retrieverPaymentsTrasaction(account Accounts) map[int]Transactions {
}
