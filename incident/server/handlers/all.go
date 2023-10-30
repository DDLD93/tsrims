package handlers

import (
	"github.com/ddld93/incident/server/middlewares"
	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	//public routes
	root := mux.NewRouter()
	root.Use(middlewares.JsonMiddleware)
	root.HandleFunc("/", RootHandler).Methods("GET")
    // Create a subrouter for /upload
    uploadRouter := root.PathPrefix("/upload").Subrouter()

    // Use FileUploadMiddleware for the /upload route
    uploadRouter.Handle("/", middlewares.FileUploadMiddleware("audio",RootHandler)).Methods("POST")
	

	

	// api.HandleFunc("/register", Register).Methods("POST")
	// api.HandleFunc("/login", Login).Methods("POST")

	// // //private routes
	// api.HandleFunc("/", GetUsers).Methods("GET")
	// api.HandleFunc("/{id}", GetUser).Methods("GET")
	// privateRoutes := publicRoutes.NewRoute().Subrouter()
	// //jwt
	// privateRoutes.Use(middlewares.JwtMiddleware)
	// privateRoutes.HandleFunc("/secret", SecretHandler)

	return root
}
