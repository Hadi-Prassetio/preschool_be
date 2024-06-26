package routes

import (
	"preschool/handlers"
	"preschool/pkg/middleware"
	"preschool/pkg/mysql"
	"preschool/repositories"

	"github.com/gorilla/mux"
)

func AdminRoutes(r *mux.Router) {
	adminRepository := repositories.RepositoryAdmin(mysql.DB)
	h := handlers.HandlerAdmin(adminRepository)

	r.HandleFunc("/register", h.Register).Methods("POST")
	r.HandleFunc("/login", h.Login).Methods("POST")
	r.HandleFunc("/check-auth", middleware.Auth(h.CheckAdmin)).Methods("GET")
}
