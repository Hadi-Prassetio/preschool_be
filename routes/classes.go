package routes

import (
	"preschool/handlers"
	// "preschool/pkg/middleware"
	"preschool/pkg/mysql"
	"preschool/repositories"

	"github.com/gorilla/mux"
)

func ClassRoutes(r *mux.Router) {
	classRepository := repositories.RepositoryClass(mysql.DB)
	h := handlers.HandlerClass(classRepository)

	r.HandleFunc("/classes", h.FindClasses).Methods("GET")
	r.HandleFunc("/class/{id}", h.GetClass).Methods("GET")
	r.HandleFunc("/class", h.CreateClass).Methods("POST")
	r.HandleFunc("/class/{id}", h.UpdateClass).Methods("PATCH")
	// r.HandleFunc("/class/{id}", middleware.Auth(h.UpdateClass)).Methods("PATCH")
	r.HandleFunc("/class/{id}", h.DeleteClass).Methods("DELETE")
	// r.HandleFunc("/class/{id}", middleware.Auth(h.DeleteClass)).Methods("DELETE")

}
