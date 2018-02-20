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

func Test_LikesService_GetLikesReceive_should_get_Likes(t *testing.T) {
	setup()
	defer teardown()
	b, _ := ioutil.ReadFile("../testdata/get-likes-receive.json")
	mux.HandleFunc("/likes/receive",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			testQueryValues(t, r, values{
				"from": 1,
			})
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Likes.GetLikesReceive(context.Background(), &GetLikesOptions{1})
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	var want *struct {
		LikedPosts []*ReceiveLikedPost `json:"likedPosts"`
	}
	json.Unmarshal(b, &want)
	if !reflect.DeepEqual(result, want.LikedPosts) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}

func Test_LikesService_GetLikesGive_should_get_Likes(t *testing.T) {
	setup()
	defer teardown()
	b, _ := ioutil.ReadFile("../testdata/get-likes-give.json")
	mux.HandleFunc("/likes/give",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			testQueryValues(t, r, values{
				"from": 1,
			})
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Likes.GetLikesGive(context.Background(), &GetLikesOptions{1})
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	var want *struct {
		LikedPosts []*GiveLikedPost `json:"likedPosts"`
	}
	json.Unmarshal(b, &want)
	if !reflect.DeepEqual(result, want.LikedPosts) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}

func Test_LikesService_GetLikesDiscover_should_get_Likes(t *testing.T) {
	setup()
	defer teardown()
	b, _ := ioutil.ReadFile("../testdata/get-likes-discover.json")
	mux.HandleFunc("/likes/discover",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			testQueryValues(t, r, values{
				"from": 1,
			})
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Likes.GetLikesDiscover(context.Background(), &GetLikesOptions{1})
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	var want *struct {
		LikedPosts []*DiscoverLikedPost `json:"likedPosts"`
	}
	json.Unmarshal(b, &want)
	if !reflect.DeepEqual(result, want.LikedPosts) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}

func Test_LikesService_ReadReceivedLikes_should_get_Likes(t *testing.T) {
	setup()
	defer teardown()
	b, _ := ioutil.ReadFile("../testdata/read-received-likes.json")
	mux.HandleFunc("/likes/receive/bookmark/save",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "POST")
			testFormValues(t, r, values{
				"likeId": 1,
			})
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Likes.ReadReceivedLikes(context.Background(), 1)
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	var want *ReadReceivedLikesResult
	json.Unmarshal(b, want)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want)
	}
}
