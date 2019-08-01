package apimessages

import "github.com/google/uuid"

// CreateMessageResponse structure of the response from the create message request
type CreateMessageResponse struct {
	ID uuid.UUID `json:"id"`
}

// GetMessageResponse message response for returning full message bodies
type GetMessageResponse struct {
	ID       uuid.UUID `json:"id"`
	UserName string    `json:"username"`
	Message  string    `json:"message"`
}
