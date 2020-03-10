package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type TransactionAdapter struct {
	AccountID       int8    `json:"AccountID"`
	OperationTypeID int8    `json:"OperationTypeID"`
	Amount          float32 `json:"Amount"`
}

func (transaction TransactionAdapter) parseJson(body Reader) (TransactionAdapter, error) {
	var newInstance TransactionAdapter
	err := json.NewDecoder(body).Decode(&newInstance)

	return newInstance, err
}

func httpTransactions(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Transactions", req.URL)

	switch req.Method {
	case "POST":
		data, err := httpTransactionRegistre(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result, err := json.Marshal(data)
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)

	}
}

func httpTransactionRegistre(req *http.Request) (Transactions, error) {
	var newInstance Transactions
	var transactionAdapter TransactionAdapter
	transactionAdapter, err := transactionAdapter.parseJson(req.Body)

	if err != nil {
		return newInstance, err
	}

	account, err := retrieverAccount(int(transactionAdapter.AccountID))
	if err != nil {
		return newInstance, err
	}

	resultAccount, err := registreTransaction(transactionAdapter, account)

	log.Println("InstanceBank.transaction", instanceBank.transaction)

	return resultAccount, err
}
