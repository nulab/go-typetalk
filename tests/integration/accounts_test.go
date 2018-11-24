package tests

import (
	"context"
	"testing"
)

func Test_V1_Accounts_GetMyProfile_should_get_own_profile(t *testing.T) {
	result, resp, err := clientV1.Accounts.GetMyProfile(context.Background())
	test(t, result, resp, err)
}

func Test_V3_Accounts_GetMyFriends_should_get_some_accounts(t *testing.T) {
	result, resp, err := clientV3.Accounts.GetMyFriends(context.Background(), spaceKey, "test", nil)
	test(t, result, resp, err)
}
