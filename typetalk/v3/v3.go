package v3

import (
	"net/http"
	"net/url"

	"github.com/nulab/go-typetalk/typetalk/internal"
)

const (
	APIVersion = "v3"
)

type service struct {
	client *internal.ClientCore
}

type Client struct {
	client *internal.ClientCore

	Accounts      *AccountsService
	Notifications *NotificationsService
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
	c.Notifications = (*NotificationsService)(common)

	return c
}
