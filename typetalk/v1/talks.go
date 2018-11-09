package v1

import (
	"context"
	"fmt"
	"time"

	. "github.com/nulab/go-typetalk/typetalk/internal"
	. "github.com/nulab/go-typetalk/typetalk/shared"
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

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/create-talk
func (s *TalksService) CreateTalk(ctx context.Context, topicId int, talkName string, postIds ...int) (*CreatedTalkResult, *Response, error) {
	u := fmt.Sprintf("topics/%d/talks", topicId)
	var result *CreatedTalkResult
	if resp, err := s.client.Post(ctx, u, &CreateTalkOptions{talkName, postIds}, &result); err != nil {
		return nil, resp, err
	} else {
		return result, resp, nil
	}
}

type updateTalkOptions struct {
	TalkName string `json:"talkName,omitempty"`
}

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/update-talk
func (s *TalksService) UpdateTalk(ctx context.Context, topicId, talkId int, talkName string) (*UpdatedTalkResult, *Response, error) {
	u, err := AddQueries(fmt.Sprintf("topics/%d/talks/%d", topicId, talkId), &updateTalkOptions{talkName})
	if err != nil {
		return nil, nil, err
	}

	var result *UpdatedTalkResult
	if resp, err := s.client.Put(ctx, u, nil, &result); err != nil {
		return nil, resp, err
	} else {
		return result, resp, nil
	}
}

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/delete-talk
func (s *TalksService) DeleteTalk(ctx context.Context, topicId, talkId int) (*DeletedTalkResult, *Response, error) {
	u := fmt.Sprintf("topics/%d/talks/%d", topicId, talkId)
	var result *DeletedTalkResult
	if resp, err := s.client.Delete(ctx, u, &result); err != nil {
		return nil, resp, err
	} else {
		return result, resp, nil
	}
}

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/get-talks
func (s *TalksService) GetTalkList(ctx context.Context, topicId int) ([]*Talk, *Response, error) {
	u := fmt.Sprintf("topics/%d/talks", topicId)
	var result *struct {
		Talks []*Talk `json:"talks"`
	}
	if resp, err := s.client.Get(ctx, u, &result); err != nil {
		return nil, resp, err
	} else {
		return result.Talks, resp, nil
	}
}

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/get-talk
func (s *TalksService) GetMessagesInTalk(ctx context.Context, topicId, talkId int, opt *GetMessagesOptions) (*MessagesInTalk, *Response, error) {
	u, err := AddQueries(fmt.Sprintf("topics/%d/talks/%d/posts", topicId, talkId), opt)
	if err != nil {
		return nil, nil, err
	}
	var result *MessagesInTalk
	if resp, err := s.client.Get(ctx, u, &result); err != nil {
		return nil, resp, err
	} else {
		return result, resp, nil
	}
}

type addMessageToTalkOptions struct {
	PostIds []int `json:"postIds[%d],omitempty"`
}

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/add-message-to-talk
func (s *TalksService) AddMessagesToTalk(ctx context.Context, topicId, talkId int, postIds ...int) (*MessagesInTalk, *Response, error) {
	u := fmt.Sprintf("topics/%d/talks/%d/posts", topicId, talkId)
	var result *MessagesInTalk
	if resp, err := s.client.Post(ctx, u, &addMessageToTalkOptions{postIds}, &result); err != nil {
		return nil, resp, err
	} else {
		return result, resp, nil
	}
}

type removeMessagesFromTalkOptions addMessageToTalkOptions

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/remove-message-from-talk
func (s *TalksService) RemoveMessagesFromTalk(ctx context.Context, topicId, talkId int, postIds ...int) (*RemovedMessagesResult, *Response, error) {
	u, err := AddQueries(fmt.Sprintf("topics/%d/talks/%d/posts", topicId, talkId), &removeMessagesFromTalkOptions{postIds})
	if err != nil {
		return nil, nil, err
	}
	var result *RemovedMessagesResult
	if resp, err := s.client.Delete(ctx, u, &result); err != nil {
		return nil, resp, err
	} else {
		return result, resp, nil
	}
}
