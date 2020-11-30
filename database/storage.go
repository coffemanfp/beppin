package database

import "database/sql"

// Storage database and database utils.
type Storage interface {
	LanguageStorage
	ProductStorage
	CategoriesStorage
	UserStorage
	FileStorage

	GetDB() *sql.DB
	SetDB(db *sql.DB)
	CloseDB()
}
