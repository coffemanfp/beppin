package database

import (
	"database/sql"
	"errors"
	"fmt"

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
	settings := config.GetSettings()

	if !settings.Database.ValidateDatabase() {
		err = errors.New(fmt.Sprint("invalid database settings", settings))
		return
	}

	dbConn, err = sql.Open("postgres", settings.Database.URL)
	if err != nil {
		err = fmt.Errorf("error opening a database connection:\n%s", err)
		return
	}

	dbConn.SetMaxOpenConns(1)

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
