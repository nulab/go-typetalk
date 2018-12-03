package v1

import (
	"context"
	"fmt"
	"time"

	"github.com/nulab/go-typetalk/typetalk/internal"
	"github.com/nulab/go-typetalk/typetalk/shared"
)

type OrganizationsService service

type Organization struct {
	Space          *Space  `json:"space"`
	MyRole         string  `json:"myRole"`
	IsPaymentAdmin bool    `json:"isPaymentAdmin"`
	MyPlan         *MyPlan `json:"myPlan"`
}

type Space struct {
	Key      string `json:"key"`
	Name     string `json:"name"`
	Enabled  bool   `json:"enabled"`
	ImageURL string `json:"imageUrl"`
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

type Group struct {
	ID         int        `json:"id"`
	Key        string     `json:"key"`
	Name       string     `json:"name"`
	Suggestion string     `json:"suggestion"`
	ImageURL   string     `json:"imageUrl"`
	CreatedAt  *time.Time `json:"createdAt"`
	UpdatedAt  *time.Time `json:"updatedAt"`
}

type OrganizationMembers struct {
	Accounts []*Account `json:"accounts"`
	Groups   []*struct {
		Group       *Group `json:"group"`
		MemberCount int    `json:"memberCount"`
	} `json:"groups"`
}

type organizationsGetOptions struct {
	ExcludesGuest bool `json:"excludesGuest,omitempty"`
}

// GetMyOrganizations fetches organizations list.
//
// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/get-spaces
func (s *OrganizationsService) GetMyOrganizations(ctx context.Context, excludesGuest bool) ([]*Organization, *shared.Response, error) {
	u, err := internal.AddQueries("spaces", &organizationsGetOptions{excludesGuest})
	if err != nil {
		return nil, nil, err
	}
	var result *struct {
		MySpaces []*Organization `json:"mySpaces"`
	}
	resp, err := s.client.Get(ctx, u, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.MySpaces, resp, nil
}

// GetOrganizationMembers fetches an organization's members list.
//
// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/get-space-members
func (s *OrganizationsService) GetOrganizationMembers(ctx context.Context, spaceKey string) (*OrganizationMembers, *shared.Response, error) {
	u := fmt.Sprintf("spaces/%s/members", spaceKey)
	var result *OrganizationMembers
	resp, err := s.client.Get(ctx, u, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}
