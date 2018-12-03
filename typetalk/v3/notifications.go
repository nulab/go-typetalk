package v3

import (
	"context"

	"github.com/nulab/go-typetalk/typetalk/shared"
)

type NotificationsService service

type ReadNotificationResult struct {
	Space  *Space  `json:"space"`
	Access *Access `json:"access"`
}

type Space struct {
	Key      string `json:"key"`
	Name     string `json:"name"`
	Enabled  bool   `json:"enabled"`
	ImageURL string `json:"imageUrl"`
}

type Access struct {
	Unopened int `json:"unopened"`
}

type readNotificationOptions struct {
	SpaceKey string `json:"spaceKey"`
}

// Typetalk API docs: https://developer.nulab-inc.com/ja/docs/typetalk/api/3/open-notification
func (s *NotificationsService) ReadNotification(ctx context.Context, spaceKey string) (*ReadNotificationResult, *shared.Response, error) {
	u := "notifications"
	var result *ReadNotificationResult
	resp, err := s.client.Put(ctx, u, &readNotificationOptions{spaceKey}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}
