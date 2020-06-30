package models

import "time"

// Offer - Offer details of a product.
type Offer struct {
	ID        int `json:"id,omitempty"`
	ProductID int `json:"productId,omitempty"`

	Type        string     `json:"type,omitempty"`
	Value       string     `json:"value,omitempty"`
	ExpiratedAt *time.Time `json:"expiratedAt,omitempty"`
	CreatedAt   *time.Time `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
}

// Offers - Alias for a offer array.
type Offers []Offer
