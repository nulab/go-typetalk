package v5

import (
	"context"

	"time"

	"github.com/nulab/go-typetalk/typetalk/shared"
)

type NotificationsService service

type NotificationCount struct {
	Statuses []*struct {
		MySpace *MySpace `json:"mySpace"`
		Access  *Access  `json:"access"`
		Like    *struct {
			Receive *struct {
				HasUnread  bool `json:"hasUnread"`
				ReadLikeID int  `json:"readLikeId"`
			} `json:"receive"`
		} `json:"like"`
		Unreads *struct {
			TopicIds   []int `json:"topicIds"`
			DMTopicIds []int `json:"dmTopicIds"`
		} `json:"unreads"`
	} `json:"statuses"`
	NotificationSettings *struct {
		FavoriteTopicMobile bool          `json:"favoriteTopicMobile"`
		DoNotDisturb        *DoNotDisturb `json:"doNotDisturb"`
	} `json:"notificationSettings"`
}

type MySpace struct {
	Space          *Space   `json:"space"`
	MyRole         string   `json:"myRole"`
	IsPaymentAdmin bool     `json:"isPaymentAdmin"`
	InvitableRoles []string `json:"invitableRoles"`
	MyPlan         MyPlan   `json:"myPlan"`
}

type Space struct {
	Key      string `json:"key"`
	Name     string `json:"name"`
	Enabled  bool   `json:"enabled"`
	ImageURL string `json:"imageUrl"`
}

type MyPlan struct {
	Plan                     *Plan       `json:"plan"`
	Enabled                  bool        `json:"enabled"`
	Trial                    interface{} `json:"trial"`
	NumberOfUsers            int         `json:"numberOfUsers"`
	NumberOfAllowedAddresses int         `json:"numberOfAllowedAddresses"`
	TotalAttachmentSize      int         `json:"totalAttachmentSize"`
	CreatedAt                *time.Time  `json:"createdAt"`
	UpdatedAt                *time.Time  `json:"updatedAt"`
}

type Access struct {
	Unopened          int `json:"unopened"`
	UnopenedExcludeDM int `json:"unopenedExcludeDM"`
}

type Plan struct {
	Key                           string `json:"key"`
	Name                          string `json:"name"`
	LimitNumberOfUsers            int    `json:"limitNumberOfUsers"`
	LimitNumberOfAllowedAddresses int    `json:"limitNumberOfAllowedAddresses"`
	LimitTotalAttachmentSize      int    `json:"limitTotalAttachmentSize"`
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

// GetNotificationCount fetches notification counts.
//
// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/5/get-notification-status/
func (s *NotificationsService) GetNotificationCount(ctx context.Context) (*NotificationCount, *shared.Response, error) {
	u := "notifications/status"
	var result *NotificationCount
	resp, err := s.client.Get(ctx, u, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}
