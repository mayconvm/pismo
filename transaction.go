package main

import (
	"errors"
	"log"
	"sort"
	"time"
)

// map operations type
var operationType = []struct {
	operationTypeID int
	description     string
	chargeOrder     int
}{
	1: {operationTypeID: 1, description: "COMPRA A VISTA", chargeOrder: 2},
	2: {operationTypeID: 2, description: "COMPRA PARCELADA", chargeOrder: 1},
	3: {operationTypeID: 3, description: "SAQUE", chargeOrder: 0},
	4: {operationTypeID: 4, description: "PAGAMENTO", chargeOrder: 0},
}

type Transactions struct {
	TransactionID   int
	AccountID       int8    `json:"AccountID"`
	OperationTypeID int8    `json:"OperationTypeID"`
	Amount          float32 `json:"Amount"`
	Balance         float32
	EventDate       time.Time
	DueDate         time.Time
}

// sort transaction
type SortTransactions map[int]Transactions

func (d SortTransactions) Len() int      { return len(d) }
func (d SortTransactions) Swap(i, j int) { d[i], d[j] = d[j], d[i] }
func (d SortTransactions) Less(i, j int) bool {
	if d[i].OperationTypeID == d[j].OperationTypeID {
		return d[i].EventDate.Before(d[j].EventDate)
	}

	return operationType[d[i].OperationTypeID].chargeOrder < operationType[d[j].OperationTypeID].chargeOrder
}

// to sort transaction
// sort by "operationType.chargeOrder"
// if equal "operationType.chargeOrder" sort by "transaction.EventDate"
func sortTrasanctions(transactions map[int]Transactions) map[int]Transactions {
	var mapSort SortTransactions = transactions

	// log.Println("before sort", mapSort)
	sort.Sort(mapSort)
	// log.Println("after sort", mapSort)

	return mapSort
}

// register new transaction
func registreTransaction(transactionAdapter TransactionAdapter, account Accounts) (Transactions, error) {
	var newTransaction Transactions

	// check limits user
	if checkLimitsUser(account, transactionAdapter) == false {
		return newTransaction, errors.New("User without limit")
	}

	key := len(instanceBank.transaction[transactionAdapter.AccountID]) + 1
	_, ok := instanceBank.transaction[transactionAdapter.AccountID][1]

	if !ok {
		instanceBank.transaction[transactionAdapter.AccountID] = make(map[int]Transactions)
	}

	amount := transactionAdapter.Amount
	if transactionAdapter.OperationTypeID != OperationTypeIDPayment && amount > 0 {
		amount *= -1
	}

	newTransaction = Transactions{
		TransactionID:   key,
		AccountID:       account.AccountID,
		Amount:          amount,
		Balance:         amount,
		OperationTypeID: transactionAdapter.OperationTypeID,
		EventDate:       time.Now(),
		DueDate:         time.Now(),
	}

	instanceBank.transaction[newTransaction.AccountID][key] = newTransaction

	return newTransaction, nil
}

// update data transaction
func updateTrasaction(transactions map[int]Transactions, account Accounts) {
	for _, transaction := range transactions {
		instanceBank.transaction[account.AccountID][transaction.TransactionID] = transaction
	}
}

// retriever all transaction if Balance < 0 to each account
func retrieverOpenTransactions(date string, account Accounts) map[int]Transactions {
	result := make(map[int]Transactions)
	for key, transaction := range instanceBank.transaction[account.AccountID] {
		if transaction.Balance < 0 {
			result[key] = transaction
		}
	}

	return result
}

// calculation transaction after payments
func calculationTransactions(date string, account Accounts) error {
	openTrasanction := retrieverOpenTransactions(date, account)
	transactionPayments := retrieverPaymentsTrasaction(account)

	openTrasanction = sortTrasanctions(openTrasanction)

	resultTransaction := make(map[int]map[int]Transactions)
	resultPayments := make(map[int]Transactions)

	for key, payment := range transactionPayments {
		resultTransaction[key], resultPayments[key] = resolveBalanceTransaction(openTrasanction, payment)

		updateTrasaction(resultTransaction[key], account)
	}

	updateTrasaction(resultPayments, account)

	log.Println("openTrasanction:", resultTransaction)
	log.Println("payments:", resultPayments)

	return nil
}

// resolve amounts transaction by payments
func resolveBalanceTransaction(transactions map[int]Transactions, payment Transactions) (map[int]Transactions, Transactions) {
	result := make(map[int]Transactions)

	for _, transaction := range transactions {
		// log.Println("Payment start-->", payment.Balance)
		// log.Println("Transaction start-->", transaction.Balance)

		if payment.Balance != 0 {
			if payment.Balance >= (transaction.Balance * -1) {
				payment.Balance += transaction.Balance
				transaction.Balance = 0
			} else {
				transaction.Balance += payment.Balance
				payment.Balance = 0
			}
		}

		result[transaction.TransactionID] = transaction
		// log.Println("Payment end-->", payment.Balance)
		// log.Println("Transaction end-->", transaction.Balance)
	}

	return result, payment
}

// register new payment
func registrePayments(transactionAdapter TransactionAdapter, account Accounts) (Transactions, error) {
	transactionAdapter.OperationTypeID = OperationTypeIDPayment

	// register new payment
	newTransaction, err := registreTransaction(transactionAdapter, account)

	if err != nil {
		return newTransaction, err
	}

	// update limits to user
	_, err = updateLimitsAccount(account, newTransaction)

	return newTransaction, err
}

// retriever all payments if Balance > 0 to each account
func retrieverPaymentsTrasaction(account Accounts) map[int]Transactions {
	result := make(map[int]Transactions)
	for key, transaction := range instanceBank.transaction[account.AccountID] {
		if transaction.OperationTypeID == OperationTypeIDPayment {
			result[key] = transaction
		}
	}

	return result
}
