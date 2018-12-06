package tests

import (
	"context"
	"testing"

	"time"

	. "github.com/nulab/go-typetalk/typetalk/v2"
)

func Test_V1_Messages_GetMessage_should_get_a_message(t *testing.T) {
	result, resp, err := clientV1.Messages.GetMessage(context.Background(), topicID, postID)
	test(t, result, resp, err)
}

func Test_V1_Messages_PostMessage_should_post_a_message(t *testing.T) {
	result, resp, err := clientV1.Messages.PostMessage(context.Background(), topicID, "go-typetalk - Test_Messages_PostMessage_should_post_a_message", nil)
	test(t, result, resp, err)
}

func Test_V1_Messages_UpdateMessage_should_update_a_message(t *testing.T) {
	result, resp, err := clientV1.Messages.UpdateMessage(context.Background(), topicID, postID, "go-typetalk - Test_Messages_UpdateMessage_should_update_a_message")
	test(t, result, resp, err)
}

func Test_V1_Messages_GetMessage_should_get_a_message_using_Typetalk_Token(t *testing.T) {
	result, resp, err := clientUsingTypetalkTokenV1.Messages.GetMessage(context.Background(), topicID, postID)
	test(t, result, resp, err)
}

func Test_V1_Messages_PostMessage_should_post_a_message_using_Typetalk_Token(t *testing.T) {
	result, resp, err := clientUsingTypetalkTokenV1.Messages.PostMessage(context.Background(), topicID, "go-typetalk - Test_Messages_PostMessage_should_post_a_message_using_Typetalk_Token", nil)
	test(t, result, resp, err)
}

func Test_V2_Messages_SearchMessages_should_some_messages(t *testing.T) {
	from := time.Date(2018, time.May, 1, 0, 0, 0, 0, time.Local)
	to := time.Now()
	opt := &SearchMessagesOptions{
		TopicIDs:       []int{topicID},
		HasAttachments: false,
		From:           &from,
		To:             &to,
	}
	result, resp, err := clientV2.Messages.SearchMessages(context.Background(), spaceKey, "test", opt)
	t.Logf("result: %v", result)
	test(t, result, resp, err)
}
