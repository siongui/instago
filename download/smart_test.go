package igdl

import (
	"os"

	"testing"
)

func ExampleDownloadPostNoLoginIfPossible(t *testing.T) {
	mgr, err := NewInstagramDownloadManager("auth.json")
	if err != nil {
		t.Error(err)
		return
	}

	err = mgr.LoadCleanDownloadManager("auth-clean.json")
	if err != nil {
		t.Error(err)
		return
	}

	_, err = mgr.DownloadPostNoLoginIfPossible(os.Getenv("CODE"))
	if err != nil {
		t.Error(err)
		return
	}
}

func ExampleDownloadAllPostsNoLoginIfPossible(t *testing.T) {
	mgr, err := NewInstagramDownloadManager("auth.json")
	if err != nil {
		t.Error(err)
		return
	}

	err = mgr.LoadCleanDownloadManager("auth-clean.json")
	if err != nil {
		t.Error(err)
		return
	}

	err = mgr.DownloadAllPostsNoLoginIfPossible(os.Getenv("USERNAME"))
	if err != nil {
		t.Error(err)
		return
	}
}
