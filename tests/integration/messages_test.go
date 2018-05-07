package tests

import (
	"context"
	"testing"
)

func Test_V1_Messages_GetMessage_should_get_a_message(t *testing.T) {
	result, resp, err := clientV1.Messages.GetMessage(context.Background(), topicId, postId)
	test(t, result, resp, err)
}

func Test_V1_Messages_PostMessage_should_post_a_message(t *testing.T) {
	result, resp, err := clientV1.Messages.PostMessage(context.Background(), topicId, "go-typetalk - Test_Messages_PostMessage_should_post_a_message", nil)
	test(t, result, resp, err)
}

func Test_V1_Messages_UpdateMessage_should_update_a_message(t *testing.T) {
	result, resp, err := clientV1.Messages.UpdateMessage(context.Background(), topicId, postId, "go-typetalk - Test_Messages_UpdateMessage_should_update_a_message")
	test(t, result, resp, err)
}

func Test_V1_Messages_GetMessage_should_get_a_message_using_Typetalk_Token(t *testing.T) {
	result, resp, err := clientUsingTypetalkTokenV1.Messages.GetMessage(context.Background(), topicId, postId)
	test(t, result, resp, err)
}

func Test_V1_Messages_PostMessage_should_post_a_message_using_Typetalk_Token(t *testing.T) {
	result, resp, err := clientUsingTypetalkTokenV1.Messages.PostMessage(context.Background(), topicId, "go-typetalk - Test_Messages_PostMessage_should_post_a_message_using_Typetalk_Token", nil)
	test(t, result, resp, err)
}

func Test_V2_Messages_SearchMessages_should_some_messages(t *testing.T) {
	result, resp, err := clientV2.Messages.SearchMessages(context.Background(), spaceKey, "test", nil)
	test(t, result, resp, err)
}
