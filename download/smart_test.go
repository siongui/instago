package igdl

import (
	"os"
	"testing"

	"github.com/siongui/instago"
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

func ExampleSmartDownloadUserStoryPostliveLayer(t *testing.T) {
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

	user := instago.IGUser{Pk: 25025320, Username: "instagram", IsPrivate: false}
	err = mgr.SmartDownloadUserStoryPostliveLayer(user, 2)
	if err != nil {
		t.Error(err)
		return
	}
}
