package main

import (
	"github.com/michaelrbond/go-rss-aggregator/controllers"
	"github.com/michaelrbond/go-rss-aggregator/types"

	"github.com/gorilla/mux"
)

// DefineRoutes defines routes
func DefineRoutes(context *types.Context) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", types.HTTPHandler{Context: context, H: controllers.DefaultHandler}.ServeHTTP)
	r.HandleFunc("/api/v1/feeds/add", types.HTTPHandler{Context: context, H: controllers.RegisterRssFeed}.ServeHTTP).Methods("POST")
	return r
}
