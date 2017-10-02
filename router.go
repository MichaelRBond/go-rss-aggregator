package main

import (
	"github.com/michaelrbond/go-rss-aggregator/controllers"

	"github.com/gorilla/mux"
)

// DefineRoutes defines routes
func DefineRoutes(context *controllers.Context) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.HTTPHandler{Context: context, H: controllers.DefaultHandler}.ServeHTTP)
	r.HandleFunc("/api/v1/feeds/add", controllers.HTTPHandler{Context: context, H: controllers.RegisterRssFeed}.ServeHTTP).Methods("POST")
	return r
}
