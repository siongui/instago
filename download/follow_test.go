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

	mgr.SaveSelfFollowers("myfollowers.json")
}
