package main

import (
	"fmt"
	"log"
	"os"

	"github.com/slack-go/slack"
)

// MessageCallback function for processing events from slack
type MessageCallback func(ev *slack.MessageEvent, rtm *slack.RTM)

//WebSocketProcessor processes events on the webhook API
type WebSocketProcessor struct {
	Token           string
	ChannelID       string
	MessageCallback MessageCallback
}

// NewWebSocketProcessor Creates a WebSocketProcessor using the slack token provided
func NewWebSocketProcessor(token string, channelID string, callback MessageCallback) *WebSocketProcessor {
	processor := new(WebSocketProcessor)
	processor.Token = token
	processor.ChannelID = channelID
	processor.MessageCallback = callback
	return processor
}

// Start Calls Slack API and begins listening for events
func (processor *WebSocketProcessor) Start() {
	api := slack.New(
		processor.Token,
		slack.OptionDebug(false),
		slack.OptionLog(log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)),
	)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	fmt.Println("Waiting for events")

	for msg := range rtm.IncomingEvents {
		//fmt.Print("Event Received: ")
		switch ev := msg.Data.(type) {
		case *slack.HelloEvent:
			// Ignore hello

		case *slack.ConnectedEvent:
			fmt.Println("Infos:", ev.Info)
			fmt.Println("Connection counter:", ev.ConnectionCount)
			rtm.SendMessage(rtm.NewOutgoingMessage("I'm back online", processor.ChannelID))

		case *slack.MessageEvent:
			processor.MessageCallback(ev, rtm)

		/*case *slack.PresenceChangeEvent:
			fmt.Printf("Presence Change: %v\n", ev)

		case *slack.LatencyReport:
			fmt.Printf("Current latency: %v\n", ev.Value)

		case *slack.DesktopNotificationEvent:
			fmt.Printf("Desktop Notification: %v\n", ev)
		*/
		case *slack.RTMError:
			fmt.Printf("Error: %s\n", ev.Error())

		case *slack.InvalidAuthEvent:
			fmt.Printf("Invalid credentials")
			return

		default:

			// Ignore other events..
			// fmt.Printf("Unexpected: %v\n", msg.Data)
		}
	}
}
