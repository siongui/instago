package instago

import (
	"os"
	"testing"
)

func ExampleGetFollowing(t *testing.T) {
	mgr, err := NewInstagramApiManager("auth.json")
	if err != nil {
		t.Error(err)
		return
	}

	users, err := mgr.GetFollowing(os.Getenv("IG_TEST_ID"))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(users)
	t.Log(len(users))
}

func ExampleGetFollowers(t *testing.T) {
	mgr, err := NewInstagramApiManager("auth.json")
	if err != nil {
		t.Error(err)
		return
	}

	users, err := mgr.GetFollowers(os.Getenv("IG_TEST_ID"))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(users)
	t.Log(len(users))
}
