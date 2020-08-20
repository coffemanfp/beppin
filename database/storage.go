package database

import "database/sql"

// Storage database and database utils.
type Storage interface {
	UserStorage

	GetDB() *sql.DB
	SetDB(db *sql.DB)
	CloseDB()
}
