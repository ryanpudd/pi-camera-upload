package main

import (
	"fmt"
	"strings"

	"github.com/slack-go/slack"
)

// MessageHandler handles messages from Slack
type MessageHandler struct {
	UserIDPrefix string
	ChannelID    string
}

// NewMessageHandler creates a message handler
func NewMessageHandler(userID string, channelID string) *MessageHandler {
	handler := new(MessageHandler)
	handler.UserIDPrefix = fmt.Sprintf("<@%s>", userID)
	handler.ChannelID = channelID
	return handler
}

// ProcessMessage callback handler from Slack
func (handler *MessageHandler) ProcessMessage(ev *slack.MessageEvent, rtm *slack.RTM) {
	fmt.Printf("Message Callback: %v\n", ev)
	fmt.Printf("Message Content: %s\n", ev.Msg.Text)

	msgContent := ev.Msg.Text
	action := strings.ToLower(strings.TrimSpace(After(msgContent, handler.UserIDPrefix)))
	response := ""

	switch action {
	case "status":
		response = GetStatus()
	}

	if response != "" {
		rtm.SendMessage(rtm.NewOutgoingMessage(response, handler.ChannelID))
	}
}

// GetStatus Creates the status string
func GetStatus() string {
	/*
	 TODO - Check Monitor is running
	 TODO - Check Disk space
	*/
	return "All good"
}
