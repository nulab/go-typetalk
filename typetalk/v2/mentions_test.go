package v2

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

func Test_MentionsService_GetMentionList_should_get_some_mentions(t *testing.T) {
	setup()
	defer teardown()
	b, _ := ioutil.ReadFile(fixturesPath + "get-mention-list.json")
	mux.HandleFunc("/mentions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testQueryValues(t, r, values{
			"spaceKey": "qwerty",
			"from":   10,
			"unread": true,
		})
		fmt.Fprint(w, string(b))
	})

	result, _, err := client.Mentions.GetMentionList(context.Background(), &GetMentionListOptions{"qwerty", 10, true})
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
