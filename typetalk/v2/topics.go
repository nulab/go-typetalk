package v2

import (
	"context"

	"time"

	. "github.com/nulab/go-typetalk/typetalk/internal"
	. "github.com/nulab/go-typetalk/typetalk/shared"
)

type TopicsService service

type Topic struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	Suggestion      string    `json:"suggestion"`
	IsDirectMessage bool      `json:"isDirectMessage"`
	LastPostedAt    time.Time `json:"lastPostedAt"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

type Unread struct {
	TopicID          int  `json:"topicId"`
	PostID           int  `json:"postId"`
	Count            int  `json:"count"`
	IsOverCountLimit bool `json:"isOverCountLimit"`
}

type FavoriteTopicWithUnread struct {
	Topic    Topic  `json:"topic"`
	Favorite bool   `json:"favorite"`
	Unread   Unread `json:"unread"`
}

type getMyTopicsOptions struct {
	SpaceKey string `json:"spaceKey"`
}

type DirectMessageTopic struct {
	Topic         *Topic         `json:"topic"`
	Unread        *Unread        `json:"unread"`
	DirectMessage *DirectMessage `json:"directMessage"`
}

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/2/get-topics/
func (s *TopicsService) GetMyTopics(ctx context.Context, spaceKey string) ([]*FavoriteTopicWithUnread, *Response, error) {
	u, err := AddQueries("topics", &getMyTopicsOptions{spaceKey})
	if err != nil {
		return nil, nil, err
	}
	var result *struct {
		Topics []*FavoriteTopicWithUnread `json:"topics"`
	}
	if resp, err := s.client.Get(ctx, u, &result); err != nil {
		return nil, resp, err
	} else {
		return result.Topics, resp, nil
	}
}

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/2/get-dm-topics
func (s *MessagesService) GetMyDirectMessageTopics(ctx context.Context, spaceKey string) ([]*DirectMessageTopic, *Response, error) {
	u, err := AddQueries("messages", &getMyTopicsOptions{spaceKey})
	if err != nil {
		return nil, nil, err
	}
	var result *struct {
		DirectMessageTopics []*DirectMessageTopic `json:"topics"`
	}
	if resp, err := s.client.Get(ctx, u, &result); err != nil {
		return nil, resp, err
	} else {
		return result.DirectMessageTopics, resp, nil
	}
}
