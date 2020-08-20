package models

import (
	"fmt"

	"github.com/coffemanfp/beppin-server/config"
)

// ResponseMessage - Response message for a end point.
type ResponseMessage struct {
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
	Content interface{} `json:"content,omitempty"`
}

// NotLimitParamProvided - Sets a message saying that the limit parameter has not been provided.
func (m *ResponseMessage) NotLimitParamProvided(limit *int) {
	if *limit > 0 {
		return
	}

	settings := config.GetSettings()

	m.Message = fmt.Sprintf(
		"Not limit param provided, setted to %d",
		settings.MaxElementsPerPagination,
	)

	*limit = settings.MaxElementsPerPagination
	return
}

// LimitParamExceeded - Sets a message saying that the limit parameter has  been provided.
func (m *ResponseMessage) LimitParamExceeded(limit *int) {
	settings := config.GetSettings()

	if *limit < settings.MaxElementsPerPagination {
		return
	}

	m.Message = fmt.Sprintf(
		"Limit of elements exceeded, setted to %d",
		settings.MaxElementsPerPagination,
	)

	*limit = settings.MaxElementsPerPagination
	return
}
