package main

import (
	"context"

	"github.com/nulab/go-typetalk/typetalk/v1"
)

func main() {
	client := v1.NewClient(nil)
	client.SetTypetalkToken("yourTypetalkToken")
	ctx := context.Background()
	topicID := 1
	message := "Hello"
	client.Messages.PostMessage(ctx, topicID, message, nil)
}
