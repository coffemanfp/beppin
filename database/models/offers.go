package models

import "database/sql"

// Offer - Offer details of a product.
type Offer struct {
	ID        int `json:"id,omitempty"`
	ProductID int `json:"productId,omitempty"`

	Type        string        `json:"type,omitempty"`
	Value       string        `json:"value,omitempty"`
	ExpiratedAt *sql.NullTime `json:"expiratedAt,omitempty"`
	CreatedAt   *sql.NullTime `json:"createdAt,omitempty"`
	UpdatedAt   *sql.NullTime `json:"updatedAt,omitempty"`
}

// Offers - Alias for a offer array.
type Offers []Offer
