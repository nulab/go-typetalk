package v2

import (
	"context"

	"time"

	. "github.com/nulab/go-typetalk/typetalk/shared"
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
	Unopened          int `json:"unopened"`
	UnopenedExcludeDM int `json:"unopenedExcludeDM"`
}

type NotificationCount struct {
	Statuses []*struct {
		MySpace *MySpace `json:"mySpace"`
		Space   *Space   `json:"space"`
		Access  *Access  `json:"access"`
		Like    *struct {
			Receive *struct {
				HasUnread  bool `json:"hasUnread"`
				ReadLikeID int  `json:"readLikeId"`
			} `json:"receive"`
		} `json:"like"`
		DirectMessage *struct {
			UnreadTopics int `json:"unreadTopics"`
		} `json:"directMessage"`
	} `json:"statuses"`
	DoNotDisturb *DoNotDisturb `json:"doNotDisturb"`
}

type MySpace struct {
	Space          *Space   `json:"space"`
	MyRole         string   `json:"myRole"`
	IsPaymentAdmin bool     `json:"isPaymentAdmin"`
	InvitableRoles []string `json:"invitableRoles"`
	MyPlan         MyPlan   `json:"myPlan"`
}

type MyPlan struct {
	Plan                *Plan       `json:"plan"`
	Enabled             bool        `json:"enabled"`
	Trial               interface{} `json:"trial"`
	NumberOfUsers       int         `json:"numberOfUsers"`
	TotalAttachmentSize int         `json:"totalAttachmentSize"`
	CreatedAt           *time.Time  `json:"createdAt"`
	UpdatedAt           *time.Time  `json:"updatedAt"`
}

type Plan struct {
	Key                      string `json:"key"`
	Name                     string `json:"name"`
	LimitNumberOfUsers       int    `json:"limitNumberOfUsers"`
	LimitTotalAttachmentSize int    `json:"limitTotalAttachmentSize"`
}

type DoNotDisturb struct {
	IsSuppressed bool       `json:"isSuppressed"`
	Manual       *Manual    `json:"manual"`
	Scheduled    *Scheduled `json:"scheduled"`
}

type Manual struct {
	RemainingTimeInMinutes interface{} `json:"remainingTimeInMinutes"`
}

type Scheduled struct {
	Enabled bool   `json:"enabled"`
	Start   string `json:"start"`
	End     string `json:"end"`
}

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/2/get-notification-status
func (s *NotificationsService) GetNotificationCount(ctx context.Context) (*NotificationCount, *Response, error) {
	u := "notifications/status"
	var result *NotificationCount
	if resp, err := s.client.Get(ctx, u, &result); err != nil {
		return nil, resp, err
	} else {
		return result, resp, nil
	}
}

// Deprecated: Use ReadNotification in github.com/nulab/go-typetalk/typetalk/v3
func (s *NotificationsService) ReadNotification(ctx context.Context) (*ReadNotificationResult, *Response, error) {
	u := "notifications"
	var result *ReadNotificationResult
	if resp, err := s.client.Put(ctx, u, nil, &result); err != nil {
		return nil, resp, err
	} else {
		return result, resp, nil
	}
}
