package main

import (
	"errors"
	"os"
)

// Environment variables used for configuration
const (
	EnvSlackToken     = "SLACK_TOKEN"
	EnvSlackChannelID = "SLACK_CHANNEL_ID"
)

// Config loaded config from environment
type Config struct {
	// Slack config
	SlackToken     string
	SlackChannelID string
}

// Validate validates the configuration has been loaded correctly
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
