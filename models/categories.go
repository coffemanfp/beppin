package models

import "time"

// Category - Product category
type Category struct {
	ID int64 `json:"id,omitempty"`

	Name              string     `json:"name,omitempty"`
	Description       string     `json:"description,omitempty"`
	RelatedCategories []string   `json:"relatedCategories,omitempty"`
	CreatedAt         *time.Time `json:"createdAt,omitempty"`
	UpdatedAt         *time.Time `json:"updatedAt,omitempty"`
}

// Categories - Alias for a categories array.
type Categories []Category

// GetIdentifier gets the first unique identifier it finds in order of importance.
func (c Category) GetIdentifier() (identifier interface{}) {
	if c.ID != 0 {
		identifier = c.ID
	} else if c.Name != "" {
		identifier = c.Name
	}

	return
}
