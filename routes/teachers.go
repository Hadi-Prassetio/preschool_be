package routes

import (
	"preschool/handlers"
	"preschool/pkg/middleware"
	"preschool/pkg/mysql"
	"preschool/repositories"

	"github.com/gorilla/mux"
)

func TeacherRoutes(r *mux.Router) {
	teacherRepository := repositories.RepositoryTeacher(mysql.DB)
	h := handlers.HandlerTeacher(teacherRepository)

	r.HandleFunc("/teachers", h.FindTeachers).Methods("GET")
	r.HandleFunc("/teacher/{id}", h.GetTeacher).Methods("GET")
	r.HandleFunc("/teacher", middleware.Auth(middleware.UploadFile(h.CreateTeacher))).Methods("POST")
	r.HandleFunc("/teacher/{id}", middleware.Auth(middleware.UploadFile(h.UpdateTeacher))).Methods("PATCH")
	r.HandleFunc("/teacher/{id}", middleware.Auth(h.DeleteTeacher)).Methods("DELETE")

}
