package routes

import (
	"ibox/handlers"
	"ibox/pkg/middleware"
	"ibox/pkg/mysql"
	"ibox/repositories"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)

	r.HandleFunc("/users", h.FindUsers).Methods("GET")
	r.HandleFunc("/user/{id}", middleware.Auth(h.GetUser)).Methods("GET")
	r.HandleFunc("/user", middleware.Auth(middleware.UploadFile(h.UpdateUser))).Methods("PATCH")
}
