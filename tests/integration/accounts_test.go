package tests

import (
	"context"
	"testing"
)

func Test_Accounts_GetMyProfile_should_get_own_profile(t *testing.T) {
	result, resp, err := client.Accounts.GetMyProfile(context.Background())
	test(t, result, resp, err)
}
