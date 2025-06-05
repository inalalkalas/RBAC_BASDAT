package handlers

import (
	"fmt"
	"net/http"

	"RBAC/config"
)

// CreateEmployee handles POST /employees
func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("employee_name")
	email := r.FormValue("employee_email")
	role := r.FormValue("employee_role")

	query := "INSERT INTO Employee (employee_name, employee_email, employee_role, create_date) VALUES (?, ?, ?, NOW())"
	result, err := config.DB.Exec(query, name, email, role)
	if err != nil {
		http.Error(w, "Failed to create employee", http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	fmt.Fprintf(w, "Employee created with ID: %d\n", id)
}
