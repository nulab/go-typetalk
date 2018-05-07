package v2

import (
	"context"

	"time"

	. "github.com/nulab/go-typetalk/typetalk/internal"
	. "github.com/nulab/go-typetalk/typetalk/shared"
)

type MessagesService service

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

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/2/search-messages/
func (s *MessagesService) SearchMessages(ctx context.Context, spaceKey, q string, opt *SearchMessagesOptions) (*SearchMessagesResult, *Response, error) {
	u, err := AddQueries("search/posts", &searchMessagesOptions{SearchMessagesOptions: opt, SpaceKey: spaceKey, Q: q})
	if err != nil {
		return nil, nil, err
	}
	var result *SearchMessagesResult
	if resp, err := s.client.Get(ctx, u, &result); err != nil {
		return nil, resp, err
	} else {
		return result, resp, nil
	}
}
