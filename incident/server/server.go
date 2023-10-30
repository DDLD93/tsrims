package server

import (
	"log"
	"net/http"
	"time"

	"github.com/ddld93/incident/server/handlers"
	"github.com/ddld93/incident/utils"
)

var appPort string

func init() {
	cfg := utils.LoadEvn()
	appPort = cfg.AppPort
}

func Start() {
	router := handlers.GetRouter()
	srv := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:" + appPort,
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	log.Printf("Server Listening on port: %s", appPort)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("Error starting server!")
	}
}
