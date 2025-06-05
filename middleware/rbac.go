package middleware

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"RBAC/models"
	"RBAC/utils"

	"github.com/gorilla/mux"
)

// RolePermissions dengan level CRUD
var RolePermissions = map[string][]string{
	"Admin": {
		"create_customer", "read_customer", "update_customer", "delete_customer",
		"create_items", "read_items", "update_items", "delete_items",
		"create_transaction", "read_transaction", "update_transaction", "delete_transaction",
		"create_employee", "read_employee", "update_employee", "delete_employee",
		"read_stock", "update_stock",
		"read_report", "generate_report",
		"read_finance", "update_finance",
	},
	"Manager": {
		"create_customer", "read_customer", "update_customer", "delete_customer",
		"create_items", "read_items", "update_items", "delete_items",
		"create_transaction", "read_transaction", "update_transaction", "delete_transaction",
		"create_employee", "read_employee", "update_employee", "delete_employee",
		"read_stock", "update_stock",
		"read_report", "generate_report",
		"read_finance", "update_finance",
	},
	"Cashier": {
		"read_customer",
		"read_items",
		"create_transaction", "read_transaction",
	},
	"Stock Keeper": {
		"read_items", "update_items",
		"read_stock", "update_stock",
		"read_report",
	},
	"Security": {
		"read_customer",
	},
	"Cleaner": {},
	"Accountant": {
		"read_transaction", "update_transaction",
		"read_report", "generate_report",
		"read_finance", "update_finance",
	},
	"Marketing": {
		"read_items", "update_items",
		"read_report", "generate_report",
	},
}

func HasPermission(role, requiredPerm string) bool {
	for _, perm := range RolePermissions[role] {
		if strings.ToLower(perm) == strings.ToLower(requiredPerm) {
			return true
		}
	}
	return false
}

func RBACMiddleware(db *sql.DB, requiredPerm string) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userID := r.URL.Query().Get("user_id")
			if userID == "" {
				http.Error(w, "Missing user_id", http.StatusUnauthorized)
				return
			}

			emp, err := models.GetEmployeeByID(db, utils.StringToInt(userID)) // ✅ Gunakan dengan prefix
			if err != nil {
				http.Error(w, "User not found", http.StatusUnauthorized)
				return
			}

			if !HasPermission(emp.Role, requiredPerm) {
				http.Error(w, fmt.Sprintf("Forbidden: You don't have '%s' access", requiredPerm), http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
