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

func Test_MessagesService_SearchMessages_should_get_some_posts(t *testing.T) {
	setup()
	defer teardown()
	from := time.Date(2018, time.May, 1, 0, 0, 0, 0, time.Local)
	to := time.Date(2018, time.May, 3, 0, 0, 0, 0, time.Local)
	b, _ := ioutil.ReadFile(fixturesPath + "search-messages.json")
	mux.HandleFunc("/search/posts",
		func(w http.ResponseWriter, r *http.Request) {
			TestMethod(t, r, "GET")
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
