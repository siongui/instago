package igdl

import (
	"os"
	"testing"
)

func ExampleDownloadPost(t *testing.T) {
	mgr, err := NewInstagramDownloadManager("auth.json")
	if err != nil {
		t.Error(err)
		return
	}

	mgr.DownloadPost(os.Getenv("IG_TEST_CODE"))
}

func ExampleDownloadPostNoLogin(t *testing.T) {
	DownloadPostNoLogin(os.Getenv("IG_TEST_CODE"))
}

func ExampleDownloadRecentPostsNoLogin(t *testing.T) {
	DownloadRecentPostsNoLogin(os.Getenv("IG_TEST_USERNAME"))
}
