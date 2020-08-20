package database

import (
	"database/sql"
	"fmt"

	"github.com/coffemanfp/beppin-server/config"
	errs "github.com/coffemanfp/beppin-server/errors"
)

var storage Storage

// Get gets the current database storage.
func Get() (s Storage, err error) {
	if storage != nil {
		s = storage
		return
	}

	storage, err = NewDefault()
	return
}

// Set sets the current database storage.
func Set(s Storage) {
	storage = s
	return
}

// NewDefault returns the default database storage.
func NewDefault() (storage Storage, err error) {
	settings := config.GetSettings()

	if !settings.Database.ValidateDatabase() {
		err = fmt.Errorf("%w", errs.ErrInvalidSettings)
		return
	}

	db, err := sql.Open("postgres", settings.Database.URL)
	if err != nil {
		err = fmt.Errorf("error opening a database connection: %v", err)
		return
	}

	db.SetMaxOpenConns(1)

	err = db.Ping()
	if err != nil {
		err = fmt.Errorf("error in ping to the database: %v", err)
		return
	}

	storage = defaultStorage{db: db}
	return
}
