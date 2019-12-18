package v1

import (
	"context"
	"time"

	"github.com/nulab/go-typetalk/typetalk/shared"
	"github.com/nulab/go-typetalk/v3/typetalk/internal"
)

// LikesService handles likes activity.
type LikesService service

// LikedPost represents a post and its likes.
type LikedPost struct {
	Post          *Post          `json:"post"`
	Likes         []*Like        `json:"likes"`
	DirectMessage *DirectMessage `json:"directMessage"`
}

// DiscoverLikedPost represents like in an organization.
type DiscoverLikedPost struct {
	*LikedPost
}

// ReceiveLikedPost represents received like.
type ReceiveLikedPost struct {
	*LikedPost
}

// GiveLikedPost represents given like.
type GiveLikedPost struct {
	*ReceiveLikedPost
	MyLike *MyLike `json:"myLike"`
}

// MyLike represents own like.
type MyLike struct {
	ID        int       `json:"id"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"createdAt"`
}

// GetLikesOptions represents parameters for likes API.
type GetLikesOptions struct {
	From int `json:"from,omitempty"`
}

// GetLikesReceive fetches received likes list.
//
// Typetalk API docs: https://developer.nulab.com/docs/typetalk/api/1/get-likes-receive/
func (s *LikesService) GetLikesReceive(ctx context.Context, opt *GetLikesOptions) ([]*ReceiveLikedPost, *shared.Response, error) {
	u, err := internal.AddQueries("likes/receive", opt)
	if err != nil {
		return nil, nil, err
	}
	var result *struct {
		LikedPosts []*ReceiveLikedPost `json:"likedPosts"`
	}
	resp, err := s.client.Get(ctx, u, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.LikedPosts, resp, nil
}

// GetLikesGive fetches given likes list. Those likes are given by your account.
//
// Typetalk API docs: https://developer.nulab.com/docs/typetalk/api/1/get-likes-give/
func (s *LikesService) GetLikesGive(ctx context.Context, opt *GetLikesOptions) ([]*GiveLikedPost, *shared.Response, error) {
	u, err := internal.AddQueries("likes/give", opt)
	if err != nil {
		return nil, nil, err
	}
	var result *struct {
		GiveLikedPost []*GiveLikedPost `json:"likedPosts"`
	}
	resp, err := s.client.Get(ctx, u, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.GiveLikedPost, resp, nil
}

// GetLikesDiscover fetches given likes list. Those likes are given by all the accounts.
//
// Typetalk API docs: https://developer.nulab.com/docs/typetalk/api/1/get-likes-discover/
func (s *LikesService) GetLikesDiscover(ctx context.Context, opt *GetLikesOptions) ([]*DiscoverLikedPost, *shared.Response, error) {
	u, err := internal.AddQueries("likes/discover", opt)
	if err != nil {
		return nil, nil, err
	}
	var result *struct {
		DiscoverLikedPost []*DiscoverLikedPost `json:"likedPosts"`
	}
	resp, err := s.client.Get(ctx, u, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.DiscoverLikedPost, resp, nil
}

type readReceivedLikesOptions struct {
	LikeID int `json:"likeId,omitempty"`
}

// ReadReceivedLikesResult represents a like that is marked as read.
type ReadReceivedLikesResult struct {
	Like struct {
		Receive struct {
			HasUnread  bool `json:"hasUnread"`
			ReadLikeID int  `json:"readLikeId"`
		} `json:"receive"`
	} `json:"like"`
}

// ReadReceivedLikes marks likes as read.
//
// Typetalk API docs: https://developer.nulab.com/docs/typetalk/api/1/save-read-likes/
func (s *LikesService) ReadReceivedLikes(ctx context.Context, likeID int) (*ReadReceivedLikesResult, *shared.Response, error) {
	u := "likes/receive/bookmark/save"

	var result *ReadReceivedLikesResult
	resp, err := s.client.Post(ctx, u, &readReceivedLikesOptions{likeID}, result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}
