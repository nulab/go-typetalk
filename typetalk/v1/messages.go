package v1

import (
	"context"
	"fmt"
	"time"

	"github.com/nulab/go-typetalk/v3/typetalk/internal"
	"github.com/nulab/go-typetalk/typetalk/shared"
)

type MessagesService service

type Post struct {
	ID          int               `json:"id"`
	TopicID     int               `json:"topicId"`
	ReplyTo     int               `json:"replyTo"`
	Message     string            `json:"message"`
	Account     *Account          `json:"account"`
	Mention     *Mention          `json:"mention"`
	Attachments []*AttachmentFile `json:"attachments"`
	Likes       []*Like           `json:"likes"`
	Talks       []*Talk           `json:"talks"`
	Links       []interface{}     `json:"links"`
	CreatedAt   *time.Time        `json:"createdAt"`
	UpdatedAt   *time.Time        `json:"updatedAt"`
}

type Like struct {
	ID        int        `json:"id"`
	PostID    int        `json:"postId"`
	TopicID   int        `json:"topicId"`
	Comment   string     `json:"comment"`
	Account   *Account   `json:"account"`
	CreatedAt *time.Time `json:"createdAt"`
}

type PostedMessageResult struct {
	Space                  *Space         `json:"space"`
	Topic                  *Topic         `json:"topic"`
	Post                   *Post          `json:"post"`
	Mentions               []*Mention     `json:"mentions"`
	ExceedsAttachmentLimit bool           `json:"exceedsAttachmentLimit"`
	DirectMessage          *DirectMessage `json:"directMessage"`
}

type UpdatedMessageResult PostedMessageResult

type Message struct {
	MySpace                *Organization `json:"mySpace"`
	Team                   interface{}   `json:"team"`
	Topic                  *Topic        `json:"topic"`
	Post                   *Post         `json:"post"`
	Replies                []*Post       `json:"replies"`
	ExceedsAttachmentLimit bool          `json:"exceedsAttachmentLimit"`
}

type LikedMessageResult struct {
	Like          *Like          `json:"like"`
	Post          *Post          `json:"post"`
	Topic         *Topic         `json:"topic"`
	DirectMessage *DirectMessage `json:"directMessage"`
}

type DirectMessage AccountStatus

type DirectMessages struct {
	Topic         *Topic         `json:"topic"`
	DirectMessage *DirectMessage `json:"directMessage"`
	Bookmark      *Bookmark      `json:"bookmark"`
	Posts         []*Post        `json:"posts"`
	HasNext       bool           `json:"hasNext"`
}

type Unread struct {
	TopicID int `json:"topicId"`
	PostID  int `json:"postId"`
	Count   int `json:"count"`
}

type DirectMessageTopic struct {
	Topic         *Topic         `json:"topic"`
	Unread        *Unread        `json:"unread"`
	DirectMessage *DirectMessage `json:"directMessage"`
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

// PostMessage posts a message.
//
// Typetalk API docs: https://developer.nulab.com/docs/typetalk/api/1/post-message
func (s *MessagesService) PostMessage(ctx context.Context, topicID int, message string, opt *PostMessageOptions) (*PostedMessageResult, *shared.Response, error) {
	u := fmt.Sprintf("topics/%v", topicID)
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

type updateMessageOptions struct {
	Message string `json:"message"`
}

// UpdateMessage updates a message.
//
// Typetalk API docs: https://developer.nulab.com/docs/typetalk/api/1/update-message
func (s *MessagesService) UpdateMessage(ctx context.Context, topicID, postID int, message string) (*UpdatedMessageResult, *shared.Response, error) {
	u := fmt.Sprintf("topics/%d/posts/%d", topicID, postID)
	var result *UpdatedMessageResult
	resp, err := s.client.Put(ctx, u, &updateMessageOptions{Message: message}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// DeleteMessage deletes a message.
//
// Typetalk API docs: https://developer.nulab.com/docs/typetalk/api/1/delete-message
func (s *MessagesService) DeleteMessage(ctx context.Context, topicID, postID int) (*Post, *shared.Response, error) {
	u := fmt.Sprintf("topics/%d/posts/%d", topicID, postID)
	var result *Post
	resp, err := s.client.Delete(ctx, u, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// GetMessage gets a message.
//
// Typetalk API docs: https://developer.nulab.com/docs/typetalk/api/1/get-message
func (s *MessagesService) GetMessage(ctx context.Context, topicID, postID int) (*Message, *shared.Response, error) {
	u := fmt.Sprintf("topics/%d/posts/%d", topicID, postID)
	var result *Message
	resp, err := s.client.Get(ctx, u, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// LikeMessage marks a message as liked.
//
// Typetalk API docs: https://developer.nulab.com/ja/docs/typetalk/api/1/favorite-topic
func (s *MessagesService) LikeMessage(ctx context.Context, topicID, postID int) (*LikedMessageResult, *shared.Response, error) {
	u := fmt.Sprintf("topics/%d/posts/%d/like", topicID, postID)
	var result *LikedMessageResult
	resp, err := s.client.Post(ctx, u, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// UnlikeMessage marks a message as unliked.
//
// Typetalk API docs: https://developer.nulab.com/ja/docs/typetalk/api/1/unfavorite-topic
func (s *MessagesService) UnlikeMessage(ctx context.Context, topicID, postID int) (*Like, *shared.Response, error) {
	u := fmt.Sprintf("topics/%d/posts/%d/like", topicID, postID)
	var result *struct {
		Like Like `json:"like"`
	}
	resp, err := s.client.Delete(ctx, u, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result.Like, resp, nil
}

// PostDirectMessage posts direct message.
//
// Deprecated: Use PostDirectMessage in github.com/nulab/go-typetalk/typetalk/v2
func (s *MessagesService) PostDirectMessage(ctx context.Context, accountName, message string, opt *PostMessageOptions) (*PostedMessageResult, *shared.Response, error) {
	u := fmt.Sprintf("messages/@%s", accountName)
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

type GetMessagesOptions struct {
	Count     int    `json:"count,omitempty"`
	From      int    `json:"from,omitempty"`
	Direction string `json:"direction,omitempty"`
}

// GetDirectMessages fetches direct messages list.
//
// Deprecated: Use GetDirectMessages in github.com/nulab/go-typetalk/typetalk/v2
func (s *MessagesService) GetDirectMessages(ctx context.Context, accountName string, opt *GetMessagesOptions) (*DirectMessages, *shared.Response, error) {
	u, err := internal.AddQueries(fmt.Sprintf("messages/@%s", accountName), opt)
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

// GetMyDirectMessageTopics fetches direct message topics list.
//
// Deprecated: Use GetMyDirectMessageTopics in github.com/nulab/go-typetalk/typetalk/v2
func (s *MessagesService) GetMyDirectMessageTopics(ctx context.Context) ([]*DirectMessageTopic, *shared.Response, error) {
	u := "messages"
	var result *struct {
		Topics []*DirectMessageTopic `json:"topics"`
	}
	resp, err := s.client.Get(ctx, u, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Topics, resp, nil
}
