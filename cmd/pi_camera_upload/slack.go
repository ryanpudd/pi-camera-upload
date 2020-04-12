package main

import (
	"github.com/slack-go/slack"
)

// NotifySlack sends the url string to the channel using the token provided
func NotifySlack(token string, channelID string, url string) error {
	api := slack.New(token)
	_, _, err := api.PostMessage(channelID, slack.MsgOptionText(url, false))
	return err
}
