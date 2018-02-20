package tests

import (
	"context"
	"testing"
)

func Test_Likes_GetLikesGive_should_get_some_GiveLikedPost(t *testing.T) {
	result, resp, err := client.Likes.GetLikesGive(context.Background(), nil)
	test(t, result, resp, err)
}

func Test_Likes_GetLikesReceive_should_get_some_ReceiveLikedPost(t *testing.T) {
	result, resp, err := client.Likes.GetLikesReceive(context.Background(), nil)
	test(t, result, resp, err)
}

func Test_Likes_GetLikesDiscover_should_get_some_DiscoverLikedPost(t *testing.T) {
	result, resp, err := client.Likes.GetLikesDiscover(context.Background(), nil)
	test(t, result, resp, err)
}
