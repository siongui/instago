package instago

import (
	"os"
	"testing"
)

func ExampleGetUserStory(t *testing.T) {
	mgr, err := NewInstagramApiManager("auth.json")
	if err != nil {
		t.Error(err)
		return
	}

	ut, err := mgr.GetUserStory(os.Getenv("IG_TEST_ID"))
	if err != nil {
		t.Error(err)
		return
	}

	for _, item := range ut.Reel.Items {
		t.Log(item)
	}

	for _, bc := range ut.PostLiveItem.Broadcasts {
		t.Log(bc)
	}
}
