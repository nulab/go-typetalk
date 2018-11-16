package v1

import (
	"context"
	"fmt"
	"time"

	. "github.com/nulab/go-typetalk/typetalk/internal"
	. "github.com/nulab/go-typetalk/typetalk/shared"
)

type AccountsService service

type Account struct {
	ID             int        `json:"id"`
	Name           string     `json:"name"`
	FullName       string     `json:"fullName"`
	Suggestion     string     `json:"suggestion"`
	ImageURL       string     `json:"imageUrl"`
	IsBot          bool       `json:"isBot"`
	Lang           string     `json:"lang"`
	TimezoneId     string     `json:"timezoneId"`
	CreatedAt      *time.Time `json:"createdAt"`
	UpdatedAt      *time.Time `json:"updatedAt"`
	ImageUpdatedAt *time.Time `json:"imageUpdatedAt"`
}

type Status struct {
	Presence *string     `json:"presence"`
	Web      interface{} `json:"web"`
	Mobile   interface{} `json:"mobile"`
}

type AccountStatus struct {
	Account *Account `json:"account"`
	Status  *Status  `json:"status"`
}

type Profile AccountStatus

type MyProfile struct {
	Account *Account `json:"account"`
	Lang    string   `json:"lang"`
	Theme   string   `json:"theme"`
	Post    *struct {
		UseCtrl       bool   `json:"useCtrl"`
		EmojiSkinTone string `json:"emojiSkinTone"`
	} `json:"post"`
}

type Friends struct {
	Count    int              `json:"count"`
	Accounts []*AccountStatus `json:"accounts"`
}

type OnlineStatus struct {
	Accounts []*AccountStatus `json:"accounts"`
}

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/get-profile
func (s *AccountsService) GetMyProfile(ctx context.Context) (*MyProfile, *Response, error) {
	u := "profile"
	var result *MyProfile
	if resp, err := s.client.Get(ctx, u, &result); err != nil {
		return nil, resp, err
	} else {
		return result, resp, nil
	}
}

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/get-friend-profile
func (s *AccountsService) GetFriendProfile(ctx context.Context, accountName string) (*Profile, *Response, error) {
	u := fmt.Sprintf("profile/%s", accountName)
	var result *Profile
	if resp, err := s.client.Get(ctx, u, &result); err != nil {
		return nil, resp, err
	} else {
		return result, resp, nil
	}
}

type GetMyFriendsOptions struct {
	Q      string `json:"q,omitempty"`
	Offset int    `json:"offset,omitempty"`
	Count  int    `json:"count,omitempty"`
}

// Deprecated: Use GetMyFrieands in github.com/nulab/go-typetalk/typetalk/v4
func (s *AccountsService) GetMyFriends(ctx context.Context, opt *GetMyFriendsOptions) (*Friends, *Response, error) {
	u, err := AddQueries("search/friends", opt)
	if err != nil {
		return nil, nil, err
	}
	var result *Friends
	if resp, err := s.client.Get(ctx, u, &result); err != nil {
		return nil, resp, err
	} else {
		return result, resp, nil
	}
}

type searchAccountsOptions struct {
	NameOrEmailAddress string `json:"nameOrEmailAddress,omitempty"`
}

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/search-accounts
func (s *AccountsService) SearchAccounts(ctx context.Context, nameOrEmailAddress string) (*Account, *Response, error) {
	u, err := AddQueries("search/accounts", &searchAccountsOptions{nameOrEmailAddress})
	if err != nil {
		return nil, nil, err
	}
	var result *Account
	if resp, err := s.client.Get(ctx, u, &result); err != nil {
		return nil, resp, err
	} else {
		return result, resp, nil
	}
}

type getOnlineStatusOptions struct {
	AccountIds []int `json:"accountIds[%d],omitempty"`
}

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/get-online-status
func (s *AccountsService) GetOnlineStatus(ctx context.Context, accountIds ...int) (*OnlineStatus, *Response, error) {
	u, err := AddQueries("accounts/status", &getOnlineStatusOptions{accountIds})
	if err != nil {
		return nil, nil, err
	}
	var result *OnlineStatus
	if resp, err := s.client.Get(ctx, u, &result); err != nil {
		return nil, resp, err
	} else {
		return result, resp, nil
	}
}
