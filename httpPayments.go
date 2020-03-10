package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func httpPaymentsRegistre(req *http.Request) (Transactions, error) {
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

	resultAccount, err := registrePayments(transactionAdapter, account)

	//TODO what is open transaction?
	calculationTransactions("data", account)

	log.Println("InstanceBank.transaction", instanceBank.transaction)

	return resultAccount, err
}

func httpPayments(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Payments", req.URL)

	switch req.Method {
	case "POST":
		data, err := httpPaymentsRegistre(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result, err := json.Marshal(data)
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
	}
}
