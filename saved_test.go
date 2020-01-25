package instago

import (
	"fmt"
	"testing"
)

func ExampleGetSavedPosts(t *testing.T) {
	mgr, err := NewInstagramApiManager("auth.json")
	if err != nil {
		t.Error(err)
		return
	}

	items, err := mgr.GetSavedPosts(20)
	if err != nil {
		t.Error(err)
		return
	}

	for _, item := range items {
		fmt.Println(item.GetPostUrl())
	}
}
