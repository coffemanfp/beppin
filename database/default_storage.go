package database

import (
	"database/sql"
)

type defaultStorage struct {
	db *sql.DB
}

// GetDB returns the connection to the database
func (dS defaultStorage) GetDB() *sql.DB {
	return dS.db
}
func (dS defaultStorage) SetDB(db *sql.DB) {
	dS.db = db
}

// CloseDB closes the connection to the database
func (dS defaultStorage) CloseDB() {
	if dS.db == nil {
		return
	}

	dS.db.Close()
	dS.db = nil
}
