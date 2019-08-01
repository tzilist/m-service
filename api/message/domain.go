package apimessages

import (
	"github.com/google/uuid"
)

// Message domain message struct
type Message struct {
	ID       uuid.UUID `json:"id"`
	UserName string    `json:"username"`
	Message  string    `json:"message"`
}

// MapToCreateResponse maps a message to a response
func (m *Message) MapToCreateResponse() CreateMessageResponse {
	return CreateMessageResponse{ID: m.ID}
}
