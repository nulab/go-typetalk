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

func Test_TopicsService_CreateTopic_should_create_a_topic(t *testing.T) {
	setup()
	defer teardown()
	b, _ := ioutil.ReadFile(fixturesPath + "create-topic.json")
	mux.HandleFunc("/topics",
		func(w http.ResponseWriter, r *http.Request) {
			TestMethod(t, r, "POST")
			TestFormValues(t, r, Values{
				"name":             "nulab",
				"spaceKey":         "balun",
				"addAccountIds[0]": 1,
				"addAccountIds[1]": 2,
				"addGroupIds[0]":   1,
				"addGroupIds[1]":   2,
				"addGroupIds[2]":   3,
			})
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Topics.CreateTopic(context.Background(), &CreateTopicOptions{"nulab", "balun", []int{1, 2}, []int{1, 2, 3}})
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	want := &TopicDetails{}
	json.Unmarshal(b, want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}

func Test_TopicsService_UpdateTopic_should_update_a_topic(t *testing.T) {
	setup()
	defer teardown()
	topicId := 1
	b, _ := ioutil.ReadFile(fixturesPath + "update-topic.json")
	mux.HandleFunc(fmt.Sprintf("/topics/%d", topicId),
		func(w http.ResponseWriter, r *http.Request) {
			TestMethod(t, r, "PUT")
			TestFormValues(t, r, Values{
				"name":        "nulab",
				"description": "This is a test topic.",
			})
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Topics.UpdateTopic(context.Background(), topicId, &UpdateTopicOptions{"nulab", "This is a test topic."})
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	want := &TopicDetails{}
	json.Unmarshal(b, want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}

func Test_TopicsService_DeleteTopic_should_delete_a_topic(t *testing.T) {
	setup()
	defer teardown()
	topicId := 1
	b, _ := ioutil.ReadFile(fixturesPath + "delete-topic.json")
	mux.HandleFunc(fmt.Sprintf("/topics/%d", topicId),
		func(w http.ResponseWriter, r *http.Request) {
			TestMethod(t, r, "DELETE")
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Topics.DeleteTopic(context.Background(), topicId)
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	want := &Topic{}
	json.Unmarshal(b, want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}

func Test_TopicsService_GetTopicDetails_should_get_a_topic_details(t *testing.T) {
	setup()
	defer teardown()
	topicId := 1
	b, _ := ioutil.ReadFile(fixturesPath + "get-topic-details.json")
	mux.HandleFunc(fmt.Sprintf("/topics/%d", topicId),
		func(w http.ResponseWriter, r *http.Request) {
			TestMethod(t, r, "GET")
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Topics.GetTopicDetails(context.Background(), topicId)
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	want := &TopicDetails{}
	json.Unmarshal(b, want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}

func Test_TopicsService_GetTopicMessages_should_get_some_topic_messages(t *testing.T) {
	setup()
	defer teardown()
	topicId := 1
	b, _ := ioutil.ReadFile(fixturesPath + "get-topic-messages.json")
	mux.HandleFunc(fmt.Sprintf("/topics/%d", topicId),
		func(w http.ResponseWriter, r *http.Request) {
			TestMethod(t, r, "GET")
			TestQueryValues(t, r, Values{
				"count":     10,
				"from":      3,
				"direction": "backward",
			})
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Topics.GetTopicMessages(context.Background(), topicId, &GetTopicMessagesOptions{10, 3, "backward"})
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	want := &TopicMessages{}
	json.Unmarshal(b, want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}

func Test_TopicsService_UpdateTopicMembers_should_add_some_topic_members(t *testing.T) {
	setup()
	defer teardown()
	topicId := 1
	b, _ := ioutil.ReadFile(fixturesPath + "update-topic-members.json")
	mux.HandleFunc(fmt.Sprintf("/topics/%d/members/update", topicId),
		func(w http.ResponseWriter, r *http.Request) {
			TestMethod(t, r, "POST")
			TestFormValues(t, r, Values{
				"addAccountIds[0]":                        1,
				"addGroupIds[0]":                          1,
				"invitations[0].email":                    "example1@nulab-inc.com",
				"invitations[0].role":                     "Admin",
				"removeAccounts[0].id":                    4,
				"removeAccounts[0].cancelSpaceInvitation": true,
				"removeGroupIds[0]":                       true,
			})
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Topics.UpdateTopicMembers(context.Background(), topicId, &UpdateTopicMembersOptions{
		AddAccountIds: []int{1},
		AddGroupIds:   []int{1},
		InvitationsEmail: []string{
			"example1@nulab-inc.com",
		},
		InvitationsRole:                     []string{"Admin"},
		RemoveAccountsId:                    []int{4},
		RemoveAccountsCancelSpaceInvitation: []bool{true},
		RemoveGroupIds:                      []bool{true},
	})
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	want := &TopicDetails{}
	json.Unmarshal(b, want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}

func Test_TopicsService_FavoriteTopic_should_favorite_topic(t *testing.T) {
	setup()
	defer teardown()
	topicId := 1
	b, _ := ioutil.ReadFile(fixturesPath + "favorite-topic.json")
	mux.HandleFunc(fmt.Sprintf("/topics/%d/favorite", topicId),
		func(w http.ResponseWriter, r *http.Request) {
			TestMethod(t, r, "POST")
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Topics.FavoriteTopic(context.Background(), topicId)
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	want := &FavoriteTopic{}
	json.Unmarshal(b, want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}

func Test_TopicsService_UnfavoriteTopic_should_unfavorite_topic(t *testing.T) {
	setup()
	defer teardown()
	topicId := 1
	b, _ := ioutil.ReadFile(fixturesPath + "unfavorite-topic.json")
	mux.HandleFunc(fmt.Sprintf("/topics/%d/favorite", topicId),
		func(w http.ResponseWriter, r *http.Request) {
			TestMethod(t, r, "DELETE")
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Topics.UnfavoriteTopic(context.Background(), topicId)
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	want := &FavoriteTopic{}
	json.Unmarshal(b, want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}

func Test_TopicsService_ReadMessagesInTopic_should_read_some_messages_in_topic(t *testing.T) {
	setup()
	defer teardown()
	topicId := 1
	postId := 1
	b, _ := ioutil.ReadFile(fixturesPath + "read-messages-in-topic.json")
	mux.HandleFunc("/bookmarks",
		func(w http.ResponseWriter, r *http.Request) {
			TestMethod(t, r, "PUT")
			TestQueryValues(t, r, Values{
				"topicId": 1,
				"postId":  1,
			})
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Topics.ReadMessagesInTopic(context.Background(), topicId, postId)
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	var want *struct {
		Unread *Unread `json:"unread"`
	}
	json.Unmarshal(b, &want)
	if !reflect.DeepEqual(result, want.Unread) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want.Unread)
	}
}

func Test_TopicsService_GetMyTopics_should_get_some_topics(t *testing.T) {
	setup()
	defer teardown()
	b, _ := ioutil.ReadFile(fixturesPath + "get-my-topics.json")
	mux.HandleFunc("/topics",
		func(w http.ResponseWriter, r *http.Request) {
			TestMethod(t, r, "GET")
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Topics.GetMyTopics(context.Background())
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	var want *struct {
		Topics []*FavoriteTopicWithUnread `json:"topics"`
	}
	json.Unmarshal(b, &want)
	if !reflect.DeepEqual(result, want.Topics) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want.Topics)
	}
}
