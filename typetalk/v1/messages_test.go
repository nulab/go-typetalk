package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	. "github.com/nulab/go-typetalk/typetalk/internal"
)

func Test_MessagesService_PostMessage_should_post_a_message(t *testing.T) {
	setup()
	defer teardown()
	topicId := 1
	b, _ := ioutil.ReadFile(fixturesPath + "post-message.json")
	mux.HandleFunc(fmt.Sprintf("/topics/%v", topicId), func(w http.ResponseWriter, r *http.Request) {
		TestMethod(t, r, http.MethodPost)
		TestFormValues(t, r, Values{
			"message":                 "hello",
			"replyTo":                 2,
			"showLinkMeta":            true,
			"fileKeys[0]":             "key0",
			"fileKeys[1]":             "key1",
			"fileKeys[2]":             "key2",
			"talkIds[0]":              0,
			"talkIds[1]":              1,
			"talkIds[2]":              2,
			"attachments[0].fileUrl":  "Url0",
			"attachments[1].fileUrl":  "Url1",
			"attachments[2].fileUrl":  "Url2",
			"attachments[0].fileName": "Name0",
			"attachments[1].fileName": "Name1",
			"attachments[2].fileName": "Name2",
		})
		fmt.Fprint(w, string(b))
	})

	result, _, err := client.Messages.PostMessage(context.Background(), topicId, "hello", &PostMessageOptions{
		ReplyTo:      2,
		ShowLinkMeta: true,
		FileKeys:     []string{"key0", "key1", "key2"},
		TalkIds:      []int{0, 1, 2},
		FileUrls:     []string{"Url0", "Url1", "Url2"},
		FileNames:    []string{"Name0", "Name1", "Name2"},
	})
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	want := &PostedMessageResult{}
	json.Unmarshal(b, want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("returned content: got %v, want %v", result, want)
	}
}

func Test_MessagesService_UpdateMessage_should_update_a_message(t *testing.T) {
	setup()
	defer teardown()
	topicId := 1
	postId := 1
	message := "hello"
	b, _ := ioutil.ReadFile(fixturesPath + "update-message.json")
	mux.HandleFunc(fmt.Sprintf("/topics/%d/posts/%d", topicId, postId),
		func(w http.ResponseWriter, r *http.Request) {
			TestMethod(t, r, http.MethodPut)
			TestFormValues(t, r, Values{"message": message})
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Messages.UpdateMessage(context.Background(), topicId, postId, message)
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	want := &UpdatedMessageResult{}
	json.Unmarshal(b, want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}

func Test_MessagesService_DeleteMessage_should_delete_a_message(t *testing.T) {
	setup()
	defer teardown()
	topicId := 1
	postId := 1
	b, _ := ioutil.ReadFile(fixturesPath + "delete-message.json")
	mux.HandleFunc(fmt.Sprintf("/topics/%d/posts/%d", topicId, postId),
		func(w http.ResponseWriter, r *http.Request) {
			TestMethod(t, r, "DELETE")
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Messages.DeleteMessage(context.Background(), topicId, postId)
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	want := &Post{}
	json.Unmarshal(b, want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Rreturned result:\n result  %v,\n want %v", result, want)
	}
}

func Test_MessagesService_GetMessage_should_get_a_message(t *testing.T) {
	setup()
	defer teardown()
	topicId := 1
	postId := 1
	b, _ := ioutil.ReadFile(fixturesPath + "get-message.json")
	mux.HandleFunc(fmt.Sprintf("/topics/%d/posts/%d", topicId, postId),
		func(w http.ResponseWriter, r *http.Request) {
			TestMethod(t, r, http.MethodGet)
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Messages.GetMessage(context.Background(), topicId, postId)
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	want := &Message{}
	json.Unmarshal(b, want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}

func Test_MessagesService_LikeMessage_should_like_a_message(t *testing.T) {
	setup()
	defer teardown()
	topicId := 1
	postId := 1
	b, _ := ioutil.ReadFile(fixturesPath + "like-message.json")
	mux.HandleFunc(fmt.Sprintf("/topics/%d/posts/%d/like", topicId, postId),
		func(w http.ResponseWriter, r *http.Request) {
			TestMethod(t, r, http.MethodPost)
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Messages.LikeMessage(context.Background(), topicId, postId)
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	want := &LikedMessageResult{}
	json.Unmarshal(b, want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}

func Test_MessagesService_UnikeMessage_should_unlike_a_message(t *testing.T) {
	setup()
	defer teardown()
	topicId := 1
	postId := 1
	b, _ := ioutil.ReadFile(fixturesPath + "unlike-message.json")
	mux.HandleFunc(fmt.Sprintf("/topics/%d/posts/%d/like", topicId, postId),
		func(w http.ResponseWriter, r *http.Request) {
			TestMethod(t, r, "DELETE")
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Messages.UnlikeMessage(context.Background(), topicId, postId)
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	want := &Like{}
	json.Unmarshal(b, want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}

func Test_MessagesService_PostDirectMessage_should_post_dairect_message(t *testing.T) {
	setup()
	defer teardown()
	accountName := "test"
	b, _ := ioutil.ReadFile(fixturesPath + "post-direct-message.json")
	mux.HandleFunc(fmt.Sprintf("/messages/@%s", accountName),
		func(w http.ResponseWriter, r *http.Request) {
			TestMethod(t, r, http.MethodPost)
			TestFormValues(t, r, Values{
				"message":                 "hello",
				"replyTo":                 2,
				"showLinkMeta":            true,
				"fileKeys[0]":             "key0",
				"fileKeys[1]":             "key1",
				"fileKeys[2]":             "key2",
				"talkIds[0]":              0,
				"talkIds[1]":              1,
				"talkIds[2]":              2,
				"attachments[0].fileUrl":  "Url0",
				"attachments[1].fileUrl":  "Url1",
				"attachments[2].fileUrl":  "Url2",
				"attachments[0].fileName": "Name0",
				"attachments[1].fileName": "Name1",
				"attachments[2].fileName": "Name2",
			})
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Messages.PostDirectMessage(context.Background(), accountName, "hello", &PostMessageOptions{
		ReplyTo:      2,
		ShowLinkMeta: true,
		FileKeys:     []string{"key0", "key1", "key2"},
		TalkIds:      []int{0, 1, 2},
		FileUrls:     []string{"Url0", "Url1", "Url2"},
		FileNames:    []string{"Name0", "Name1", "Name2"},
	})
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	want := &PostedMessageResult{}
	json.Unmarshal(b, want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}

func Test_MessagesService_GetDirectMessages_should_get_some_direct_messages(t *testing.T) {
	setup()
	defer teardown()
	accountName := "test"
	b, _ := ioutil.ReadFile(fixturesPath + "get-direct-messages.json")
	mux.HandleFunc(fmt.Sprintf("/messages/@%s", accountName),
		func(w http.ResponseWriter, r *http.Request) {
			TestMethod(t, r, http.MethodGet)
			TestQueryValues(t, r, Values{
				"count":     10,
				"from":      1,
				"direction": "backward",
			})
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Messages.GetDirectMessages(context.Background(), accountName, &GetMessagesOptions{10, 1, "backward"})
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	want := &DirectMessages{}
	json.Unmarshal(b, want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}

func Test_MessagesService_GetMyDirectMessageTopics_should_get_some_topics_of_direct_message(t *testing.T) {
	setup()
	defer teardown()
	b, _ := ioutil.ReadFile(fixturesPath + "get-my-direct-message-topics.json")
	mux.HandleFunc("/messages",
		func(w http.ResponseWriter, r *http.Request) {
			TestMethod(t, r, http.MethodGet)
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Messages.GetMyDirectMessageTopics(context.Background())
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	var want *struct {
		Topics []*DirectMessageTopic `json:"topics"`
	}
	json.Unmarshal(b, &want)
	if !reflect.DeepEqual(result, want.Topics) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}
