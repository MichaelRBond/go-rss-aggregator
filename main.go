package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/michaelrbond/go-rss-aggregator/configuration"
	"github.com/michaelrbond/go-rss-aggregator/database-utils"
	"github.com/michaelrbond/go-rss-aggregator/logger"
	"github.com/michaelrbond/go-rss-aggregator/syncEngine"
	"github.com/michaelrbond/go-rss-aggregator/types"
)

func main() {
	config := configuration.GetConfig()
	logger.Info("Starting Go-RSS-Aggregator")

	logger.Info("Getting database connection")
	db := databaseUtils.GetDatabase(config.Mysql)
	defer databaseUtils.Close(db)

	logger.Info("Performing Database Migrations")
	databaseUtils.Migrate(db, config.Dbmigrations.Files)

	context := &types.Context{Config: config, Db: db}

	router := DefineRoutes(context)
	http.Handle("/", router)
	logger.Info(fmt.Sprintf("Listening on port %d", config.Server.Port))
	go http.ListenAndServe(fmt.Sprintf(":%d", config.Server.Port), nil)

	syncEngineTicker := time.NewTicker(config.SyncEngine.IntervalInSeconds * time.Second)
	for range syncEngineTicker.C {
		syncEngine.SyncRssFeeds(context)
	}

	// TODO : Catch SIGINT and clean up syncEngineTicker
}
