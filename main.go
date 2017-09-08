package main

import "fmt"
import "net/http"

func main() {
	config := GetConfig()
	fmt.Printf("Starting Go-RSS-Aggregator\n")
	router := DefineRoutes()
	http.Handle("/", router)
	fmt.Printf("Listening on port %d", config.Server.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", config.Server.Port), nil)
}
