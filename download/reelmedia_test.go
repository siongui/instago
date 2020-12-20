package igdl

import (
	"testing"
)

func ExampleDownloadUserReelMediaByName(t *testing.T) {
	mgr, err := NewInstagramDownloadManager("auth.json")
	if err != nil {
		panic(err)
	}

	mgr.DownloadUserReelMediaByName("instagram")
}

func ExampleDownloadUserReelMediaLayer(t *testing.T) {
	mgr, err := NewInstagramDownloadManager("auth.json")
	if err != nil {
		panic(err)
	}

	mgr.DownloadUserReelMediaLayer("25025320", 2, 15)
}
