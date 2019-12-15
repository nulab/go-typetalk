package v2

import (
	"context"
	"time"

	"github.com/nulab/go-typetalk/v3/typetalk/internal"
	"github.com/nulab/go-typetalk/v3/typetalk/shared"
)

type MentionsService service

type Mention struct {
	ID     int        `json:"id"`
	ReadAt *time.Time `json:"readAt"`
	Post   *Post      `json:"post"`
}

type GetMentionListOptions struct {
	From   int  `json:"from,omitempty"`
	Unread bool `json:"unread,omitempty"`
}

type getMentionListOptions struct {
	*GetMentionListOptions
	SpaceKey string `json:"spaceKey"`
}

// GetMentionList fetches mentions list.
//
// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/2/get-mentions
func (s *MentionsService) GetMentionList(ctx context.Context, spaceKey string, opt *GetMentionListOptions) ([]*Mention, *shared.Response, error) {
	u, err := internal.AddQueries("mentions", &getMentionListOptions{opt, spaceKey})
	if err != nil {
		return nil, nil, err
	}
	var result *struct {
		Mentions []*Mention `json:"mentions"`
	}
	resp, err := s.client.Get(ctx, u, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Mentions, resp, nil
}
