package main

import (
	"goyak/handler"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

/*
	Think of handler as controller in MVC frameworks. Every controller has controller methods for list, retrieve, update,
	create and delete. Note that I am using Django convention here. List is equivalent to index in Rails; retrieve is
	equivalent to show in Rails.
*/

func LoadRoutes() http.Handler {
	// Defining middleware
	logMiddleware := NewServerLoggingMiddleware()

	muxRouter := mux.NewRouter().StrictSlash(true)

	// Name-spacing the API
	api := muxRouter.PathPrefix("/api").Subrouter()

	api.Handle("/users/", handler.NewUserListHandler()).Methods("GET")
	api.Handle("/users/{id:[0-9]+}", handler.NewUserRetrieveHandler()).Methods("GET")
	api.Handle("/users/", handler.NewUserCreateHandler()).Methods("POST")
	api.Handle("/authenticate/", handler.NewSessionCreateHandler()).Methods("POST")

	muxRouter.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	return handlers.CORS()(logMiddleware(muxRouter))
}
