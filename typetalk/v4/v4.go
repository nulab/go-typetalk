package v4

import (
	"net/http"
	"net/url"

	"github.com/nulab/go-typetalk/v3/typetalk/internal"
)

const (
	APIVersion = "v4"
)

type service struct {
	client *internal.ClientCore
}

type Client struct {
	client *internal.ClientCore

	Accounts *AccountsService
}

func (c *Client) SetTypetalkToken(token string) *Client {
	c.client.TypetalkToken = token
	return c
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(internal.DefaultBaseURL + APIVersion + "/")

	c := &Client{client: &internal.ClientCore{Client: httpClient, BaseURL: baseURL, UserAgent: internal.UserAgent}}

	common := &service{client: c.client}

	c.Accounts = (*AccountsService)(common)

	return c
}
