package apimessages

import "github.com/google/uuid"

// CreateMessageRequest structure of a create message request
type CreateMessageRequest struct {
	UserName string `json:"username" validate:"required"`
	Message  string `json:"message" validate:"required"`
}

// MapToMessage maps a CreateMessageRequest to a Message
func (mreq *CreateMessageRequest) MapToMessage() Message {
	return Message{
		UserName: mreq.UserName,
		Message:  mreq.Message,
		ID:       uuid.New(),
	}
}
