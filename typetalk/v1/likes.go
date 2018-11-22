package v1

import (
	"context"
	. "github.com/nulab/go-typetalk/typetalk/internal"
	. "github.com/nulab/go-typetalk/typetalk/shared"
	"time"
)

type LikesService service

type LikedPost struct {
	Post          *Post          `json:"post"`
	Likes         []*Like        `json:"likes"`
	DirectMessage *DirectMessage `json:"directMessage"`
}

type DiscoverLikedPost struct {
	*LikedPost
}

type ReceiveLikedPost struct {
	*LikedPost
}

type GiveLikedPost struct {
	*ReceiveLikedPost
	MyLike *MyLike `json:"myLike"`
}

type MyLike struct {
	ID        int       `json:"id"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"createdAt"`
}

type GetLikesOptions struct {
	From int `json:"from,omitempty"`
}

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/get-likes-receive/
func (s *LikesService) GetLikesReceive(ctx context.Context, opt *GetLikesOptions) ([]*ReceiveLikedPost, *Response, error) {
	u, err := AddQueries("likes/receive", opt)
	if err != nil {
		return nil, nil, err
	}
	var result *struct {
		LikedPosts []*ReceiveLikedPost `json:"likedPosts"`
	}
	if resp, err := s.client.Get(ctx, u, &result); err != nil {
		return nil, resp, err
	} else {
		return result.LikedPosts, resp, nil
	}
}

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/get-likes-give/
func (s *LikesService) GetLikesGive(ctx context.Context, opt *GetLikesOptions) ([]*GiveLikedPost, *Response, error) {
	u, err := AddQueries("likes/give", opt)
	if err != nil {
		return nil, nil, err
	}
	var result *struct {
		GiveLikedPost []*GiveLikedPost `json:"likedPosts"`
	}
	if resp, err := s.client.Get(ctx, u, &result); err != nil {
		return nil, resp, err
	} else {
		return result.GiveLikedPost, resp, nil
	}
}

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/get-likes-discover/
func (s *LikesService) GetLikesDiscover(ctx context.Context, opt *GetLikesOptions) ([]*DiscoverLikedPost, *Response, error) {
	u, err := AddQueries("likes/discover", opt)
	if err != nil {
		return nil, nil, err
	}
	var result *struct {
		DiscoverLikedPost []*DiscoverLikedPost `json:"likedPosts"`
	}
	if resp, err := s.client.Get(ctx, u, &result); err != nil {
		return nil, resp, err
	} else {
		return result.DiscoverLikedPost, resp, nil
	}
}

type readReceivedLikesOptions struct {
	LikeId int `json:"likeId,omitempty"`
}

type ReadReceivedLikesResult struct {
	Like struct {
		Receive struct {
			HasUnread  bool `json:"hasUnread"`
			ReadLikeID int  `json:"readLikeId"`
		} `json:"receive"`
	} `json:"like"`
}

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/1/save-read-likes/
func (s *LikesService) ReadReceivedLikes(ctx context.Context, likeId int) (*ReadReceivedLikesResult, *Response, error) {
	u := "likes/receive/bookmark/save"

	var result *ReadReceivedLikesResult
	if resp, err := s.client.Post(ctx, u, &readReceivedLikesOptions{likeId}, result); err != nil {
		return nil, resp, err
	} else {
		return result, resp, nil
	}
}
