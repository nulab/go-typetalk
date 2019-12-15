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

func Test_NotificationsService_GetNotificationList_should_get_some_notifications(t *testing.T) {
	setup()
	defer teardown()
	b, _ := ioutil.ReadFile(fixturesPath + "get-notification-list.json")
	mux.HandleFunc("/notifications",
		func(w http.ResponseWriter, r *http.Request) {
			TestMethod(t, r, http.MethodGet)
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Notifications.GetNotificationList(context.Background())
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	want := &NotificationList{}
	json.Unmarshal(b, want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}

func Test_NotificationsService_GetNotificationCount_should_get_notification_count(t *testing.T) {
	setup()
	defer teardown()
	b, _ := ioutil.ReadFile(fixturesPath + "get-notification-count.json")
	mux.HandleFunc("/notifications/status",
		func(w http.ResponseWriter, r *http.Request) {
			TestMethod(t, r, http.MethodGet)
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Notifications.GetNotificationCount(context.Background())
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	want := &NotificationCount{}
	json.Unmarshal(b, want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}

func Test_NotificationsService_ReadNotification_should_read_notification(t *testing.T) {
	setup()
	defer teardown()
	b, _ := ioutil.ReadFile(fixturesPath + "read-notification.json")
	mux.HandleFunc("/notifications",
		func(w http.ResponseWriter, r *http.Request) {
			TestMethod(t, r, http.MethodPut)
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Notifications.ReadNotification(context.Background())
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	var want *struct {
		Access *Access `json:"access"`
	}
	json.Unmarshal(b, &want)
	if !reflect.DeepEqual(result, want.Access) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want.Access)
	}
}
