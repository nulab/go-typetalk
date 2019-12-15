package v2

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	. "github.com/nulab/go-typetalk/v3/typetalk/internal"
)

func Test_MentionsService_GetMentionList_should_get_some_mentions(t *testing.T) {
	setup()
	defer teardown()
	b, _ := ioutil.ReadFile(fixturesPath + "get-mention-list.json")
	mux.HandleFunc("/mentions", func(w http.ResponseWriter, r *http.Request) {
		TestMethod(t, r, http.MethodGet)
		TestQueryValues(t, r, Values{
			"spaceKey": "qwerty",
			"from":     10,
			"unread":   true,
		})
		fmt.Fprint(w, string(b))
	})

	result, _, err := client.Mentions.GetMentionList(context.Background(), "qwerty", &GetMentionListOptions{10, true})
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
