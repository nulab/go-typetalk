package tests

import (
	"context"
	"testing"
)

func Test_V1_Likes_GetLikesGive_should_get_some_GiveLikedPost(t *testing.T) {
	result, resp, err := clientV1.Likes.GetLikesGive(context.Background(), nil)
	test(t, result, resp, err)
}

func Test_V1_Likes_GetLikesReceive_should_get_some_ReceiveLikedPost(t *testing.T) {
	result, resp, err := clientV1.Likes.GetLikesReceive(context.Background(), nil)
	test(t, result, resp, err)
}

func Test_V1_Likes_GetLikesDiscover_should_get_some_DiscoverLikedPost(t *testing.T) {
	result, resp, err := clientV1.Likes.GetLikesDiscover(context.Background(), nil)
	test(t, result, resp, err)
}


func Test_V2_Likes_GetLikesGive_should_get_some_GiveLikedPost(t *testing.T) {
	result, resp, err := clientV2.Likes.GetLikesGive(context.Background(), spaceKey, nil)
	test(t, result, resp, err)
}

func Test_V2_Likes_GetLikesReceive_should_get_some_ReceiveLikedPost(t *testing.T) {
	result, resp, err := clientV2.Likes.GetLikesReceive(context.Background(), spaceKey,nil)
	test(t, result, resp, err)
}

func Test_V2_Likes_GetLikesDiscover_should_get_some_DiscoverLikedPost(t *testing.T) {
	result, resp, err := clientV2.Likes.GetLikesDiscover(context.Background(), spaceKey,nil)
	test(t, result, resp, err)
}