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
	var want *NotificationCount
	json.Unmarshal(b, &want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}

func Test_NotificationsService_GetNotificationCount_errorResponse(t *testing.T) {
	_, _, err := client.Notifications.GetNotificationCount(context.Background())
	if err == nil {
		t.Errorf("Expected error to be returned")
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
	var want *ReadNotificationResult
	json.Unmarshal(b, &want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}

func Test_NotificationsService_ReadNotification_errorResponse(t *testing.T) {
	_, _, err := client.Notifications.ReadNotification(context.Background())
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}
