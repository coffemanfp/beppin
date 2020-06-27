package models

import "time"

// Product - Product for the app.
type Product struct {
	ID     int   `json:"id"`
	UserID int   `json:"userId"`
	Offer  Offer `json:"offer"`

	Name        string `json:"name"`
	Description string `json:"description"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
