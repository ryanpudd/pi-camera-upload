package main

import (
	"errors"
	"os"
)

// Environment variables used for configuration
const (
	EnvSlackToken         = "SLACK_TOKEN"
	EnvSlackChannelID     = "SLACK_CHANNEL_ID"
	EnvS3BucketName       = "S3_BUCKET_NAME"
	EnvS3BucketRegion     = "S3_BUCKET_REGION"
	EnvAWSAccessKey       = "AWS_ACCESS_KEY"
	EnvAWSSecretAccessKey = "AWS_SECRET_ACCESS_KEY"
)

// Config Configuration structure
type Config struct {
	// Slack config
	SlackToken     string
	SlackChannelID string
	// AWS config
	S3BucketName       string
	S3BucketRegion     string
	AWSAccessKey       string
	AWSSecretAccessKey string
}

// Validate checks the environments variables required have been set, throws an error otherwise
func (c *Config) Validate() error {
	if c.SlackToken == "" {
		return errors.New("Must set env var " + EnvSlackToken)
	}
	if c.SlackChannelID == "" {
		return errors.New("Must set env var " + EnvSlackChannelID)
	}
	if c.S3BucketName == "" {
		return errors.New("Must set env var " + EnvS3BucketName)
	}
	if c.S3BucketRegion == "" {
		return errors.New("Must set env var " + EnvS3BucketRegion)
	}
	if c.AWSAccessKey == "" {
		return errors.New("Must set env var " + EnvAWSAccessKey)
	}
	if c.AWSSecretAccessKey == "" {
		return errors.New("Must set env var " + EnvAWSSecretAccessKey)
	}

	return nil
}

// NewConfigFromEnv Loads all the environment variables as config
func NewConfigFromEnv() *Config {
	return &Config{
		SlackToken:         os.Getenv(EnvSlackToken),
		SlackChannelID:     os.Getenv(EnvSlackChannelID),
		S3BucketName:       os.Getenv(EnvS3BucketName),
		S3BucketRegion:     os.Getenv(EnvS3BucketRegion),
		AWSAccessKey:       os.Getenv(EnvAWSAccessKey),
		AWSSecretAccessKey: os.Getenv(EnvAWSSecretAccessKey),
	}
}
