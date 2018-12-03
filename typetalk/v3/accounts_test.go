package v3

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

func Test_AccountsService_GetMyFriends_should_get_some_friends(t *testing.T) {
	setup()
	defer teardown()
	b, _ := ioutil.ReadFile(fixturesPath + "get-my-friends.json")
	mux.HandleFunc("/search/friends", func(w http.ResponseWriter, r *http.Request) {
		TestMethod(t, r, http.MethodGet)
		TestQueryValues(t, r, Values{
			"spaceKey": "qwerty",
			"q":        "hello",
			"offset":   10,
			"count":    2,
		})
		fmt.Fprint(w, string(b))
	})

	result, _, err := client.Accounts.GetMyFriends(context.Background(), "qwerty", "hello", &GetMyFriendsOptions{
		Offset: 10,
		Count:  2,
	})
	if err != nil {
		t.Errorf("returned error: %v", err)
	}

	var want *struct {
		Accounts []*Account `json:"accounts"`
	}
	json.Unmarshal(b, &want)
	if !reflect.DeepEqual(result, want.Accounts) {
		t.Errorf("returned content: got %v, want %v", result, want)
	}
}
