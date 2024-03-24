package main

import (
	"net/http"

	"github.com/byClebert/go-finance/cmd/api-server/handler"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("POST /v1/expenses", handler.CreateExpenseHandle)
	mux.HandleFunc("PUT /v1/expenses/{id}", handler.UpdateExpenseHandle)
	mux.HandleFunc("DELETE /v1/expenses/{id}", handler.DeleteExpenseHandle)
	mux.HandleFunc("GET /v1/expenses", handler.ListExpensesHandle)
	mux.HandleFunc("GET /v1/expenses/{id}", handler.GetExpenseHandle)

	http.ListenAndServe(":3000", mux)
}
