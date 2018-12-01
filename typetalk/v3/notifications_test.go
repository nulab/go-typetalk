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

func Test_NotificationsService_ReadNotification_should_read_notification(t *testing.T) {
	setup()
	defer teardown()
	b, _ := ioutil.ReadFile(fixturesPath + "read-notification.json")
	mux.HandleFunc("/notifications",
		func(w http.ResponseWriter, r *http.Request) {
			TestMethod(t, r, http.MethodPut)
			TestFormValues(t, r, Values{
				"spaceKey": "qwerty",
			})
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Notifications.ReadNotification(context.Background(), "qwerty")
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	var want *ReadNotificationResult
	json.Unmarshal(b, &want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}
