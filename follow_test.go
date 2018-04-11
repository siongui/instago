package instago

import (
	"os"
	"testing"
)

func ExampleGetFollowing(t *testing.T) {
	mgr := NewInstagramApiManager(
		os.Getenv("IG_DS_USER_ID"),
		os.Getenv("IG_SESSIONID"),
		os.Getenv("IG_CSRFTOKEN"))
	users, err := mgr.GetFollowing(os.Getenv("IG_TEST_ID"))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(users)
	t.Log(len(users))
}

func ExampleGetFollowers(t *testing.T) {
	mgr := NewInstagramApiManager(
		os.Getenv("IG_DS_USER_ID"),
		os.Getenv("IG_SESSIONID"),
		os.Getenv("IG_CSRFTOKEN"))
	users, err := mgr.GetFollowers(os.Getenv("IG_TEST_ID"))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(users)
	t.Log(len(users))
}
