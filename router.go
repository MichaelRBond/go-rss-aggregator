package main

import "github.com/gorilla/mux"
import "github.com/michaelrbond/go-rss-aggregator/controllers"

func DefineRoutes() (*mux.Router) {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.DefaultHandler)
	return r
}