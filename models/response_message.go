package models

import (
	"fmt"

	"github.com/coffemanfp/beppin-server/config"
)

// Helper consts to get the content type on the Content field.
const (
	TypeAvatar = "avatar"

	TypeCategories = "categories"
	TypeCategory   = "category"

	TypeLanguages = "languages"
	TypeLanguage  = "language"

	TypeOffers = "offers"
	TypeOffer  = "offer"

	TypeProducts = "products"
	TypeProduct  = "product"

	TypeUsers = "users"
	TypeUser  = "user"

	TypeToken = "token"
)

// ResponseMessage - Response message for a end point.
type ResponseMessage struct {
	Message     string      `json:"message,omitempty"`
	Error       string      `json:"error,omitempty"`
	Content     interface{} `json:"content,omitempty"`
	ContentType string      `json:"contentType,omitempty"` // Helper field to find the object content type on the Content field.
}

// NotLimitParamProvided - Sets a message saying that the limit parameter has not been provided.
func (m *ResponseMessage) NotLimitParamProvided(limit *uint64) {
	if *limit > 0 {
		return
	}

	settings := config.GetSettings()

	m.Message = fmt.Sprintf(
		"Not limit param provided, setted to %d",
		settings.MaxElementsPerPagination,
	)

	*limit = uint64(settings.MaxElementsPerPagination)
	return
}

// LimitParamExceeded - Sets a message saying that the limit parameter has  been provided.
func (m *ResponseMessage) LimitParamExceeded(limit *uint64) {
	settings := config.GetSettings()

	if *limit < settings.MaxElementsPerPagination {
		return
	}

	m.Message = fmt.Sprintf(
		"Limit of elements exceeded, setted to %d",
		settings.MaxElementsPerPagination,
	)

	*limit = uint64(settings.MaxElementsPerPagination)
	return
}
