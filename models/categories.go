package models

import "time"

// Category - Product category
type Category struct {
	ID int64 `json:"id,omitempty"`

	Name              string     `json:"name,omitempty"`
	RelatedCategories []string   `json:"relatedCategories,omitempty"`
	CreatedAt         *time.Time `json:"createdAt,omitempty"`
}

// Categories - Alias for a categories array.
type Categories []Category
