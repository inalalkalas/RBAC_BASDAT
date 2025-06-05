package main

import (
	"fmt"
	"log"
	"net/http"

	"RBAC/config"
	"RBAC/handlers"
	"RBAC/middleware"

	"github.com/gorilla/mux"
)

func main() {
	config.ConnectDB()

	r := mux.NewRouter()

	// Customer Routes
	r.Handle("/customers", middleware.RBACMiddleware(config.DB, "read_customer")(http.HandlerFunc(handlers.GetCustomers))).Methods("GET")
	r.Handle("/customers/{id}", middleware.RBACMiddleware(config.DB, "read_customer")(http.HandlerFunc(handlers.GetCustomerByID))).Methods("GET")
	r.Handle("/customers", middleware.RBACMiddleware(config.DB, "create_customer")(http.HandlerFunc(handlers.CreateCustomer))).Methods("POST")

	// Items Grocery Routes
	r.Handle("/items", middleware.RBACMiddleware(config.DB, "read_items")(http.HandlerFunc(handlers.GetAllItems))).Methods("GET")
	r.Handle("/items/{id}", middleware.RBACMiddleware(config.DB, "update_items")(http.HandlerFunc(handlers.UpdateItem))).Methods("PUT")
	r.Handle("/items/{id}", middleware.RBACMiddleware(config.DB, "read_items")(http.HandlerFunc(handlers.GetItemByID))).Methods("GET")

	// Transaction Routes
	r.Handle("/transactions", middleware.RBACMiddleware(config.DB, "create_transaction")(http.HandlerFunc(handlers.CreateTransaction))).Methods("POST")
	r.Handle("/transactions", middleware.RBACMiddleware(config.DB, "read_transaction")(http.HandlerFunc(handlers.GetAllTrachanctions))).Methods("GET")
	r.Handle("/transactions/{id}", middleware.RBACMiddleware(config.DB, "read_transaction")(http.HandlerFunc(handlers.GetTransactionByID))).Methods("GET")

	// Stock Routes
	r.Handle("/stock/{id}", middleware.RBACMiddleware(config.DB, "update_stock")(http.HandlerFunc(handlers.UpdateStock))).Methods("PUT")

	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
