package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// constants with transaction operation
const (
	OperationTypeIDWithdrawal = 3

	OperationTypeIDCredit = 2

	OperationTypeIDPayment = 4
)

// helper to read buffer json
type Reader interface {
	Read(p []byte) (n int, err error)
}

// insntance back to save transactions
var instanceBank = struct {
	accounts    map[int]Accounts
	transaction map[int8]map[int]Transactions
}{accounts: make(map[int]Accounts), transaction: make(map[int8]map[int]Transactions)}

// stat method
func main() {
	r := mux.NewRouter()

	// routes to accounts
	a := r.PathPrefix("/accounts").Subrouter()
	a.HandleFunc("/", httpAccounts)
	a.HandleFunc("", httpAccounts)
	a.HandleFunc("/{id}", httpAccounts)

	// routes to transaction
	t := r.PathPrefix("/transactions").Subrouter()
	t.HandleFunc("/", httpTransactions)
	t.HandleFunc("", httpTransactions)
	t.HandleFunc("/{id}", httpTransactions)

	// routes to payments
	p := r.PathPrefix("/payments").Subrouter()
	p.HandleFunc("/", httpPayments)
	p.HandleFunc("", httpPayments)

	http.Handle("/", r)
	http.ListenAndServe(":8090", nil)
}
