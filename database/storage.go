package database

import "database/sql"

// Storage database and database utils.
type Storage interface {
	LanguageStorage
	ProductStorage
	UserStorage

	GetDB() *sql.DB
	SetDB(db *sql.DB)
	CloseDB()
}
