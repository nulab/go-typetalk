# go-typetalk [![Build Status](https://travis-ci.org/nulab/go-typetalk.svg?branch=master)](https://travis-ci.org/nulab/go-typetalk) [![Coverage Status](https://coveralls.io/repos/github/nulab/go-typetalk/badge.svg?branch=master)](https://coveralls.io/github/nulab/go-typetalk?branch=master)

go-typetalk is a GO client library for accessing the [Typetalk API](http://developer.nulab-inc.com/docs/typetalk).

## Prerequisite

To use this library, you must have a valid [client id and client secret](https://developer.nulab-inc.com/docs/typetalk/auth#oauth2) provided by Typetalk and register a new client application. Or you can use the [Typetalk Token](https://developer.nulab-inc.com/docs/typetalk/auth#tttoken).

## Installation

This package can be installed with the go get command:

```
$ go get github.com/nulab/go-typetalk
```

## Usage

### Access APIs using Typetalk Token

``` go
package main

import (
	"context"
	"github.com/nulab/go-typetalk/typetalk"
)

func main() {
	client := typetalk.NewClient(nil).SetTypetalkToken("yourTypetalkToken")
	ctx := context.Background()
	topicId, postId := 1, 2
	message := "Hello"
	profile, resp, err := client.Messages.PostMessage(ctx, topicId, postId, message, nil)
}
```

### Access APIs using OAuth2 Access Token

``` go
package main

import (
	"context"
	"encoding/json"
	"github.com/nulab/go-typetalk/typetalk"
	"golang.org/x/oauth2"
	"net/http"
	"net/url"
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
	resp, err := http.PostForm("https://typetalk.com/oauth2/access_token", form)
	if err != nil {
		print("Client Credential request returned error")
	}
	v := &AccessToken{}
	json.NewDecoder(resp.Body).Decode(v)

	tc := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: v.AccessToken},
	))

	client := typetalk.NewClient(tc)
	profile, resp, err := client.Accounts.GetMyProfile(context.Background())
}
```

## Bugs and Feedback

For bugs, questions and discussions please use the Github Issues.

## License

[MIT License](http://www.opensource.org/licenses/mit-license.php)
