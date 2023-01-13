package routes

import (
	"ibox/handlers"
	"ibox/pkg/middleware"
	"ibox/pkg/mysql"
	"ibox/repositories"

	"github.com/gorilla/mux"
)

func Transaction(r *mux.Router) {
	transactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepository)

	r.HandleFunc("/transactions", middleware.Auth(h.FindTransactions)).Methods("GET")
	r.HandleFunc("/incomes", middleware.Auth(h.FindIncomes)).Methods("GET")
	r.HandleFunc("/transaction", middleware.Auth(h.CreateTransaction)).Methods("POST")
	// Notification from midtrans route here ...
	r.HandleFunc("/notification", h.Notification).Methods("POST")
}
