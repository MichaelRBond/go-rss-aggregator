package main

import (
	"fmt"
	"net/http"

	"github.com/michaelrbond/go-rss-aggregator/configuration"
	"github.com/michaelrbond/go-rss-aggregator/controllers"
	"github.com/michaelrbond/go-rss-aggregator/database-utils"
	"github.com/michaelrbond/go-rss-aggregator/logger"
)

func main() {
	config := configuration.GetConfig()
	logger.Info("Starting Go-RSS-Aggregator")

	logger.Info("Getting database connection")
	db := databaseUtils.GetDatabase(config.Mysql)
	defer databaseUtils.Close(db)

	logger.Info("Performing Database Migrations")
	databaseUtils.Migrate(db, config.Dbmigrations.Files)

	context := &controllers.Context{Config: config, Db: db}

	router := DefineRoutes(context)
	http.Handle("/", router)
	logger.Info(fmt.Sprintf("Listening on port %d", config.Server.Port))
	http.ListenAndServe(fmt.Sprintf(":%d", config.Server.Port), nil)
}
