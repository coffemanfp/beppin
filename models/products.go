package models

import "time"

// Product - Product for the app.
type Product struct {
	ID     int64  `json:"id,omitempty"`
	UserID int64  `json:"userID,omitempty"`
	Offer  *Offer `json:"offer,omitempty"`

	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	Categories  []string `json:"categories,omitempty"`
	Price       float64  `json:"price"`

	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// Products - Alias for a product array.
type Products []Product

// Validate - Validates a product.
func (p Product) Validate() (valid bool) {
	valid = true

	switch "" {
	case p.Name:
	case p.Description:
		valid = false
		return
	}

	switch 0 {
	case int(p.UserID):
	case int(p.Price):
		valid = false
	}
	return
}
