package models

import "database/sql"

// Product - Product for the app.
type Product struct {
	ID     int
	UserID int
	Offer  *Offer

	Name        string
	Description string
	Categories  []string

	CreatedAt *sql.NullTime
	UpdatedAt *sql.NullTime
}

// Products - Alias for a product array.
type Products []Product

// ValidateUpdate - Validates a product for update.
func (p Product) ValidateUpdate() (valid bool) {
	valid = true

	switch "" {
	case p.Name:
	case p.Description:
		valid = false
	}

	return
}

// Validate - Validates a product.
func (p Product) Validate() (valid bool) {
	valid = true

	valid = p.ValidateUpdate()

	if p.UserID == 0 {
		valid = false
	}
	return
}
