package api

import (
	"hello/internals/app/handlers"

	"github.com/gorilla/mux"
)

func CreateRoutes(deployHandler *handlers.DeployHandler, healthHandler *handlers.HealthHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/deployment/create", deployHandler.Create).Methods("POST")
	r.HandleFunc("api/health", healthHandler.Status).Methods("GET")
	r.NotFoundHandler = r.NewRoute().HandlerFunc(handlers.NotFound).GetHandler()
	return r
}
