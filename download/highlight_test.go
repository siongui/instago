package igdl

import (
	"fmt"
	"testing"
)

func ExampleDownloadUserStoryHighlightsByName(t *testing.T) {
	mgr, err := NewInstagramDownloadManager("auth.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	mgr.DownloadUserStoryHighlightsByName("instagram")
}
