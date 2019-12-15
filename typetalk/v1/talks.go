package v1

import (
	"context"
	"fmt"
	"time"

	"github.com/nulab/go-typetalk/v3/typetalk/internal"
	"github.com/nulab/go-typetalk/v3/typetalk/shared"
)

type TalksService service

type Talk struct {
	ID         int         `json:"id"`
	TopicID    int         `json:"topicId"`
	Name       string      `json:"name"`
	Suggestion string      `json:"suggestion"`
	CreatedAt  *time.Time  `json:"createdAt"`
	UpdatedAt  *time.Time  `json:"updatedAt"`
	Backlog    interface{} `json:"backlog"`
}

type CreatedTalkResult struct {
	Topic   *Topic `json:"topic"`
	Talk    *Talk  `json:"talk"`
	PostIds []int  `json:"postIds"`
}

type UpdatedTalkResult struct {
	Topic *Topic `json:"topic"`
	Talk  *Talk  `json:"talk"`
}

type DeletedTalkResult CreatedTalkResult

type RemovedMessagesResult CreatedTalkResult

type MessagesInTalk struct {
	MySpace       *Organization  `json:"mySpace"`
	Topic         *Topic         `json:"topic"`
	DirectMessage *DirectMessage `json:"directMessage"`
	Talk          *Talk          `json:"talk"`
	Posts         []*Post        `json:"posts"`
	HasNext       bool           `json:"hasNext"`
}

type CreateTalkOptions struct {
	TalkName string `json:"talkName"`
	PostIds  []int  `json:"postIds[%d],omitempty"`
}

// CreateTalk creates a talk.
//
// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/create-talk
func (s *TalksService) CreateTalk(ctx context.Context, topicID int, talkName string, postIds ...int) (*CreatedTalkResult, *shared.Response, error) {
	u := fmt.Sprintf("topics/%d/talks", topicID)
	var result *CreatedTalkResult
	resp, err := s.client.Post(ctx, u, &CreateTalkOptions{talkName, postIds}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

type updateTalkOptions struct {
	TalkName string `json:"talkName,omitempty"`
}

// UpdateTalk updates a talk.
//
// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/update-talk
func (s *TalksService) UpdateTalk(ctx context.Context, topicID, talkID int, talkName string) (*UpdatedTalkResult, *shared.Response, error) {
	u, err := internal.AddQueries(fmt.Sprintf("topics/%d/talks/%d", topicID, talkID), &updateTalkOptions{talkName})
	if err != nil {
		return nil, nil, err
	}

	var result *UpdatedTalkResult
	resp, err := s.client.Put(ctx, u, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// DeleteTalk deletes a talk.
//
// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/delete-talk
func (s *TalksService) DeleteTalk(ctx context.Context, topicID, talkID int) (*DeletedTalkResult, *shared.Response, error) {
	u := fmt.Sprintf("topics/%d/talks/%d", topicID, talkID)
	var result *DeletedTalkResult
	resp, err := s.client.Delete(ctx, u, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// GetTalkList fetches talks list.
//
// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/get-talks
func (s *TalksService) GetTalkList(ctx context.Context, topicID int) ([]*Talk, *shared.Response, error) {
	u := fmt.Sprintf("topics/%d/talks", topicID)
	var result *struct {
		Talks []*Talk `json:"talks"`
	}
	resp, err := s.client.Get(ctx, u, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Talks, resp, nil
}

// GetMessagesInTalk fetches messages list in a talk.
//
// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/get-talk
func (s *TalksService) GetMessagesInTalk(ctx context.Context, topicID, talkID int, opt *GetMessagesOptions) (*MessagesInTalk, *shared.Response, error) {
	u, err := internal.AddQueries(fmt.Sprintf("topics/%d/talks/%d/posts", topicID, talkID), opt)
	if err != nil {
		return nil, nil, err
	}
	var result *MessagesInTalk
	resp, err := s.client.Get(ctx, u, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

type addMessageToTalkOptions struct {
	PostIds []int `json:"postIds[%d],omitempty"`
}

// AddMessagesToTalk adds messages to a talk.
//
// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/add-message-to-talk
func (s *TalksService) AddMessagesToTalk(ctx context.Context, topicID, talkID int, postIds ...int) (*MessagesInTalk, *shared.Response, error) {
	u := fmt.Sprintf("topics/%d/talks/%d/posts", topicID, talkID)
	var result *MessagesInTalk
	resp, err := s.client.Post(ctx, u, &addMessageToTalkOptions{postIds}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

type removeMessagesFromTalkOptions addMessageToTalkOptions

// RemoveMessagesFromTalk removes messages from a talk.
//
// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/remove-message-from-talk
func (s *TalksService) RemoveMessagesFromTalk(ctx context.Context, topicID, talkID int, postIds ...int) (*RemovedMessagesResult, *shared.Response, error) {
	u, err := internal.AddQueries(fmt.Sprintf("topics/%d/talks/%d/posts", topicID, talkID), &removeMessagesFromTalkOptions{postIds})
	if err != nil {
		return nil, nil, err
	}
	var result *RemovedMessagesResult
	resp, err := s.client.Delete(ctx, u, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}
