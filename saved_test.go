package instago

import (
	"fmt"
	"os"
	"testing"
)

func ExampleGetSavedPosts(t *testing.T) {
	mgr := NewInstagramApiManager(
		os.Getenv("IG_DS_USER_ID"),
		os.Getenv("IG_SESSIONID"),
		os.Getenv("IG_CSRFTOKEN"))

	items, err := mgr.GetSavedPosts()
	if err != nil {
		t.Error(err)
		return
	}

	for _, item := range items {
		fmt.Println(item.GetPostUrl())
	}
}
