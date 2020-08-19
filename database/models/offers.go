package models

import "database/sql"

// Offer - Offer details of a product.
type Offer struct {
	ID        int64 `json:"id,omitempty"`
	ProductID int64 `json:"productId,omitempty"`

	Type        string        `json:"type,omitempty"`
	Value       string        `json:"value,omitempty"`
	ExpiratedAt *sql.NullTime `json:"expiratedAt,omitempty"`
	CreatedAt   *sql.NullTime `json:"createdAt,omitempty"`
	UpdatedAt   *sql.NullTime `json:"updatedAt,omitempty"`
}

// GetIdentifier gets the first unique identifier it finds in order of importance.
func (o Offer) GetIdentifier() (identifier interface{}) {
	if o.ID != 0 {
		identifier = o.ID
	}

	return
}

// Offers - Alias for a offer array.
type Offers []Offer
