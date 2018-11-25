package v2

import (
	"context"
	"fmt"

	"time"

	. "github.com/nulab/go-typetalk/typetalk/internal"
	. "github.com/nulab/go-typetalk/typetalk/shared"
)

type MessagesService service

type DirectMessages struct {
	Topic         *Topic         `json:"topic"`
	DirectMessage *DirectMessage `json:"directMessage"`
	Bookmark      *Bookmark      `json:"bookmark"`
	Posts         []*Post        `json:"posts"`
	HasNext       bool           `json:"hasNext"`
}

type Bookmark struct {
	PostID    int       `json:"postId"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type SearchMessagesResult struct {
	Count     int     `json:"count"`
	Posts     []*Post `json:"posts"`
	IsLimited bool    `json:"isLimited"`
}

type SearchMessagesOptions struct {
	TopicIDs       []int      `json:"topicIds,omitempty"`
	HasAttachments bool       `json:"hasAttachments,omitempty"`
	AccountIDs     []int      `json:"accountIds,omitempty"`
	From           *time.Time `json:"from,omitempty"`
	To             *time.Time `json:"to,omitempty"`
}

type searchMessagesOptions struct {
	*SearchMessagesOptions
	SpaceKey string `json:"spaceKey"`
	Q        string `json:"q"`
}

type GetMessagesOptions struct {
	Count     int    `json:"count,omitempty"`
	From      int    `json:"from,omitempty"`
	Direction string `json:"direction,omitempty"`
}

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/2/get-direct-messages
func (s *MessagesService) GetDirectMessages(ctx context.Context, spaceKey, accountName string, opt *GetMessagesOptions) (*DirectMessages, *Response, error) {
	u, err := AddQueries(fmt.Sprintf("spaces/%s/messages/@%s", spaceKey, accountName), opt)
	if err != nil {
		return nil, nil, err
	}
	var result *DirectMessages
	resp, err := s.client.Get(ctx, u, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/2/search-messages/
func (s *MessagesService) SearchMessages(ctx context.Context, spaceKey, q string, opt *SearchMessagesOptions) (*SearchMessagesResult, *Response, error) {
	u, err := AddQueries("search/posts", &searchMessagesOptions{SearchMessagesOptions: opt, SpaceKey: spaceKey, Q: q})
	if err != nil {
		return nil, nil, err
	}
	var result *SearchMessagesResult
	resp, err := s.client.Get(ctx, u, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}
