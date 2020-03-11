package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type AccountsAdapterLimit struct {
	Amout float32 `json:"amount"`
}

type AccountsAdapter struct {
	AvailableCreditLimit     AccountsAdapterLimit `json:"available_credit_limit"`
	AvailableWithdrawalLimit AccountsAdapterLimit `json:"available_withdrawal_limit"`
}

func (account AccountsAdapter) parseJson(body Reader) (Accounts, error) {
	var newAccountAdapter AccountsAdapter
	err := json.NewDecoder(body).Decode(&newAccountAdapter)

	newAccount := Accounts{
		AvailableCreditLimit:     newAccountAdapter.AvailableCreditLimit.Amout,
		AvailableWithdrawalLimit: newAccountAdapter.AvailableWithdrawalLimit.Amout,
	}

	return newAccount, err
}

// switch to action with accounts
func httpAccounts(w http.ResponseWriter, req *http.Request) {
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

	printTableBankAccounts()
}

// retriever data account
func httpAccountsRetriever(req *http.Request) (Accounts, error) {
	var account Accounts
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		return account, err
	}

	account, err = retrieverAccount(id)
	return account, err
}

// update data accounts
func httpAccountsUpdate(req *http.Request) (Accounts, error) {
	var account Accounts
	var accountsAdapter AccountsAdapter

	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		return account, err
	}
	dataRequest, err := accountsAdapter.parseJson(req.Body)
	account, err = updateAccount(id, dataRequest.AvailableCreditLimit, dataRequest.AvailableWithdrawalLimit)

	return account, err
}

// register new account
func httpAccountsRegistre(req *http.Request) (Accounts, error) {
	var newAccount Accounts
	var accountsAdapter AccountsAdapter

	newAccount, err := accountsAdapter.parseJson(req.Body)

	if err != nil {
		return newAccount, err
	}

	//instanceBank[newAccount.AccountID] = newAccount
	resultAccount, err := registreAccount(newAccount)

	return resultAccount, err
}
