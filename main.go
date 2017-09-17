package main

import (
	"fmt"
	"net/http"

	"github.com/michaelrbond/go-rss-aggregator/configuration"
	"github.com/michaelrbond/go-rss-aggregator/database-utils"
)

func main() {
	config := configuration.GetConfig()
	fmt.Printf("Starting Go-RSS-Aggregator\n")

	fmt.Printf("Getting database connection\n")
	db := databaseUtils.GetDatabase(config.Mysql)
	defer databaseUtils.Close(db)

	fmt.Printf("Performing Database Migrations\n")
	databaseUtils.Migrate(db, config.Dbmigrations.Files)

	router := DefineRoutes()
	http.Handle("/", router)
	fmt.Printf("Listening on port %d\n", config.Server.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", config.Server.Port), nil)
}
