package models

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

// Common messages
var (
	MessageNotLimitParam      = "Not limit param provided"
	MessageLimitParamExceeded = "Limit of elements exceeded"
)

// ResponseMessage - Response message for a end point.
type ResponseMessage struct {
	Message     string      `json:"message,omitempty"`
	Error       string      `json:"error,omitempty"`
	Content     interface{} `json:"content,omitempty"`
	ContentType string      `json:"contentType,omitempty"` // Helper field to find the object content type on the Content field.
}
