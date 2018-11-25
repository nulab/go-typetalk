package v1

import (
	"context"
	"fmt"
	"time"

	. "github.com/nulab/go-typetalk/typetalk/internal"
	. "github.com/nulab/go-typetalk/typetalk/shared"
)

type TopicsService service

type Topic struct {
	ID              int        `json:"id"`
	Name            string     `json:"name"`
	Description     string     `json:"description"`
	Suggestion      string     `json:"suggestion"`
	IsDirectMessage bool       `json:"isDirectMessage"`
	LastPostedAt    *time.Time `json:"lastPostedAt"`
	CreatedAt       *time.Time `json:"createdAt"`
	UpdatedAt       *time.Time `json:"updatedAt"`
}

type TopicDetails struct {
	Topic   *Topic        `json:"topic"`
	MySpace *Organization `json:"mySpace"`
	Teams   []interface{} `json:"teams"`
	Groups  []*struct {
		Group       *Group `json:"group"`
		MemberCount int    `json:"memberCount"`
	} `json:"groups"`
	Accounts             []*Account    `json:"accounts"`
	InvitingAccounts     []interface{} `json:"invitingAccounts"`
	Invites              []interface{} `json:"invites"`
	AccountsForAPI       []interface{} `json:"accountsForApi"`
	Integrations         []interface{} `json:"integrations"`
	RemainingInvitations interface{}   `json:"remainingInvitations"`
}

type Bookmark struct {
	PostID    int       `json:"postId"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type TopicMessages struct {
	MySpace                *Organization `json:"mySpace"`
	Team                   interface{}   `json:"team"`
	Topic                  *Topic        `json:"topic"`
	Bookmark               *Bookmark     `json:"bookmark"`
	Posts                  []*Post       `json:"posts"`
	HasNext                bool          `json:"hasNext"`
	ExceedsAttachmentLimit bool          `json:"exceedsAttachmentLimit"`
}

type FavoriteTopic struct {
	Topic    *Topic `json:"topic"`
	Favorite bool   `json:"favorite"`
}

type FavoriteTopicWithUnread struct {
	FavoriteTopic
	Unread Unread `json:"unread"`
}

type CreateTopicOptions struct {
	Name          string `json:"name,omitempty"`
	SpaceKey      string `json:"spaceKey,omitempty"`
	AddAccountIds []int  `json:"addAccountIds[%d],omitempty"`
	AddGroupIds   []int  `json:"addGroupIds[%d],omitempty"`
}

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/create-topic
func (s *TopicsService) CreateTopic(ctx context.Context, opt *CreateTopicOptions) (*TopicDetails, *Response, error) {
	u := "topics"
	var result *TopicDetails
	resp, err := s.client.Post(ctx, u, opt, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

type UpdateTopicOptions struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/update-topic
func (s *TopicsService) UpdateTopic(ctx context.Context, topicId int, opt *UpdateTopicOptions) (*TopicDetails, *Response, error) {
	u := fmt.Sprintf("topics/%d", topicId)
	var result *TopicDetails
	resp, err := s.client.Put(ctx, u, opt, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/delete-topic
func (s *TopicsService) DeleteTopic(ctx context.Context, topicId int) (*Topic, *Response, error) {
	u := fmt.Sprintf("topics/%d", topicId)
	var result *Topic
	resp, err := s.client.Delete(ctx, u, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/get-topic-details
func (s *TopicsService) GetTopicDetails(ctx context.Context, topicId int) (*TopicDetails, *Response, error) {
	u := fmt.Sprintf("topics/%d", topicId)
	var result *TopicDetails
	resp, err := s.client.Get(ctx, u, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

type GetTopicMessagesOptions struct {
	Count     int    `json:"count,omitempty"`
	From      int    `json:"from,omitempty"`
	Direction string `json:"direction,omitempty"`
}

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/get-messages
func (s *TopicsService) GetTopicMessages(ctx context.Context, topicId int, opt *GetTopicMessagesOptions) (*TopicMessages, *Response, error) {
	u, err := AddQueries(fmt.Sprintf("topics/%d", topicId), opt)
	if err != nil {
		return nil, nil, err
	}
	var result *TopicMessages
	resp, err := s.client.Get(ctx, u, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

type UpdateTopicMembersOptions struct {
	AddAccountIds                       []int    `json:"addAccountIds[%d],omitempty"`
	AddGroupIds                         []int    `json:"addGroupIds[%d],omitempty"`
	InvitationsEmail                    []string `json:"invitations[%d].email,omitempty"`
	InvitationsRole                     []string `json:"invitations[%d].role,omitempty"`
	RemoveAccountsId                    []int    `json:"removeAccounts[%d].id,omitempty"`
	RemoveAccountsCancelSpaceInvitation []bool   `json:"removeAccounts[%d].cancelSpaceInvitation,omitempty"`
	RemoveGroupIds                      []bool   `json:"removeGroupIds[%d],omitempty"`
}

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/update-topic-members
func (s *TopicsService) UpdateTopicMembers(ctx context.Context, topicId int, opt *UpdateTopicMembersOptions) (*TopicDetails, *Response, error) {
	u := fmt.Sprintf("topics/%d/members/update", topicId)
	var result *TopicDetails
	resp, err := s.client.Post(ctx, u, opt, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/favorite-topic
func (s *TopicsService) FavoriteTopic(ctx context.Context, topicId int) (*FavoriteTopic, *Response, error) {
	u := fmt.Sprintf("topics/%d/favorite", topicId)
	var result *FavoriteTopic
	resp, err := s.client.Post(ctx, u, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/unfavorite-topic
func (s *TopicsService) UnfavoriteTopic(ctx context.Context, topicId int) (*FavoriteTopic, *Response, error) {
	u := fmt.Sprintf("topics/%d/favorite", topicId)
	var result *FavoriteTopic
	resp, err := s.client.Delete(ctx, u, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

type readMessagesInTopicOptions struct {
	TopicID int `json:"topicId,omitempty"`
	PostID  int `json:"postId,omitempty"`
}

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/save-read-topic
func (s *TopicsService) ReadMessagesInTopic(ctx context.Context, topicId, postId int) (*Unread, *Response, error) {
	u, err := AddQueries("bookmarks", &readMessagesInTopicOptions{topicId, postId})
	if err != nil {
		return nil, nil, err
	}
	var result *struct {
		Unread *Unread `json:"unread"`
	}
	resp, err := s.client.Put(ctx, u, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Unread, resp, nil
}

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/get-topics
func (s *TopicsService) GetMyTopics(ctx context.Context) ([]*FavoriteTopicWithUnread, *Response, error) {
	u := "topics"
	var result *struct {
		Topics []*FavoriteTopicWithUnread `json:"topics"`
	}
	resp, err := s.client.Get(ctx, u, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Topics, resp, nil
}
