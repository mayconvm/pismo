package main

import (
	"testing"
)

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
