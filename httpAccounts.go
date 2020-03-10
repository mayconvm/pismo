package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func httpAccounts(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Accounts", req.URL)

	switch req.Method {
	case "GET":
		account, err := httpAccountsRetriever(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result, err := json.Marshal(account)
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)

	case "POST":
		account, err := httpAccountsRegistre(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result, err := json.Marshal(account)
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)

	case "PATCH":
		account, err := httpAccountsUpdate(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result, err := json.Marshal(account)
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
	}

	// log.Println(instanceBank)
}

func httpAccountsRegistre(req *http.Request) (Accounts, error) {
	var newAccount Accounts
	newAccount, err := newAccount.parseJson(req.Body)

	if err != nil {
		return newAccount, err
	}

	//instanceBank[newAccount.AccountID] = newAccount
	resultAccount, err := registreAccount(newAccount)

	return resultAccount, err
}

func httpAccountsRetriever(req *http.Request) (Accounts, error) {
	var account Accounts
	params := mux.Vars(req)
	// log.Println(params)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		return account, err
	}

	account, err = retrieverAccount(id)
	return account, err
}

func httpAccountsUpdate(req *http.Request) (Accounts, error) {
	var account Accounts
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		return account, err
	}
	dataRequest, err := account.parseJson(req.Body)
	account, err = updateAccount(id, dataRequest.AvailableCreditLimit, dataRequest.AvailableWithdrawalLimit)

	return account, err
}
