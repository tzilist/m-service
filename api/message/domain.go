package apimessages

import (
	"github.com/google/uuid"
)

// Message domain message struct
type Message struct {
	ID       uuid.UUID
	UserName string
	Message  string
}

// MapToCreateResponse maps a message to a response
func (m *Message) MapToCreateResponse() CreateMessageResponse {
	return CreateMessageResponse{ID: m.ID}
}

// MapToGetResponse maps a message to a response
func (m *Message) MapToGetResponse() GetMessageResponse {
	return GetMessageResponse{
		ID:       m.ID,
		UserName: m.UserName,
		Message:  m.Message,
	}
}
