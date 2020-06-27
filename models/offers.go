package models

import "time"

// Offer - Offer details of a product.
type Offer struct {
	ID        int `json:"id,omitempty"`
	ProductID int `json:"product_id,omitempty"`

	Type        string    `json:"type,omitempty"`
	Value       string    `json:"value,omitempty"`
	ExpiratedAt time.Time `json:"expirated_at,omitempty"`
}

// Offers - Alias for a offer array.
type Offers []Offer
