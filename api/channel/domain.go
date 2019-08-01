package apichannels

import (
	"sync"

	"github.com/google/uuid"
	messages "github.com/tzilist/m-service/api/message"
)

// Channel struct of a channel
type Channel struct {
	Name     string `validate:"required, validchannelname"`
	Messages []messages.Message
	// prevent reads and writes at the same time
	mux sync.Mutex
}

// XXX: this is a hack because there is no database for this challenge :)
// Everything below would idealy be extracted into a models/repositories folder
// for interacting with the database

// Channels Map of channels
var Channels = map[string]*Channel{}

// New creates a new channel (note this is not a pointer)
func New(name string) *Channel {
	return &Channel{
		Name:     name,
		Messages: []messages.Message{},
	}

}

// AddMessage appends a message to the list of messages
func (c *Channel) AddMessage(message messages.Message) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.Messages = append(c.Messages, message)
}

// GetAllMessages returns all messages in a channel
func (c *Channel) GetAllMessages() []messages.Message {
	c.mux.Lock()
	defer c.mux.Unlock()

	return c.Messages
}

// GetMessagesSinceID gets all messages from an ID and later
func (c *Channel) GetMessagesSinceID(id uuid.UUID) ([]messages.Message, error) {
	c.mux.Lock()
	defer c.mux.Unlock()

	messages := []messages.Message{}
	found := false
	for _, m := range c.Messages {
		// if the id has been found already, append to the slice to return
		if found {
			messages = append(messages, m)
			continue
		}

		// if the ids are equal, we will set found to true to return everything after this message
		if id == m.ID {
			found = true
			continue
		}
	}
	if !found {
		return nil, MessageNotFound
	}
	return messages, nil
}

// AddToGlobalChannelList addes the channel to the global channel list
func (c *Channel) AddToGlobalChannelList() error {
	if _, ok := Channels[c.Name]; ok {
		return AlreadyExists
	}

	Channels[c.Name] = c

	return nil
}

// FindChannelByName finds a channel from the global channel list
func FindChannelByName(name string) (*Channel, error) {
	if channel, ok := Channels[name]; ok {
		return channel, nil
	}

	return nil, ChannelNotFound
}
