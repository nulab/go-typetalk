package v1

import (
	"context"

	"github.com/nulab/go-typetalk/v3/typetalk/shared"
)

type NotificationsService service

type Invites struct {
	Teams  []interface{} `json:"teams"`
	Topics []*Topic      `json:"topics"`
}

type NotificationList struct {
	Mentions []*Mention `json:"mentions"`
	Invites  *Invites   `json:"invites"`
}

type Access struct {
	Unopened int `json:"unopened"`
}

type NotificationCount struct {
	Mention *struct {
		Unread int `json:"unread"`
	} `json:"mention"`
	Access *Access `json:"access"`
	Invite *struct {
		Team *struct {
			Pending int `json:"pending"`
		} `json:"team"`
		Topic *struct {
			Pending int `json:"pending"`
		} `json:"topic"`
	} `json:"invite"`
	Like *struct {
		Receive interface{} `json:"receive"`
	} `json:"like"`
	DirectMessage *struct {
		UnreadTopics int `json:"unreadTopics"`
	} `json:"directMessage"`
}

// GetNotificationList fetches notifications list.
//
// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/get-notifications
func (s *NotificationsService) GetNotificationList(ctx context.Context) (*NotificationList, *shared.Response, error) {
	u := "notifications"
	var result *NotificationList
	resp, err := s.client.Get(ctx, u, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// GetNotificationCount fetches notification counts.
//
// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/get-notification-status
func (s *NotificationsService) GetNotificationCount(ctx context.Context) (*NotificationCount, *shared.Response, error) {
	u := "notifications/status"
	var result *NotificationCount
	resp, err := s.client.Get(ctx, u, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// ReadNotification marks notifications as read.
//
// Deprecated: Use ReadNotification in github.com/nulab/go-typetalk/v3/typetalk/v3
func (s *NotificationsService) ReadNotification(ctx context.Context) (*Access, *shared.Response, error) {
	u := "notifications"
	var result *struct {
		Access *Access `json:"access"`
	}
	resp, err := s.client.Put(ctx, u, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Access, resp, nil
}
