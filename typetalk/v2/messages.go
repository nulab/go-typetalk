package v2

import (
	"context"
	"fmt"

	"time"

	"github.com/nulab/go-typetalk/typetalk/internal"
	"github.com/nulab/go-typetalk/typetalk/shared"
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

type PostMessageOptions struct {
	ReplyTo      int      `json:"replyTo,omitempty"`
	ShowLinkMeta bool     `json:"showLinkMeta,omitempty"`
	FileKeys     []string `json:"fileKeys[%d],omitempty"`
	TalkIds      []int    `json:"talkIds[%d],omitempty"`
	FileUrls     []string `json:"attachments[%d].fileUrl,omitempty"`
	FileNames    []string `json:"attachments[%d].fileName,omitempty"`
}

type postMessageOptions struct {
	*PostMessageOptions
	Message string `json:"message,omitempty"`
}

type PostedMessageResult struct {
	Space                  *Space         `json:"space"`
	Topic                  *Topic         `json:"topic"`
	Post                   *Post          `json:"post"`
	Mentions               []*Mention     `json:"mentions"`
	ExceedsAttachmentLimit bool           `json:"exceedsAttachmentLimit"`
	DirectMessage          *DirectMessage `json:"directMessage"`
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

// GetDirectMessages fetches direct messages.
//
// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/2/get-direct-messages
func (s *MessagesService) GetDirectMessages(ctx context.Context, spaceKey, accountName string, opt *GetMessagesOptions) (*DirectMessages, *shared.Response, error) {
	u, err := internal.AddQueries(fmt.Sprintf("spaces/%s/messages/@%s", spaceKey, accountName), opt)
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

// PostDirectMessage posts direct message.
//
// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/2/post-direct-message
func (s *MessagesService) PostDirectMessage(ctx context.Context, spaceKey, accountName, message string, opt *PostMessageOptions) (*PostedMessageResult, *shared.Response, error) {
	u := fmt.Sprintf("spaces/%s/messages/@%s", spaceKey, accountName)
	if opt == nil {
		opt = &PostMessageOptions{}
	}
	var result *PostedMessageResult
	resp, err := s.client.Post(ctx, u, &postMessageOptions{opt, message}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// SearchMessages searches messages.
//
// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/2/search-messages/
func (s *MessagesService) SearchMessages(ctx context.Context, spaceKey, q string, opt *SearchMessagesOptions) (*SearchMessagesResult, *shared.Response, error) {
	u, err := internal.AddQueries("search/posts", &searchMessagesOptions{SearchMessagesOptions: opt, SpaceKey: spaceKey, Q: q})
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
