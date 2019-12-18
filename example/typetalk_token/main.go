package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	v1 "github.com/nulab/go-typetalk/typetalk/v1"
	"golang.org/x/oauth2"
)

type AccessToken struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

func main() {
	form := url.Values{}
	form.Add("client_id", "yourClientId")
	form.Add("client_secret", "yourClientSecret")
	form.Add("grant_type", "client_credentials")
	form.Add("scope", "topic.read,topic.post,topic.write,topic.delete,my")
	oauth2resp, err := http.PostForm("https://typetalk.com/oauth2/access_token", form)
	if err != nil {
		print("Client Credential request returned error")
	}
	v := &AccessToken{}
	json.NewDecoder(oauth2resp.Body).Decode(v)

	tc := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: v.AccessToken},
	))

	client := v1.NewClient(tc)
	client.Accounts.GetMyProfile(context.Background())
}
