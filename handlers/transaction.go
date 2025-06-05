package handlers

import (
	"RBAC/config"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateTransaction handles POST /transactions
func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	itemID, _ := strconv.Atoi(r.FormValue("item_id"))
	custID, _ := strconv.Atoi(r.FormValue("customer_id"))
	empID, _ := strconv.Atoi(r.FormValue("employee_id"))
	amount, _ := strconv.ParseFloat(r.FormValue("total_amount"), 64)

	query := "INSERT INTO transaction (Items_GroceryID, CustomerID, EmployeeID, total_amount, transaction_date) VALUES (?, ?, ?, ?, NOW())"
	result, err := config.DB.Exec(query, itemID, custID, empID, amount)
	if err != nil {
		http.Error(w, "Failed to create transaction", http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	fmt.Fprintf(w, "Transaction created with ID: %d\n", id)
}

func GetAllTrachanctions(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT TransactionID, Items_GroceryID, CustomerID, EmployeeID, total_amount, transaction_date FROM transaction")
	if err != nil {
		http.Error(w, "Failed to fetch transactions", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id, itemID, custID, empID int
		var amount float64
		var date string
		rows.Scan(&id, &itemID, &custID, &empID, &amount, &date)
		fmt.Fprintf(w, "Transaction ID: %d | Item ID: %d | Customer ID: %d | Employee ID: %d | Amount: %.2f | Date: %s\n", id, itemID, custID, empID, amount, date)
	}
}

func GetTransactionByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, _ := strconv.Atoi(idStr)

	var itemID, custID, empID int
	var amount float64
	var date string
	err := config.DB.QueryRow("SELECT Items_GroceryID, CustomerID, EmployeeID, total_amount, transaction_date FROM transaction WHERE TransactionID = ?", id).
		Scan(&itemID, &custID, &empID, &amount, &date)

	if err != nil {
		http.Error(w, "Transaction not found", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Transaction ID: %d | Item ID: %d | Customer ID: %d | Employee ID: %d | Amount: %.2f | Date: %s\n", id, itemID, custID, empID, amount, date)
}
