package controllers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	channels "github.com/tzilist/m-service/api/channel"
	messages "github.com/tzilist/m-service/api/message"
	"github.com/tzilist/m-service/pkg/util/responses"
	"github.com/tzilist/m-service/pkg/util/validator"
)

// GetChannelMessages returns all or some channel messages from an ID
func GetChannelMessages(ctx echo.Context) error {
	channelName := ctx.Param("channel")

	if !validator.IsAlphanumericOrDash(channelName) {
		return responses.BadRequest("Invalid Channel Name")
	}

	channel, err := channels.FindChannelByName(channelName)

	if err != nil {
		if err == channels.ChannelNotFound {
			return responses.NotFound("Channel not found")
		}

		return responses.InternalServerError("Error finding channel")
	}

	rawLastID := ctx.QueryParam("last_id")

	var rawMessages []messages.Message

	if rawLastID == "" {
		rawMessages = channel.GetAllMessages()
	} else {
		lastID, err := uuid.Parse(rawLastID)
		if err != nil {
			return responses.BadRequest("Invalid message ID")
		}

		rawMessages, err = channel.GetMessagesSinceID(lastID)

		if err != nil {
			if err == channels.MessageNotFound {
				return responses.NotFound("Message not found")
			}
		}

	}

	var messageResponses []messages.GetMessageResponse
	for _, m := range rawMessages {
		messageResponses = append(messageResponses, m.MapToGetResponse())
	}

	return ctx.JSON(http.StatusOK, messageResponses)
}

// PostChannelMessage controller for handling posating to /:channel/message
func PostChannelMessage(ctx echo.Context) error {
	channelName := ctx.Param("channel")

	if !validator.IsAlphanumericOrDash(channelName) {
		return responses.BadRequest("Invalid Channel Name")
	}

	channel, err := channels.FindChannelByName(channelName)
	if err != nil {
		if err != channels.ChannelNotFound {
			return responses.InternalServerError("Error finding channel")
		}

		channel = channels.New(channelName)
		channel.AddToGlobalChannelList()
	}

	messageRequest := new(messages.CreateMessageRequest)

	if err := ctx.Bind(messageRequest); err != nil {
		return responses.BadRequest("Failed to parse JSON to create message request")
	}

	if err = ctx.Validate(messageRequest); err != nil {
		return responses.BadRequest("Invalid message data")
	}

	message := messageRequest.MapToMessage()

	channel.AddMessage(message)

	return ctx.JSON(http.StatusCreated, message.MapToCreateResponse())
}
