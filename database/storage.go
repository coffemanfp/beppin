package database

import "database/sql"

// Storage database and database utils.
type Storage interface {
	UserStorage
	ProductStorage

	GetDB() *sql.DB
	SetDB(db *sql.DB)
	CloseDB()
}
