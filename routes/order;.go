package routes

import (
	"ibox/handlers"
	"ibox/pkg/middleware"
	"ibox/pkg/mysql"
	"ibox/repositories"

	"github.com/gorilla/mux"
)

func OrderRoutes(r *mux.Router) {
	orderRepository := repositories.RepositoryOrder(mysql.DB)
	h := handlers.HandlerOrder(orderRepository)

	r.HandleFunc("/orders", h.FindOrder).Methods("GET")
	r.HandleFunc("/order/{id}", h.GetOrder).Methods("GET")
	r.HandleFunc("/order", middleware.Auth(h.CreateOrder)).Methods("POST")
	r.HandleFunc("/order/{id}", middleware.Auth(h.DeleteOrder)).Methods("DELETE")
	r.HandleFunc("/order/{id}", middleware.Auth(h.UpdateOrder)).Methods("PATCH")
}
