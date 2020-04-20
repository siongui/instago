package igdl

import (
	"testing"
)

func ExampleSaveSelfFollowers(t *testing.T) {
	mgr, err := NewInstagramDownloadManager("auth.json")
	if err != nil {
		t.Error(err)
		return
	}

	mgr.SaveSelfFollowers(mgr.GetSelfId() + "-followers.json")
}

func ExampleSaveSelfFollowing(t *testing.T) {
	mgr, err := NewInstagramDownloadManager("auth.json")
	if err != nil {
		t.Error(err)
		return
	}

	mgr.SaveSelfFollowing(mgr.GetSelfId() + "-following.json")
}

func ExampleLoadFollowUsers(t *testing.T) {
	users, err := LoadFollowUsers("following-or-followers.json")
	if err != nil {
		t.Error(err)
		return
	}

	for _, user := range users {
		t.Log(user)
	}
}

func ExampleSaveSelfFollow(t *testing.T) {
	mgr, err := NewInstagramDownloadManager("auth.json")
	if err != nil {
		t.Error(err)
		return
	}

	mgr.SaveSelfFollow()
}
