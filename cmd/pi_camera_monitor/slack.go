package main

import (
	"github.com/slack-go/slack"
)

// GetBotUserID converts token into user ID
func GetBotUserID(token string) (string, error) {
	api := slack.New(token)
	response, err := api.AuthTest()

	if err != nil {
		return "", nil
	}

	return response.UserID, nil
}
