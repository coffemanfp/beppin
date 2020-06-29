package models

import "time"

// Product - Product for the app.
type Product struct {
	ID     int   `json:"id"`
	UserID int   `json:"userId"`
	Offer  Offer `json:"offer"`

	Name        string   `json:"name"`
	Description string   `json:"description"`
	Categories  []string `json:"categories"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
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
