package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"RBAC/config"

	"github.com/gorilla/mux"
)

// UpdateStock handles PUT /stock/{id}
func UpdateStock(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID, _ := strconv.Atoi(vars["id"])
	qty, _ := strconv.Atoi(r.FormValue("quantity"))
	empID, _ := strconv.Atoi(r.FormValue("employee_id"))

	query := "UPDATE Items_Grocery SET Items_Groceryc_quantity = Items_Groceryc_quantity + ? WHERE Items_GroceryID = ?"
	_, err := config.DB.Exec(query, qty, itemID)
	if err != nil {
		http.Error(w, "Failed to update stock", http.StatusInternalServerError)
		return
	}

	// Log update stock
	config.DB.Exec("INSERT INTO Update_Stock (Update_Stock_product, Update_Stock_quantity, EmployeeID, Items_GroceryID, create_date) VALUES (?, ?, ?, ?, NOW())",
		"", qty, empID, itemID)

	fmt.Fprintf(w, "Stock updated successfully\n")
}
