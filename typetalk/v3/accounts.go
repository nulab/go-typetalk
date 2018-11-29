package v3

import (
	"context"
	"time"

	"github.com/nulab/go-typetalk/typetalk/internal"
	"github.com/nulab/go-typetalk/typetalk/shared"
)

type AccountsService service

type Account struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	FullName   string     `json:"fullName"`
	Suggestion string     `json:"suggestion"`
	ImageURL   string     `json:"imageUrl"`
	IsBot      bool       `json:"isBot"`
	CreatedAt  *time.Time `json:"createdAt"`
	UpdatedAt  *time.Time `json:"updatedAt"`
}

type GetMyFriendsOptions struct {
	Offset int `json:"offset,omitempty"`
	Count  int `json:"count,omitempty"`
}

type getMyFriendsOptions struct {
	*GetMyFriendsOptions
	SpaceKey string `json:"spaceKey"`
	Q        string `json:"q"`
}

// Deprecated: Use GetMyFrieands in github.com/nulab/go-typetalk/typetalk/v4
func (s *AccountsService) GetMyFriends(ctx context.Context, spaceKey, q string, opt *GetMyFriendsOptions) ([]*Account, *shared.Response, error) {
	u, err := internal.AddQueries("search/friends", &getMyFriendsOptions{GetMyFriendsOptions: opt, SpaceKey: spaceKey, Q: q})
	if err != nil {
		return nil, nil, err
	}
	var result *struct {
		Accounts []*Account `json:"accounts"`
	}
	resp, err := s.client.Get(ctx, u, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Accounts, resp, nil
}
