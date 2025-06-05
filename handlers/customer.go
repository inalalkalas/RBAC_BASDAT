package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"RBAC/config"

	"github.com/gorilla/mux"
)

// CreateCustomer handles POST /customers
func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var name, phone string
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	name = r.FormValue("customer_name")
	phone = r.FormValue("customer_phone")

	query := "INSERT INTO customer (customer_name, customer_phone, create_date) VALUES (?, ?, NOW())"
	result, err := config.DB.Exec(query, name, phone)
	if err != nil {
		http.Error(w, "Failed to create customer", http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	fmt.Fprintf(w, "Customer created with ID: %d\n", id)
}

// GetCustomers handles GET /customers
func GetCustomers(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT CustomerID, customer_name, customer_phone FROM customer WHERE delete_date IS NULL")
	if err != nil {
		http.Error(w, "Failed to fetch customers", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name, phone string
		rows.Scan(&id, &name, &phone)
		fmt.Fprintf(w, "ID: %d | Name: %s | Phone: %s\n", id, name, phone)
	}
}

// GetCustomerByID handles GET /customers/{id}
func GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, _ := strconv.Atoi(idStr)

	var name, phone string
	err := config.DB.QueryRow("SELECT customer_name, customer_phone FROM customer WHERE CustomerID = ?", id).Scan(&name, &phone)
	if err != nil {
		http.Error(w, "Customer not found", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Customer: %s | Phone: %s\n", name, phone)
}

// UpdateCustomer handles PUT /customers/{id}
func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, _ := strconv.Atoi(idStr)

	var name, phone string
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	name = r.FormValue("customer_name")
	phone = r.FormValue("customer_phone")

	query := "UPDATE customer SET customer_name = ?, customer_phone = ? WHERE CustomerID = ?"
	_, err := config.DB.Exec(query, name, phone, id)
	if err != nil {
		http.Error(w, "Failed to update customer", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Customer updated successfully\n")
}

// DeleteCustomer handles DELETE /customers/{id}
func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, _ := strconv.Atoi(idStr)

	// Soft delete
	_, err := config.DB.Exec("UPDATE customer SET delete_date = NOW() WHERE CustomerID = ?", id)
	if err != nil {
		http.Error(w, "Failed to delete customer", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Customer deleted successfully\n")
}
