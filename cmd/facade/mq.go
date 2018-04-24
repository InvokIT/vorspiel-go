package main

import (
	"context"
	"os"

	"cloud.google.com/go/pubsub"
	"github.com/invokit/vorspiel-backend/pkg/gcp/gcpmq"
	"github.com/invokit/vorspiel-backend/pkg/mq"
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

	mqClient := gcpmq.New(pubsubClient)

	return mqClient, nil
}
