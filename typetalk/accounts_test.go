package typetalk

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

func Test_AccountsService_GetMyProfile_should_get_a_profile(t *testing.T) {
	setup()
	defer teardown()
	b, _ := ioutil.ReadFile("../testdata/get-my-profile.json")
	mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, string(b))
	})

	result, _, err := client.Accounts.GetMyProfile(context.Background())
	if err != nil {
		t.Errorf("returned error: %v", err)
	}

	want := &Profile{}
	json.Unmarshal(b, want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("returned content: got %v, want %v", result, want)
	}
}

func Test_AccountsService_GetFriendProfile_should_get_a_profile(t *testing.T) {
	setup()
	defer teardown()
	b, _ := ioutil.ReadFile("../testdata/get-friend-profile.json")
	accountName := "ahorowitz"
	mux.HandleFunc(fmt.Sprintf("/profile/%s", accountName), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, string(b))
	})

	result, _, err := client.Accounts.GetFriendProfile(context.Background(), accountName)
	if err != nil {
		t.Errorf("returned error: %v", err)
	}

	want := &Profile{}
	json.Unmarshal(b, want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("returned content: got %v, want %v", result, want)
	}
}

func Test_AccountsService_GetMyFriends_should_get_some_friends(t *testing.T) {
	setup()
	defer teardown()
	b, _ := ioutil.ReadFile("../testdata/get-my-friends.json")
	mux.HandleFunc("/search/friends", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testQueryValues(t, r, values{
			"q":      "test",
			"offset": 10,
			"count":  2,
		})
		fmt.Fprint(w, string(b))
	})

	result, _, err := client.Accounts.GetMyFriends(context.Background(), &GetMyFriendsOptions{
		Q:      "test",
		Offset: 10,
		Count:  2,
	})
	if err != nil {
		t.Errorf("returned error: %v", err)
	}

	want := &Friends{}
	json.Unmarshal(b, want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("returned content: got %v, want %v", result, want)
	}
}

func Test_AccountsService_SearchAccounts_should_get_some_accounts(t *testing.T) {
	setup()
	defer teardown()
	b, _ := ioutil.ReadFile("../testdata/search-accounts.json")
	mux.HandleFunc("/search/accounts", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testQueryValues(t, r, values{
			"nameOrEmailAddress": "test",
		})
		fmt.Fprint(w, string(b))
	})

	result, _, err := client.Accounts.SearchAccounts(context.Background(), "test")
	if err != nil {
		t.Errorf("returned error: %v", err)
	}

	want := &Account{}
	json.Unmarshal(b, want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("returned content: got %v, want %v", result, want)
	}
}

func Test_AccountsService_GetOnlineStatus_should_get_some_accounts_online_status(t *testing.T) {
	setup()
	defer teardown()
	b, _ := ioutil.ReadFile("../testdata/get-online-status.json")
	mux.HandleFunc("/accounts/status", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testQueryValues(t, r, values{
			"accountIds[0]": 1,
			"accountIds[1]": 2,
			"accountIds[2]": 3,
		})
		fmt.Fprint(w, string(b))
	})

	result, _, err := client.Accounts.GetOnlineStatus(context.Background(), 1, 2, 3)
	if err != nil {
		t.Errorf("returned error: %v", err)
	}

	want := &OnlineStatus{}
	json.Unmarshal(b, want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("returned content: got %v, want %v", result, want)
	}
}
