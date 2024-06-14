package api

import (
	"hello/internals/app/handlers"

	"github.com/gorilla/mux"
)

func CreateRoutes(deployHandler *handlers.DeployHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/deploys/create", deployHandler.Create).Methods("POST")
	r.NotFoundHandler = r.NewRoute().HandlerFunc(handlers.NotFound).GetHandler()
	return r
}
