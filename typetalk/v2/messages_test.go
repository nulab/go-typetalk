package v2

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"time"

	. "github.com/nulab/go-typetalk/typetalk/internal"
)

func Test_MessagesService_GetDirectMessages_should_get_some_direct_messages(t *testing.T) {
	setup()
	defer teardown()
	spaceKey := "qwerty"
	accountName := "test"
	b, _ := ioutil.ReadFile(fixturesPath + "get-direct-messages.json")
	mux.HandleFunc(fmt.Sprintf("/spaces/%s/messages/@%s", spaceKey, accountName),
		func(w http.ResponseWriter, r *http.Request) {
			TestMethod(t, r, http.MethodGet)
			TestQueryValues(t, r, Values{
				"count":     10,
				"from":      1,
				"direction": "backward",
			})
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Messages.GetDirectMessages(context.Background(), spaceKey, accountName, &GetMessagesOptions{10, 1, "backward"})
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	want := &DirectMessages{}
	json.Unmarshal(b, want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}

func Test_MessagesService_PostDirectMessage_should_post_dairect_message(t *testing.T) {
	setup()
	defer teardown()
	spaceKey := "qwerty"
	accountName := "test"
	b, _ := ioutil.ReadFile(fixturesPath + "post-direct-message.json")
	mux.HandleFunc(fmt.Sprintf("/spaces/%s/messages/@%s", spaceKey, accountName),
		func(w http.ResponseWriter, r *http.Request) {
			TestMethod(t, r, "POST")
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

	result, _, err := client.Messages.PostDirectMessage(context.Background(), spaceKey, accountName, "hello", &PostMessageOptions{
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

func Test_MessagesService_SearchMessages_should_get_some_posts(t *testing.T) {
	setup()
	defer teardown()
	from := time.Date(2018, time.May, 1, 0, 0, 0, 0, time.Local)
	to := time.Date(2018, time.May, 3, 0, 0, 0, 0, time.Local)
	b, _ := ioutil.ReadFile(fixturesPath + "search-messages.json")
	mux.HandleFunc("/search/posts",
		func(w http.ResponseWriter, r *http.Request) {
			TestMethod(t, r, http.MethodGet)
			TestQueryValues(t, r, Values{
				"spaceKey":       "qwerty",
				"q":              "hello",
				"hasAttachments": true,
				"from":           from.Format(time.RFC3339),
				"to":             to.Format(time.RFC3339),
			})
			fmt.Fprint(w, string(b))
		})

	query := &SearchMessagesOptions{TopicIDs: []int{}, HasAttachments: true, AccountIDs: []int{}, From: &from, To: &to}
	result, _, err := client.Messages.SearchMessages(context.Background(), "qwerty", "hello", query)
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	want := &SearchMessagesResult{}
	json.Unmarshal(b, want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}
