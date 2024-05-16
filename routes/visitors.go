package routes

import (
	"preschool/handlers"
	// "preschool/pkg/middleware"
	"preschool/pkg/mysql"
	"preschool/repositories"

	"github.com/gorilla/mux"
)

func VisitorRoutes(r *mux.Router) {
	visitorRepository := repositories.RepositoryVisitor(mysql.DB)
	h := handlers.HandlerVisitor(visitorRepository)

	r.HandleFunc("/visitors", h.FindVisitors).Methods("GET")
	r.HandleFunc("/visitor-wait", h.FindVisitorWait).Methods("GET")
	r.HandleFunc("/visitor-accept", h.FindVisitorAccepted).Methods("GET")
	r.HandleFunc("/visitor-cancel", h.FindVisitorCancel).Methods("GET")
	r.HandleFunc("/visitor/{id}", h.GetVisitor).Methods("GET")
	r.HandleFunc("/visitor", h.CreateVisitor).Methods("POST")
	r.HandleFunc("/visitor/{id}", h.UpdateVisitor).Methods("PATCH")
	// r.HandleFunc("/visitor/{id}", middleware.Auth(h.UpdateVisitor)).Methods("PATCH")
	r.HandleFunc("/visitor/{id}", h.DeleteVisitor).Methods("DELETE")
	// r.HandleFunc("/visitor/{id}", middleware.Auth(h.DeleteVisitor)).Methods("DELETE")

}
