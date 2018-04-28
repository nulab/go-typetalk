package typetalk_token

import (
	"context"

	"github.com/nulab/go-typetalk/typetalk/v1"
)

func main() {
	client := v1.NewClient(nil)
	client.SetTypetalkToken("yourTypetalkToken")
	ctx := context.Background()
	topicId := 1
	message := "Hello"
	profile, resp, err := client.Messages.PostMessage(ctx, topicId, message, nil)
}
