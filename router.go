package main

import "github.com/gorilla/mux"
import "github.com/michaelrbond/go-rss-aggregator/controllers"

// DefineRoutes defines routes
func DefineRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.DefaultHandler)
	r.HandleFunc("/api/v1/feeds/add", controllers.RegisterRssFeed).Methods("POST")
	return r
}
