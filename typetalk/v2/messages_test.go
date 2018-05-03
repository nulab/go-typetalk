package v2

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

func Test_MessagesService_SearchMessages_should_get_some_posts(t *testing.T) {
	setup()
	defer teardown()
	b, _ := ioutil.ReadFile(fixturesPath + "search-messages.json")
	mux.HandleFunc("/search/posts",
		func(w http.ResponseWriter, r *http.Request) {
			TestMethod(t, r, "GET")
			TestQueryValues(t, r, Values{
				"spaceKey":       "qwerty",
				"q":              "hello",
				"hasAttachments": true,
			})
			fmt.Fprint(w, string(b))
		})

	query := &SearchMessagesOptions{TopicIDs: []int{}, HasAttachments: true, AccountIDs: []int{}}
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
