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
	s = storage
	return
}

// Set sets the current database storage.
func Set(s Storage) {
	storage = s
	return
}

// New creates a new database storage.
func New(db *sql.DB) (s Storage) {
	s = defaultStorage{db: db}
	storage = s
	return
}

// NewDefault creates the default database storage.
func NewDefault() (s Storage, err error) {
	settings := config.GetSettings()

	if settings.Database == nil || !settings.Database.ValidateDatabase() {
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

	s = defaultStorage{db: db}
	return
}
