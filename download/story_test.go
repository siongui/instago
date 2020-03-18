package igdl

import (
	"fmt"
	"testing"
)

func TestDownloadUserStoryByNameLayer(t *testing.T) {
	mgr, err := NewInstagramDownloadManager("auth.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	mgr.DownloadUserStoryByNameLayer("instagram", 2)
}
