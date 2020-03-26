package igdl

import (
	"fmt"
	"testing"
)

func ExampleDownloadSavedPosts(t *testing.T) {
	mgr, err := NewInstagramDownloadManager("auth.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	mgr.DownloadSavedPosts(12, true)
}

func ExampleIsInCollection(t *testing.T) {
	mgr, err := NewInstagramDownloadManager("auth.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	items, err := mgr.apimgr.GetSavedPosts(10)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, item := range items {
		fmt.Println(item.GetPostUrl())
		fmt.Println(mgr.IsInCollection(item, "Taiwan"))
	}
}
