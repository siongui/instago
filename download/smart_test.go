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

	mgr.DownloadPostNoLoginIfPossible(os.Getenv("CODE"))
}
