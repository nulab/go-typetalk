package v1

import (
	"context"
	"fmt"
	"time"

	. "github.com/nulab/go-typetalk/typetalk/internal"
	. "github.com/nulab/go-typetalk/typetalk/shared"
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

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/post-message
func (s *MessagesService) PostMessage(ctx context.Context, topicId int, message string, opt *PostMessageOptions) (*PostedMessageResult, *Response, error) {
	u := fmt.Sprintf("topics/%v", topicId)
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

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/update-message
func (s *MessagesService) UpdateMessage(ctx context.Context, topicId, postId int, message string) (*UpdatedMessageResult, *Response, error) {
	u := fmt.Sprintf("topics/%d/posts/%d", topicId, postId)
	var result *UpdatedMessageResult
	resp, err := s.client.Put(ctx, u, &updateMessageOptions{Message: message}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/delete-message
func (s *MessagesService) DeleteMessage(ctx context.Context, topicId, postId int) (*Post, *Response, error) {
	u := fmt.Sprintf("topics/%d/posts/%d", topicId, postId)
	var result *Post
	resp, err := s.client.Delete(ctx, u, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/get-message
func (s *MessagesService) GetMessage(ctx context.Context, topicId, postId int) (*Message, *Response, error) {
	u := fmt.Sprintf("topics/%d/posts/%d", topicId, postId)
	var result *Message
	resp, err := s.client.Get(ctx, u, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// Typetalk API docs: https://developer.nulab-inc.com/ja/docs/typetalk/api/1/favorite-topic
func (s *MessagesService) LikeMessage(ctx context.Context, topicId, postId int) (*LikedMessageResult, *Response, error) {
	u := fmt.Sprintf("topics/%d/posts/%d/like", topicId, postId)
	var result *LikedMessageResult
	resp, err := s.client.Post(ctx, u, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// Typetalk API docs: https://developer.nulab-inc.com/ja/docs/typetalk/api/1/unfavorite-topic
func (s *MessagesService) UnlikeMessage(ctx context.Context, topicId, postId int) (*Like, *Response, error) {
	u := fmt.Sprintf("topics/%d/posts/%d/like", topicId, postId)
	var result *struct {
		Like Like `json:"like"`
	}
	resp, err := s.client.Delete(ctx, u, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result.Like, resp, nil
}

// Typetalk API docs: https://developer.nulab-inc.com/ja/docs/typetalk/api/1/post-direct-message
func (s *MessagesService) PostDirectMessage(ctx context.Context, accountName, message string, opt *PostMessageOptions) (*PostedMessageResult, *Response, error) {
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

// Deprecated: Use GetDirectMessages in github.com/nulab/go-typetalk/typetalk/v2
func (s *MessagesService) GetDirectMessages(ctx context.Context, accountName string, opt *GetMessagesOptions) (*DirectMessages, *Response, error) {
	u, err := AddQueries(fmt.Sprintf("messages/@%s", accountName), opt)
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

// Deprecated: Use GetMyDirectMessageTopics in github.com/nulab/go-typetalk/typetalk/v2
func (s *MessagesService) GetMyDirectMessageTopics(ctx context.Context) ([]*DirectMessageTopic, *Response, error) {
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
