package database

import (
	"database/sql"
	"fmt"

	"github.com/coffemanfp/beppin-server/config"
	errs "github.com/coffemanfp/beppin-server/errors"
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
		err = fmt.Errorf("%w", errs.ErrInvalidSettings)
		return
	}

	dbConn, err = sql.Open("postgres", settings.Database.URL)
	if err != nil {
		err = fmt.Errorf("error opening a database connection: %v", err)
		return
	}

	dbConn.SetMaxOpenConns(1)

	db = dbConn

	err = dbConn.Ping()
	if err != nil {
		err = fmt.Errorf("error in ping to the database: %v", err)
	}

	return
}

// CloseConn - ...
func CloseConn() {
	if db == nil {
		return
	}

	db.Close()
	db = nil
}
