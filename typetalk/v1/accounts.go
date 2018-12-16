package v1

import (
	"context"
	"fmt"
	"time"

	"github.com/nulab/go-typetalk/typetalk/internal"
	"github.com/nulab/go-typetalk/typetalk/shared"
)

// AccountsService handles communication with the account related API.
type AccountsService service

// Account represents Typetalk account information.
type Account struct {
	ID             int        `json:"id"`
	Name           string     `json:"name"`
	FullName       string     `json:"fullName"`
	Suggestion     string     `json:"suggestion"`
	ImageURL       string     `json:"imageUrl"`
	IsBot          bool       `json:"isBot"`
	Lang           string     `json:"lang"`
	TimezoneID     string     `json:"timezoneId"`
	CreatedAt      *time.Time `json:"createdAt"`
	UpdatedAt      *time.Time `json:"updatedAt"`
	ImageUpdatedAt *time.Time `json:"imageUpdatedAt"`
}

// Status represents online status of the user.
type Status struct {
	Presence *string     `json:"presence"`
	Web      interface{} `json:"web"`
	Mobile   interface{} `json:"mobile"`
}

// AccountStatus contains account and status information.
type AccountStatus struct {
	Account *Account `json:"account"`
	Status  *Status  `json:"status"`
}

// Profile is alias for AccountStatus.
type Profile AccountStatus

// MyProfile represents the user's information.
type MyProfile struct {
	Account *Account `json:"account"`
	Lang    string   `json:"lang"`
	Theme   string   `json:"theme"`
	Post    *struct {
		UseCtrl       bool   `json:"useCtrl"`
		EmojiSkinTone string `json:"emojiSkinTone"`
	} `json:"post"`
}

// Friends represents accounts search result.
type Friends struct {
	Count    int              `json:"count"`
	Accounts []*AccountStatus `json:"accounts"`
}

// OnlineStatus contains accounts information.
type OnlineStatus struct {
	Accounts []*AccountStatus `json:"accounts"`
}

// GetMyProfile fetches the user's account information.
//
// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/get-profile
func (s *AccountsService) GetMyProfile(ctx context.Context) (*MyProfile, *shared.Response, error) {
	u := "profile"
	var result *MyProfile
	resp, err := s.client.Get(ctx, u, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// GetFriendProfile fetches other user's account information.
//
// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/get-friend-profile
func (s *AccountsService) GetFriendProfile(ctx context.Context, accountName string) (*Profile, *shared.Response, error) {
	u := fmt.Sprintf("profile/%s", accountName)
	var result *Profile
	resp, err := s.client.Get(ctx, u, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// GetMyFriendsOptions represents request parameters for "search/friends" API.
type GetMyFriendsOptions struct {
	Q      string `json:"q,omitempty"`
	Offset int    `json:"offset,omitempty"`
	Count  int    `json:"count,omitempty"`
}

// GetMyFriends searches other user who belong to a topic in common.
//
// Deprecated: Use GetMyFrieands in github.com/nulab/go-typetalk/typetalk/v4
func (s *AccountsService) GetMyFriends(ctx context.Context, opt *GetMyFriendsOptions) (*Friends, *shared.Response, error) {
	u, err := internal.AddQueries("search/friends", opt)
	if err != nil {
		return nil, nil, err
	}
	var result *Friends
	resp, err := s.client.Get(ctx, u, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

type searchAccountsOptions struct {
	NameOrEmailAddress string `json:"nameOrEmailAddress,omitempty"`
}

// SearchAccounts searches acocunts by name or mail address.
//
// Deprecated: Use GetMyFrieands in github.com/nulab/go-typetalk/typetalk/v4
func (s *AccountsService) SearchAccounts(ctx context.Context, nameOrEmailAddress string) (*Account, *shared.Response, error) {
	u, err := internal.AddQueries("search/accounts", &searchAccountsOptions{nameOrEmailAddress})
	if err != nil {
		return nil, nil, err
	}
	var result *Account
	resp, err := s.client.Get(ctx, u, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

type getOnlineStatusOptions struct {
	AccountIds []int `json:"accountIds[%d],omitempty"`
}

// GetOnlineStatus fetches an user's online status.
//
// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/get-online-status
func (s *AccountsService) GetOnlineStatus(ctx context.Context, accountIds ...int) (*OnlineStatus, *shared.Response, error) {
	u, err := internal.AddQueries("accounts/status", &getOnlineStatusOptions{accountIds})
	if err != nil {
		return nil, nil, err
	}
	var result *OnlineStatus
	resp, err := s.client.Get(ctx, u, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}
