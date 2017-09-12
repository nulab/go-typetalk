package tests

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"testing"

	"github.com/nulab/go-typetalk/typetalk"
	"golang.org/x/oauth2"
)

var (
	client                   *typetalk.Client
	clientUsingTypetalkToken *typetalk.Client
	topicId                  int
	postId                   int
)

type AccessToken struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

func init() {
	clientId := os.Getenv("TT_CLIENT_ID")
	clientSecret := os.Getenv("TT_CLIENT_SECRET")
	if v, err := strconv.Atoi(os.Getenv("TT_TOPIC_ID")); err == nil {
		topicId = v
	}
	if v, err := strconv.Atoi(os.Getenv("TT_POST_ID")); err == nil {
		postId = v
	}
	if clientId == "" || clientSecret == "" {
		print("!!! Integration test using OAuth2 requires client_id and client_secret. !!!\n\n")
		client = typetalk.NewClient(nil)
	} else {
		form := url.Values{}
		form.Add("client_id", clientId)
		form.Add("client_secret", clientSecret)
		form.Add("grant_type", "client_credentials")
		form.Add("scope", "topic.read,topic.post,topic.write,topic.delete,my")
		resp, err := http.PostForm("https://typetalk.com/oauth2/access_token", form)
		if err != nil {
			print("Client Credential request returned error")
		}
		if resp == nil {
			print("Client Credential request returned nil response")
		}
		v := &AccessToken{}
		json.NewDecoder(resp.Body).Decode(v)

		tc := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: v.AccessToken},
		))
		client = typetalk.NewClient(tc)
	}

	typetalkToken := os.Getenv("TT_TOKEN")
	if typetalkToken == "" {
		print("!!! Integration test using Typetalk Token requires Typetalk Token. !!!\n\n")
		clientUsingTypetalkToken = typetalk.NewClient(nil)
	} else {
		clientUsingTypetalkToken = typetalk.NewClient(nil).SetTypetalkToken(typetalkToken)
	}
}

func test(t *testing.T, result interface{}, resp *typetalk.Response, err error) {
	if err != nil {
		t.Fatalf("Returned error: %v", err)
	}
	if resp == nil {
		t.Fatal("Returned nil response")
	}
	if result == nil {
		t.Error("Returned nil result")
	}
}
