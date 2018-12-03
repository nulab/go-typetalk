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

func Test_MentionsService_ReadMention_should_read_a_mention(t *testing.T) {
	setup()
	defer teardown()
	mentionID := 1
	b, _ := ioutil.ReadFile(fixturesPath + "read-mention.json")
	mux.HandleFunc(fmt.Sprintf("/mentions/%d", mentionID), func(w http.ResponseWriter, r *http.Request) {
		TestMethod(t, r, http.MethodPut)
		fmt.Fprint(w, string(b))
	})

	result, _, err := client.Mentions.ReadMention(context.Background(), mentionID)
	if err != nil {
		t.Errorf("returned error: %v", err)
	}
	var want *struct {
		Mention *Mention `json:"mention"`
	}
	json.Unmarshal(b, &want)
	if !reflect.DeepEqual(result, want.Mention) {
		t.Errorf("returned content: got  %v, want %v", result.ID, want.Mention.ID)
	}
}

func Test_MentionsService_GetMentionList_should_get_some_mentions(t *testing.T) {
	setup()
	defer teardown()
	b, _ := ioutil.ReadFile(fixturesPath + "get-mention-list.json")
	mux.HandleFunc("/mentions", func(w http.ResponseWriter, r *http.Request) {
		TestMethod(t, r, http.MethodGet)
		TestQueryValues(t, r, Values{
			"from":   10,
			"unread": true,
		})
		fmt.Fprint(w, string(b))
	})

	result, _, err := client.Mentions.GetMentionList(context.Background(), &GetMentionListOptions{10, true})
	if err != nil {
		t.Errorf("returned error: %v", err)
	}
	var want *struct {
		Mentions []*Mention `json:"mentions"`
	}
	json.Unmarshal(b, &want)
	if !reflect.DeepEqual(result, want.Mentions) {
		t.Errorf("returned content: got  %v, want %v", result, want.Mentions)
	}
}
