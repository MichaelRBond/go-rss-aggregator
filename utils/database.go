package utils

import (
	"database/sql"
	"fmt"

	// load mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/mysql"
	// load the source importer
	_ "github.com/mattes/migrate/source/file"

	"github.com/michaelrbond/go-rss-aggregator/configuration"
	"github.com/michaelrbond/go-rss-aggregator/logger"
)

var database *sql.DB

// DatabaseMigrate runs the MySQL migrationn scripts on startup
func DatabaseMigrate(db *sql.DB, migrationsPath string) {
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", migrationsPath),
		"mysql",
		driver,
	)

	if err != nil {
		logger.Panic(fmt.Sprintf("Error starting DB Migration: %s\n", err.Error()))
	}

	if err = m.Up(); err != nil && err.Error() != "no change" {
		logger.Panic(fmt.Sprintf("Error performing DB Migration: %s\n", err.Error()))
	}
}

// DatabaseGet returns an open database
func DatabaseGet(config configuration.MysqlConfig) *sql.DB {
	if database != nil {
		return database
	}
	connectionStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?multiStatements=true",
		config.User,
		config.Password,
		config.Server,
		config.Port,
		config.Database)
	db, err := sql.Open("mysql", connectionStr)

	if err != nil {
		logger.Panic(fmt.Sprintf("Error getting database: %s\n", err.Error()))
	}

	database = db
	return db
}

// DatabaseClose terminates a database connection
func DatabaseClose(db *sql.DB) {
	db.Close()
}
