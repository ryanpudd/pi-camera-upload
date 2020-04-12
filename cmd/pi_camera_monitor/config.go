package main

import (
	"errors"
	"os"
)

const (
	EnvSlackToken     = "SLACK_TOKEN"
	EnvSlackChannelID = "SLACK_CHANNEL_ID"
)

type Config struct {
	// Slack config
	SlackToken     string
	SlackChannelID string
}

func (c *Config) Validate() error {
	if c.SlackToken == "" {
		return errors.New("Must set env var " + EnvSlackToken)
	}
	if c.SlackChannelID == "" {
		return errors.New("Must set env var " + EnvSlackChannelID)
	}

	return nil
}

// NewConfigFromEnv Loads all the environment variables as config
func NewConfigFromEnv() *Config {
	return &Config{
		SlackToken:     os.Getenv(EnvSlackToken),
		SlackChannelID: os.Getenv(EnvSlackChannelID),
	}
}
