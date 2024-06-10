package api

import (
	"hello/internals/app/handlers"

	"github.com/gorilla/mux"
)

func CreateRoutes(userHandler *handlers.UsersHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/users/create", userHandler.Create).Methods("POST")
	r.NotFoundHandler = r.NewRoute().HandlerFunc(handlers.NotFound).GetHandler()
	return r
}
