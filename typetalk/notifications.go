package typetalk

import (
	"context"
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

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/get-notifications
func (s *NotificationsService) GetNotificationList(ctx context.Context) (*NotificationList, *Response, error) {
	u := "notifications"
	var result *NotificationList
	if resp, err := s.client.get(ctx, u, &result); err != nil {
		return nil, resp, err
	} else {
		return result, resp, nil
	}
}

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/get-notification-status
func (s *NotificationsService) GetNotificationCount(ctx context.Context) (*NotificationCount, *Response, error) {
	u := "notifications/status"
	var result *NotificationCount
	if resp, err := s.client.get(ctx, u, &result); err != nil {
		return nil, resp, err
	} else {
		return result, resp, nil
	}
}

// Typetalk API docs: https://developer.nulab-inc.com/ja/docs/typetalk/api/1/open-notification
func (s *NotificationsService) ReadNotification(ctx context.Context) (*Access, *Response, error) {
	u := "notifications"
	var result *struct {
		Access *Access `json:"access"`
	}
	if resp, err := s.client.put(ctx, u, nil, &result); err != nil {
		return nil, resp, err
	} else {
		return result.Access, resp, nil
	}
}
