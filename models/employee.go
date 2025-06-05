package models

import "database/sql"

type Employee struct {
	ID    int
	Name  string
	Email string
	Role  string
}

func GetEmployeeByID(db *sql.DB, id int) (Employee, error) {
	var emp Employee
	err := db.QueryRow("SELECT EmployeeID, employee_name, employee_email, employee_role FROM Employee WHERE EmployeeID = ?", id).
		Scan(&emp.ID, &emp.Name, &emp.Email, &emp.Role)
	return emp, err
}
