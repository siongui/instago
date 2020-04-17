package instago

import (
	"fmt"
	"os"
	"testing"
)

func ExampleGetSavedPosts(t *testing.T) {
	mgr, err := NewInstagramApiManager("auth.json")
	if err != nil {
		t.Error(err)
		return
	}

	items, err := mgr.GetSavedPosts(12)
	if err != nil {
		t.Error(err)
		return
	}

	for _, item := range items {
		err = PrintPostItem(&item)
		if err != nil {
			t.Error(err)
			return
		}
		fmt.Println(item.SavedCollectionIds)
		PrintTaggedUsers(item.Usertags)
		for _, cm := range item.CarouselMedia {
			PrintTaggedUsers(cm.Usertags)
		}
	}
}

func ExampleGetSavedCollection(t *testing.T) {
	mgr, err := NewInstagramApiManager("auth.json")
	if err != nil {
		t.Error(err)
		return
	}

	items, err := mgr.GetSavedCollection(os.Getenv("COLLECTION_ID"))
	if err != nil {
		t.Error(err)
		return
	}

	for _, item := range items {
		fmt.Println(item.GetPostUrl())
		fmt.Println(item.SavedCollectionIds)
	}
}

func ExampleGetSavedCollectionList(t *testing.T) {
	mgr, err := NewInstagramApiManager("auth.json")
	if err != nil {
		t.Error(err)
		return
	}

	collections, err := mgr.GetSavedCollectionList()
	if err != nil {
		t.Error(err)
		return
	}

	for _, collection := range collections {
		fmt.Println(collection)
	}
}
