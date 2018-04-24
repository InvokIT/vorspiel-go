package main

import (
	"context"
	"os"

	"cloud.google.com/go/pubsub"
	googlemq "github.com/invokit/vorspiel-lib/google/mq"
	"github.com/invokit/vorspiel-lib/mq"
	"google.golang.org/api/option"
)

func BuildMq() (mq.Client, error) {
	projectId := os.Getenv("GOOGLE_PROJECT_ID")
	apiKey := os.Getenv("GOOGLE_API_KEY")

	options := make([]option.ClientOption, 0)

	if apiKey != "" {
		options = append(options, option.WithAPIKey(apiKey))
	}

	pubsubClient, err := pubsub.NewClient(context.Background(), projectId, options...)
	if err != nil {
		return nil, err
	}

	mqClient := googlemq.New(pubsubClient)

	return mqClient, nil
}