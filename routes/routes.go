package routes

import "github.com/gorilla/mux"

func RouteInit(r *mux.Router) {
	TeacherRoutes(r)
	AuthRoutes(r)
}
