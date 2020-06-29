package models

// ResponseMessage - Response message for a end point.
type ResponseMessage struct {
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
	Content interface{} `json:"content,omitempty"`
}
