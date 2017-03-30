package typetalk

import (
	"context"
	"fmt"
	"time"
)

type MentionsService service

type Mention struct {
	ID     int        `json:"id"`
	ReadAt *time.Time `json:"readAt"`
	Post   *Post      `json:"post"`
}

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/save-read-mention
func (s *MentionsService) ReadMention(ctx context.Context, mentionId int) (*Mention, *Response, error) {
	u := fmt.Sprintf("mentions/%d", mentionId)
	var result *struct {
		Mention Mention `json:"mention"`
	}
	if resp, err := s.client.put(ctx, u, nil, &result); err != nil {
		return nil, resp, err
	} else {
		return &result.Mention, resp, nil
	}
}

type GetMentionListOptions struct {
	From   int  `json:"from,omitempty"`
	Unread bool `json:"unread,omitempty"`
}

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/get-mentions
func (s *MentionsService) GetMentionList(ctx context.Context, opt *GetMentionListOptions) ([]*Mention, *Response, error) {
	u, err := addQueries("mentions", opt)
	if err != nil {
		return nil, nil, err
	}
	var result *struct {
		Mentions []*Mention `json:"mentions"`
	}
	if resp, err := s.client.get(ctx, u, &result); err != nil {
		return nil, resp, err
	} else {
		return result.Mentions, resp, nil
	}
}
