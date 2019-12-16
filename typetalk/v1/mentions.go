package v1

import (
	"context"
	"fmt"
	"time"

	"github.com/nulab/go-typetalk/typetalk/shared"
	"github.com/nulab/go-typetalk/v3/typetalk/internal"
)

type MentionsService service

type Mention struct {
	ID     int        `json:"id"`
	ReadAt *time.Time `json:"readAt"`
	Post   *Post      `json:"post"`
}

// ReadMention marks a mention as read.
//
// Typetalk API docs: https://developer.nulab.com/docs/typetalk/api/1/save-read-mention
func (s *MentionsService) ReadMention(ctx context.Context, mentionID int) (*Mention, *shared.Response, error) {
	u := fmt.Sprintf("mentions/%d", mentionID)
	var result *struct {
		Mention Mention `json:"mention"`
	}
	resp, err := s.client.Put(ctx, u, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result.Mention, resp, nil
}

type GetMentionListOptions struct {
	From   int  `json:"from,omitempty"`
	Unread bool `json:"unread,omitempty"`
}

// GetMentionList fetches mentions list.
//
// Deprecated: Use GetMentionList v2
func (s *MentionsService) GetMentionList(ctx context.Context, opt *GetMentionListOptions) ([]*Mention, *shared.Response, error) {
	u, err := internal.AddQueries("mentions", opt)
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
