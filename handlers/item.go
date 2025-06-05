// handlers/item.go
package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"RBAC/config"

	"github.com/gorilla/mux"
)

type Item struct {
	ID      int     `json:"id"`
	Product string  `json:"product"`
	Price   float64 `json:"price"`
	Stock   int     `json:"stock"`
	ISBN    string  `json:"isbn,omitempty"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type SuccessResponse1 struct {
	Data interface{} `json:"data,omitempty"`
}

func GetAllItems(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT Items_GroceryID, Items_Grocery_product, Items_Grocery_price, Items_Groceryc_quantity FROM Items_Grocery")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Failed to fetch items"})
		return
	}
	defer rows.Close()

	var items []Item

	for rows.Next() {
		var id int
		var product string
		var price float64
		var qty int
		rows.Scan(&id, &product, &price, &qty)
		items = append(items, Item{
			ID:      id,
			Product: product,
			Price:   price,
			Stock:   qty,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(SuccessResponse1{
		Data: items,
	})
}

func GetItemByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, `{"error": "Invalid item ID"}`, http.StatusBadRequest)
		return
	}

	var product string
	var price float64
	var qty int

	row := config.DB.QueryRow("SELECT Items_Grocery_product, Items_Grocery_price, Items_Groceryc_quantity FROM Items_Grocery WHERE Items_GroceryID = ?", id)
	err = row.Scan(&product, &price, &qty)

	if err == sql.ErrNoRows {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Item not found"})
		return
	} else if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Database error"})
		return
	}

	item := Item{
		ID:      id,
		Product: product,
		Price:   price,
		Stock:   qty,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(SuccessResponse1{
		Data: item,
	})
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var product, isbn string
	var price float64
	var qty int

	if err := r.ParseForm(); err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Error parsing form"})
		return
	}

	product = r.FormValue("product")
	isbn = r.FormValue("isbn")
	price, _ = strconv.ParseFloat(r.FormValue("price"), 64)
	qty, _ = strconv.Atoi(r.FormValue("quantity"))

	query := "INSERT INTO Items_Grocery (Items_Grocery_product, Items_Grocery_price, Items_Grocery_isbn, Items_Groceryc_quantity, create_date) VALUES (?, ?, ?, ?, NOW())"
	result, err := config.DB.Exec(query, product, price, isbn, qty)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Failed to create item"})
		return
	}

	id, _ := result.LastInsertId()
	response := SuccessResponse{
		Message: "Item created successfully",
		Data: map[string]interface{}{
			"id": id,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Invalid item ID"})
		return
	}

	var product, isbn string
	var price float64
	var qty int

	if err := r.ParseForm(); err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Error parsing form"})
		return
	}

	product = r.FormValue("product")
	isbn = r.FormValue("isbn")
	price, _ = strconv.ParseFloat(r.FormValue("price"), 64)
	qty, _ = strconv.Atoi(r.FormValue("quantity"))

	query := "UPDATE Items_Grocery SET Items_Grocery_product = ?, Items_Grocery_price = ?, Items_Grocery_isbn = ?, Items_Groceryc_quantity = ?, last_update = NOW() WHERE Items_GroceryID = ?"
	res, err := config.DB.Exec(query, product, price, isbn, qty, id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Failed to update item"})
		return
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Item not found or no changes made"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(SuccessResponse{
		Message: "Item updated successfully",
	})
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Invalid item ID"})
		return
	}

	res, err := config.DB.Exec("DELETE FROM Items_Grocery WHERE Items_GroceryID = ?", id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Failed to delete item"})
		return
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Item not found or already deleted"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(SuccessResponse{
		Message: "Item deleted successfully",
	})
}
