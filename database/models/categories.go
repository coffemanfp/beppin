package models

import "database/sql"

// Category - Product category
type Category struct {
	ID int

	Name              string
	RelatedCategories []string
	CreatedAt         *sql.NullTime
}

// GetIdentifier gets the first unique identifier it finds in order of importance.
func (c Category) GetIdentifier() (identifier interface{}) {
	if c.ID != 0 {
		identifier = c.ID
	} else if c.Name != "" {
		identifier = c.Name
	}

	return
}

// Categories - Alias for a categories array.
type Categories []Category
