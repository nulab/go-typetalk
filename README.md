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

### Import

Use v1:
``` go
import "github.com/nulab/go-typetalk/typetalk/v1"
```
Use v2:
``` go
import "github.com/nulab/go-typetalk/typetalk/v2"
```
Use v3:
``` go
import "github.com/nulab/go-typetalk/typetalk/v3"
```

### Access APIs using Typetalk Token

``` go
package main

import (
	"context"

	"github.com/nulab/go-typetalk/typetalk/v1"
)

func main() {
	client := v1.NewClient(nil)
	client.SetTypetalkToken("yourTypetalkToken")
	ctx := context.Background()
	topicId := 1
	message := "Hello"
	profile, resp, err := client.Messages.PostMessage(ctx, topicId, message, nil)
}
```

### Access APIs using OAuth2 Access Token

``` go
package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/nulab/go-typetalk/typetalk/v1"
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
	profile, resp, err := client.Accounts.GetMyProfile(context.Background())
}
```

## Bugs and Feedback

For bugs, questions and discussions please use the Github Issues.

## License

[MIT License](http://www.opensource.org/licenses/mit-license.php)
