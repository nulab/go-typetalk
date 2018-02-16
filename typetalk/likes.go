package typetalk

import (
	"context"
)

type LikesService service

type LikedPost struct {
	Post          *Post          `json:"post"`
	Likes         []*Like        `json:"likes"`
	DirectMessage *DirectMessage `json:"directMessage"`
}

type GetLikesOptions struct {
	From int `json:"from,omitempty"`
}

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/get-likes-receive/
func (s *MessagesService) GetLikes(ctx context.Context, opt *GetLikesOptions) ([]*LikedPost, *Response, error) {
	u, err := addQueries("likes/receive", opt)
	if err != nil {
		return nil, nil, err
	}
	var result *struct {
		LikedPosts []*LikedPost `json:"likedPosts"`
	}
	if resp, err := s.client.get(ctx, u, &result); err != nil {
		return nil, resp, err
	} else {
		return result.LikedPosts, resp, nil
	}
}
