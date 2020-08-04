package v1

import (
	"context"
	"fmt"
	"time"

	"github.com/nulab/go-typetalk/typetalk/shared"
)

type StatusesService service

type UserStatus struct {
	ID                     int       `json:"id"`
	AccountID              int       `json:"accountId"`
	SpaceID                int       `json:"spaceId"`
	Emoji                  string    `json:"emoji"`
	Message                string    `json:"message"`
	ClearAt                time.Time `json:"clearAt"`
	IsNotificationDisabled bool      `json:"isNotificationDisabled"`
	CreatedAt              time.Time `json:"createdAt"`
	UpdatedAt              time.Time `json:"updatedAt"`
}

type SaveUserStatusOptions struct {
	Message                string
	ClearAt                string
	IsNotificationDisabled bool
}

type saveUserStatusOptions struct {
	Emoji                  string `json:"emoji"`
	Message                string `json:"message,omitempty"`
	ClearAt                string `json:"clearAt,omitempty"`
	IsNotificationDisabled bool   `json:"isNotificationDisabled,omitempty"`
}

type SaveUserStatusResult struct {
	UserStatus *UserStatus `json:"userStatus"`
}

// SaveUserStatus save a user status.
//
// Typetalk API docs: https://developer.nulab.com/docs/typetalk/api/1/save-user-status/#save-user-status
func (s *StatusesService) SaveUserStatus(ctx context.Context, spaceKey, emoji string, opt *SaveUserStatusOptions) (*SaveUserStatusResult, *shared.Response, error) {
	u := fmt.Sprintf("spaces/%s/userStatuses", spaceKey)
	var (
		result *SaveUserStatusResult
		params saveUserStatusOptions
	)
	params.Emoji = emoji
	if opt != nil {
		params.Message = opt.Message
		params.ClearAt = opt.ClearAt
		params.IsNotificationDisabled = opt.IsNotificationDisabled
	}
	resp, err := s.client.Post(ctx, u, &params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}
