package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/kataras/tablewriter"
	"github.com/landoop/tableprinter"
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
	port := ":8090"
	r := mux.NewRouter()

	// routes to accounts
	a := r.PathPrefix("/accounts").Subrouter()
	a.HandleFunc("", httpAccounts).Methods("POST")
	a.HandleFunc("/{id}", httpAccounts).Methods("PATCH")
	a.HandleFunc("/{id}/limits", httpAccounts).Methods("GET")

	// routes to transaction
	t := r.PathPrefix("/transactions").Subrouter()
	t.HandleFunc("", httpTransactions).Methods("POST")

	// routes to payments
	p := r.PathPrefix("/payments").Subrouter()
	p.HandleFunc("", httpPayments).Methods("POST")

	fmt.Println("Server: localhost" + port)

	http.Handle("/", r)
	http.ListenAndServe(port, nil)
}

func printTableBankAccounts() {
	printer := tableprinter.New(os.Stdout)

	log.Println("Accounts")

	printer.BorderTop = true
	printer.BorderBottom = true
	printer.BorderLeft = true
	printer.BorderRight = true
	printer.CenterSeparator = "│"
	printer.ColumnSeparator = "│"
	printer.RowSeparator = "─"
	printer.HeaderBgColor = tablewriter.BgBlackColor
	printer.HeaderFgColor = tablewriter.FgGreenColor

	printer.Print(instanceBank.accounts)
}

func printTableBankTransaction() {
	printer := tableprinter.New(os.Stdout)

	log.Println("Transaction")

	printer.BorderTop = true
	printer.BorderBottom = true
	printer.BorderLeft = true
	printer.BorderRight = true
	printer.CenterSeparator = "│"
	printer.ColumnSeparator = "│"
	printer.RowSeparator = "─"
	printer.HeaderBgColor = tablewriter.BgBlackColor
	printer.HeaderFgColor = tablewriter.FgGreenColor

	printer.Print(instanceBank.transaction)
}
