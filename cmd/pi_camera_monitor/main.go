package main

import (
	"fmt"
	"log"
)

func main() {
	// Load the config
	cfg := NewConfigFromEnv()
	if err := cfg.Validate(); err != nil {
		log.Println(err)
		return
	}

	userID, err := GetBotUserID(cfg.SlackToken)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("Bot User ID: %s\n", userID)

	handler := NewMessageHandler(userID, cfg.SlackChannelID)

	// Start the slack listener
	processor := NewWebSocketProcessor(cfg.SlackToken, cfg.SlackChannelID, handler.ProcessMessage)
	processor.Start()
}
