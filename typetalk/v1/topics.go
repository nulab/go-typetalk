package v1

import (
	"context"
	"fmt"
	"time"

	"github.com/nulab/go-typetalk/typetalk/shared"
	"github.com/nulab/go-typetalk/v3/typetalk/internal"
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

// CreateTopic creates a topic.
//
// Typetalk API docs: https://developer.nulab.com/docs/typetalk/api/1/create-topic
func (s *TopicsService) CreateTopic(ctx context.Context, opt *CreateTopicOptions) (*TopicDetails, *shared.Response, error) {
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

// UpdateTopic updates a topic.
//
// Typetalk API docs: https://developer.nulab.com/docs/typetalk/api/1/update-topic
func (s *TopicsService) UpdateTopic(ctx context.Context, topicID int, opt *UpdateTopicOptions) (*TopicDetails, *shared.Response, error) {
	u := fmt.Sprintf("topics/%d", topicID)
	var result *TopicDetails
	resp, err := s.client.Put(ctx, u, opt, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// DeleteTopic deletes a topic.
//
// Typetalk API docs: https://developer.nulab.com/docs/typetalk/api/1/delete-topic
func (s *TopicsService) DeleteTopic(ctx context.Context, topicID int) (*Topic, *shared.Response, error) {
	u := fmt.Sprintf("topics/%d", topicID)
	var result *Topic
	resp, err := s.client.Delete(ctx, u, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// GetTopicDetails fetches a topic's detailed information.
//
// Typetalk API docs: https://developer.nulab.com/docs/typetalk/api/1/get-topic-details
func (s *TopicsService) GetTopicDetails(ctx context.Context, topicID int) (*TopicDetails, *shared.Response, error) {
	u := fmt.Sprintf("topics/%d", topicID)
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

// GetTopicMessages fetches messages list in a topic.
//
// Typetalk API docs: https://developer.nulab.com/docs/typetalk/api/1/get-messages
func (s *TopicsService) GetTopicMessages(ctx context.Context, topicID int, opt *GetTopicMessagesOptions) (*TopicMessages, *shared.Response, error) {
	u, err := internal.AddQueries(fmt.Sprintf("topics/%d", topicID), opt)
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
	RemoveAccountsID                    []int    `json:"removeAccounts[%d].id,omitempty"`
	RemoveAccountsCancelSpaceInvitation []bool   `json:"removeAccounts[%d].cancelSpaceInvitation,omitempty"`
	RemoveGroupIds                      []bool   `json:"removeGroupIds[%d],omitempty"`
}

// UpdateTopicMembers updates members in a topic.
//
// Typetalk API docs: https://developer.nulab.com/docs/typetalk/api/1/update-topic-members
func (s *TopicsService) UpdateTopicMembers(ctx context.Context, topicID int, opt *UpdateTopicMembersOptions) (*TopicDetails, *shared.Response, error) {
	u := fmt.Sprintf("topics/%d/members/update", topicID)
	var result *TopicDetails
	resp, err := s.client.Post(ctx, u, opt, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// FavoriteTopic marks a topic as favorite.
//
// Typetalk API docs: https://developer.nulab.com/docs/typetalk/api/1/favorite-topic
func (s *TopicsService) FavoriteTopic(ctx context.Context, topicID int) (*FavoriteTopic, *shared.Response, error) {
	u := fmt.Sprintf("topics/%d/favorite", topicID)
	var result *FavoriteTopic
	resp, err := s.client.Post(ctx, u, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// UnfavoriteTopic marks a topic as unfavorite.
//
// Typetalk API docs: https://developer.nulab.com/docs/typetalk/api/1/unfavorite-topic
func (s *TopicsService) UnfavoriteTopic(ctx context.Context, topicID int) (*FavoriteTopic, *shared.Response, error) {
	u := fmt.Sprintf("topics/%d/favorite", topicID)
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

// ReadMessagesInTopic mark a message as read.
//
// Typetalk API docs: https://developer.nulab.com/docs/typetalk/api/1/save-read-topic
func (s *TopicsService) ReadMessagesInTopic(ctx context.Context, topicID, postID int) (*Unread, *shared.Response, error) {
	u, err := internal.AddQueries("bookmarks", &readMessagesInTopicOptions{topicID, postID})
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

// GetMyTopics fetches topics list.
//
// Deprecated: Use GetMyTopics v2
func (s *TopicsService) GetMyTopics(ctx context.Context) ([]*FavoriteTopicWithUnread, *shared.Response, error) {
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
