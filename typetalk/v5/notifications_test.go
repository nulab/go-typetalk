package v5

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
