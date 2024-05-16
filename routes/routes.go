package routes

import "github.com/gorilla/mux"

func RouteInit(r *mux.Router) {
	TeacherRoutes(r)
	AdminRoutes(r)
	VisitorRoutes(r)
	ClassRoutes(r)
	EnrollmentRoutes(r)
	PageRoutes(r)
}
