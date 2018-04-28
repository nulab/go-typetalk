package v1

import (
	"net/http"
	"net/url"

	"github.com/nulab/go-typetalk/typetalk/internal"
)

const (
	ApiVersion = "v1"
)

type service struct {
	client *internal.ClientCore
}

type Client struct {
	client *internal.ClientCore

	Accounts      *AccountsService
	Files         *FilesService
	Mentions      *MentionsService
	Messages      *MessagesService
	Notifications *NotificationsService
	Organizations *OrganizationsService
	Talks         *TalksService
	Topics        *TopicsService
	Likes         *LikesService
}

func (c *Client) SetTypetalkToken(token string) *Client {
	c.client.TypetalkToken = token
	return c
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(internal.DefaultBaseURL + ApiVersion + "/")

	c := &Client{client: &internal.ClientCore{Client: httpClient, BaseURL: baseURL, UserAgent: internal.UserAgent}}

	common := &service{client: c.client}

	c.Accounts = (*AccountsService)(common)
	c.Files = (*FilesService)(common)
	c.Mentions = (*MentionsService)(common)
	c.Messages = (*MessagesService)(common)
	c.Notifications = (*NotificationsService)(common)
	c.Organizations = (*OrganizationsService)(common)
	c.Talks = (*TalksService)(common)
	c.Topics = (*TopicsService)(common)
	c.Likes = (*LikesService)(common)

	return c
}
