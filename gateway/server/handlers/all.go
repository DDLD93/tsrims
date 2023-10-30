package handlers

import (
	"github.com/ddld93/tsrims/auth/server/middlewares"
	"github.com/ddld93/tsrims/auth/utils"
	"github.com/gorilla/mux"
)

var cfg *utils.Config

func init() {
	res := utils.LoadEvn()
	cfg = res
}

func GetRouter() *mux.Router {
	router := mux.NewRouter()

	router.Use(middlewares.JsonMiddleware)

	router.PathPrefix("/api/v1/nin").Handler(proxyHandler(cfg.MicroServices.NIN))
	router.PathPrefix("/api/v1/incident").Handler(proxyHandler(cfg.MicroServices.IncidentService))
	//
	router.HandleFunc("/", RootHandler)
	router.HandleFunc("/api/v1/auth/register", Register).Methods("POST")
	router.HandleFunc("/api/v1/auth/login", Login).Methods("POST")

	return router
}
