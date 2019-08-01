package apichannels

// ChannelError types of errors that can be thrown
type ChannelError string

func (che ChannelError) Error() string {
	return string(che)
}

// ChannelNotFound error for when the channel is not found
const ChannelNotFound = ChannelError("Channel not found")

// MessageNotFound error for when the channel is not found
const MessageNotFound = ChannelError("Channel not found")

// AlreadyExists error for when trying to create a channel but it already exists
const AlreadyExists = ChannelError("Channel already exists")
