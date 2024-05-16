package routes

import (
	"preschool/handlers"
	// "preschool/pkg/middleware"
	"preschool/pkg/mysql"
	"preschool/repositories"

	"github.com/gorilla/mux"
)

func PageRoutes(r *mux.Router) {
	pageRepository := repositories.RepositoryPage(mysql.DB)
	h := handlers.HandlerPage(pageRepository)

	r.HandleFunc("/pages", h.FindPages).Methods("GET")
	r.HandleFunc("/page/{id}", h.GetPage).Methods("GET")
	r.HandleFunc("/page", h.CreatePage).Methods("POST")
	// r.HandleFunc("/page", middleware.Auth(h.CreatePage)).Methods("POST")
	r.HandleFunc("/page/{id}", h.UpdatePage).Methods("PATCH")
	// r.HandleFunc("/page/{id}", middleware.Auth(h.UpdatePage)).Methods("PATCH")
	r.HandleFunc("/page/{id}", h.DeletePage).Methods("DELETE")
	// r.HandleFunc("/page/{id}", middleware.Auth(h.DeletePage)).Methods("DELETE")

}
