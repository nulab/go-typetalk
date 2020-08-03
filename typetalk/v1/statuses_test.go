package v1

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

func Test_StatusesService_SaveUserStatus_should_save_a_user_status(t *testing.T) {
	setup()
	defer teardown()
	spaceKey := "foo"
	b, _ := ioutil.ReadFile(fixturesPath + "save-user-status.json")
	mux.HandleFunc(fmt.Sprintf("/spaces/%v/userStatuses", spaceKey), func(w http.ResponseWriter, r *http.Request) {
		TestMethod(t, r, http.MethodPost)
		TestFormValues(t, r, Values{
			"emoji":                  ":musical_note:",
			"message":                "hello",
			"clearAt":                "2019-09-06T10:49:00Z",
			"isNotificationDisabled": "false",
		})
		fmt.Fprint(w, string(b))
	})

	result, _, err := client.Statuses.SaveUserStatus(context.Background(), spaceKey, ":musical_note:", &SaveUserStatusOptions{
		Message:                stringPtr("hello"),
		ClearAt:                stringPtr("2019-09-06T10:49:00Z"),
		IsNotificationDisabled: boolPrt(false),
	})
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	want := &SaveUserStatusResult{}
	json.Unmarshal(b, want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("returned content: got %v, want %v", result, want)
	}
}

func stringPtr(s string) *string {
	return &s
}

func boolPrt(b bool) *bool {
	return &b
}
