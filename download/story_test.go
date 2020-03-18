package igdl

import (
	"fmt"
	"testing"
)

func ExampleDownloadUserStoryByNameLayer(t *testing.T) {
	mgr, err := NewInstagramDownloadManager("auth.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	mgr.DownloadUserStoryByNameLayer("instagram", 2)
}

func TestDownloadUserStoryLayer(t *testing.T) {
	mgr, err := NewInstagramDownloadManager("auth.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	mgr.DownloadUserStoryLayer(25025320, 2)
}
