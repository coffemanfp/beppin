package models

import "database/sql"

// Category - Product category
type Category struct {
	ID int

	Name              string
	RelatedCategories []string
	CreatedAt         *sql.NullTime
}

// Categories - Alias for a categories array.
type Categories []Category
