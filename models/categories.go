package models

// Category - Product category
type Category struct {
	ID int `json:"id,omitempty"`

	Name              string   `json:"name,omitempty"`
	RelatedCategories []string `json:"relatedCategories,omitempty"`
}

// Categories - Alias for a categories array.
type Categories []Category
