package routes

import (
	"preschool/handlers"
	// "preschool/pkg/middleware"
	"preschool/pkg/mysql"
	"preschool/repositories"

	"github.com/gorilla/mux"
)

func EnrollmentRoutes(r *mux.Router) {
	enrollmentRepository := repositories.RepositoryEnrollment(mysql.DB)
	h := handlers.HandlerEnrollment(enrollmentRepository)

	r.HandleFunc("/enrollments", h.FindEnrollments).Methods("GET")
	r.HandleFunc("/enrollment-wait", h.FindEnrollmentWait).Methods("GET")
	r.HandleFunc("/enrollment-accept", h.FindEnrollmentAccepted).Methods("GET")
	r.HandleFunc("/enrollment-cancel", h.FindEnrollmentCancel).Methods("GET")
	r.HandleFunc("/enrollment/{id}", h.GetEnrollment).Methods("GET")
	r.HandleFunc("/enrollment", h.CreateEnrollment).Methods("POST")
	r.HandleFunc("/enrollment/{id}", h.UpdateEnrollment).Methods("PATCH")
	// r.HandleFunc("/enrollment/{id}", middleware.Auth(h.UpdateEnrollment)).Methods("PATCH")
	r.HandleFunc("/enrollment/{id}", h.DeleteEnrollment).Methods("DELETE")
	// r.HandleFunc("/enrollment/{id}", middleware.Auth(h.DeleteEnrollment)).Methods("DELETE")

}
