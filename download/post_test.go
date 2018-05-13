package igdl

import (
	"os"
	"testing"
)

func ExampleDownloadPost(t *testing.T) {
	mgr, err := NewInstagramDownloadManager(
		os.Getenv("IG_DS_USER_ID"),
		os.Getenv("IG_SESSIONID"),
		os.Getenv("IG_CSRFTOKEN"))
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
