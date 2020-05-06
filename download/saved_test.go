package igdl

import (
	"fmt"
	"testing"
)

func ExampleDownloadSavedPosts(t *testing.T) {
	mgr, err := NewInstagramDownloadManager("auth.json")
	if err != nil {
		t.Error(err)
		return
	}

	mgr.DownloadSavedPosts(12, true)
}

func ExampleIsInCollection(t *testing.T) {
	mgr, err := NewInstagramDownloadManager("auth.json")
	if err != nil {
		t.Error(err)
		return
	}
	mgr.LoadCollectionList()
	if err != nil {
		t.Error(err)
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

func ExampleCollectionId2Name(t *testing.T) {
	mgr, err := NewInstagramDownloadManager("auth.json")
	if err != nil {
		t.Error(err)
		return
	}
	mgr.LoadCollectionList()
	if err != nil {
		t.Error(err)
		return
	}

	items, err := mgr.apimgr.GetSavedPosts(10)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, item := range items {
		fmt.Println(item.GetPostUrl())
		for _, id := range item.SavedCollectionIds {
			fmt.Println(mgr.CollectionId2Name(id))
		}
	}
}

func ExampleCollectionName2Id(t *testing.T) {
	mgr, err := NewInstagramDownloadManager("auth.json")
	if err != nil {
		t.Error(err)
		return
	}
	mgr.LoadCollectionList()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(mgr.CollectionName2Id("MyCollectionName"))
}
