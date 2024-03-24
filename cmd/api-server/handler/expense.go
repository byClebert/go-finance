package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/byClebert/go-finance/internal/model"
)

var expenses []model.Expense

func CreateExpenseHandle(w http.ResponseWriter, r *http.Request) {
	var expense model.Expense

	err := json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	expense.ID = len(expenses) + 1
	expense.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

	expenses = append(expenses, expense)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(expense)
}

func UpdateExpenseHandle(w http.ResponseWriter, r *http.Request) {
	expenseIdString := r.PathValue("id")
	expenseId, err := strconv.Atoi(expenseIdString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var expense model.Expense

	err2 := json.NewDecoder(r.Body).Decode(&expense)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}

	for i := range expenses {
		if expenses[i].ID == expenseId {
			expenses[i].Amount = expense.Amount
			expenses[i].Description = expense.Description
			expenses[i].Category = expense.Category
			expenses[i].UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
			expense = expenses[i]
			break
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(expense)
}

func DeleteExpenseHandle(w http.ResponseWriter, r *http.Request) {
	expenseIdString := r.PathValue("id")
	expenseId, err := strconv.Atoi(expenseIdString)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	for i := range expenses {
		if expenses[i].ID == expenseId {
			expenses[i].DeletedAt = time.Now().Format("2024-03-24 13:36")
			break
		}
	}

	w.WriteHeader(http.StatusAccepted)
}

func ListExpensesHandle(w http.ResponseWriter, r *http.Request) {
	var activeExpenses []model.Expense

	for i := range expenses {
		if expenses[i].DeletedAt == "" {
			activeExpenses = append(activeExpenses, expenses[i])
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(activeExpenses)
}

func GetExpenseHandle(w http.ResponseWriter, r *http.Request) {
	expenseIdString := r.PathValue("id")
	expenseId, err := strconv.Atoi(expenseIdString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var expense model.Expense

	for i := range expenses {
		if expenses[i].ID == expenseId {
			expense = expenses[i]
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(expense)
}
