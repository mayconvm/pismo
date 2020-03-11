package main

import (
	"encoding/json"
	"net/http"
)

type MapTransactionPaymentAdapter []TransactionAdapter

func (transaction MapTransactionPaymentAdapter) parseJson(body Reader) (MapTransactionPaymentAdapter, error) {
	var mapTransactionPaymentAdapter MapTransactionPaymentAdapter
	err := json.NewDecoder(body).Decode(&mapTransactionPaymentAdapter)

	return mapTransactionPaymentAdapter, err
}

// switch to action with payments
func httpPayments(w http.ResponseWriter, req *http.Request) {
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

	printTableBankTransaction()
}

// register new payments
func httpPaymentsRegistre(req *http.Request) (map[int]Transactions, error) {
	var mapTransactionPaymentAdapter MapTransactionPaymentAdapter
	resultMapTransaction := make(map[int]Transactions)

	mapTransactionPaymentAdapter, err := mapTransactionPaymentAdapter.parseJson(req.Body)
	if err != nil {
		return resultMapTransaction, err
	}

	for key, transactionAdapter := range mapTransactionPaymentAdapter {
		account, err := retrieverAccount(int(transactionAdapter.AccountID))
		if err != nil {
			return resultMapTransaction, err
		}

		resultMapTransaction[key], err = registrePayments(transactionAdapter, account)

		//TODO what is open transaction?
		calculationTransactions("data", account)
	}

	// log.Println("InstanceBank.transaction", instanceBank.transaction)

	return resultMapTransaction, err
}
