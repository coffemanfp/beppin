package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/coffemanfp/beppin-server/config"
	_ "github.com/lib/pq"
)

var db *sql.DB

// Get - Get the conn to the database.
func Get() (dbConn *sql.DB, err error) {
	if db != nil {
		dbConn = db
		return
	}

	dbConn, err = OpenConn()
	if err != nil {
		return
	}

	db = dbConn
	return
}

// OpenConn - Open a conn to the database.
func OpenConn() (dbConn *sql.DB, err error) {
	settings, err := config.GetSettings()
	if err != nil {
		return
	}

	log.Println(settings.Database)

	if !settings.ValidateDatabase() {
		err = errors.New(fmt.Sprint("database settings are not populated", settings))
		return
	}

	dbConn, err = sql.Open("postgres", fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		settings.Database.User,
		settings.Database.Password,
		settings.Database.Name,
		settings.Database.Host,
		settings.Database.Port,
	))

	// dbConn, err = sql.Open("postgres", fmt.Sprintf(
	// 	"user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
	// 	"beppin",
	// 	"b26edb839ea81fe0801f9d47f17aaeaac2e1162fe28a3487f51a3ee7716d1ef7",
	// 	"beppin",
	// 	"localhost",
	// 	5432,
	// ))

	if err != nil {
		err = fmt.Errorf("error opening a database connection:\n%s", err)
		return
	}

	dbConn.SetMaxIdleConns(1)

	db = dbConn

	err = dbConn.Ping()
	if err != nil {
		err = fmt.Errorf("error in ping to the database:\n%s", err)
	}

	return
}

// CloseConn - ...
func CloseConn() {
	if db == nil {
		return
	}

	db.Close()
}
