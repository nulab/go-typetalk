package tests

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"testing"

	"github.com/nulab/go-typetalk/typetalk/shared"
	"github.com/nulab/go-typetalk/typetalk/v1"
	"github.com/nulab/go-typetalk/typetalk/v2"
	"github.com/nulab/go-typetalk/typetalk/v3"
	"golang.org/x/oauth2"
)

var (
	clientV1                   *v1.Client
	clientV2                   *v2.Client
	clientV3                   *v3.Client
	clientUsingTypetalkTokenV1 *v1.Client
	clientUsingTypetalkTokenV2 *v2.Client
	clientUsingTypetalkTokenV3 *v3.Client
	topicId                    int
	postId                     int
	spaceKey                   string
)

type AccessToken struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

func init() {
	spaceKey = os.Getenv("TT_SPACE_KEY")
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
		clientV1 = v1.NewClient(nil)
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
		clientV1 = v1.NewClient(tc)
		clientV2 = v2.NewClient(tc)
		clientV3 = v3.NewClient(tc)
	}

	clientUsingTypetalkTokenV1 = v1.NewClient(nil)
	clientUsingTypetalkTokenV2 = v2.NewClient(nil)
	clientUsingTypetalkTokenV3 = v3.NewClient(nil)

	typetalkToken := os.Getenv("TT_TOKEN")
	if typetalkToken != "" {
		clientUsingTypetalkTokenV1.SetTypetalkToken(typetalkToken)
		clientUsingTypetalkTokenV2.SetTypetalkToken(typetalkToken)
		clientUsingTypetalkTokenV3.SetTypetalkToken(typetalkToken)
	} else {
		print("!!! Integration test using Typetalk Token requires Typetalk Token. !!!\n\n")
	}
}

func test(t *testing.T, result interface{}, resp *shared.Response, err error) {
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
