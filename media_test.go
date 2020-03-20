package instago

import (
	"fmt"
	"os"
	"testing"
)

func TestGetMediaInfo(t *testing.T) {
	mgr, err := NewInstagramApiManager("auth.json")
	if err != nil {
		t.Error(err)
		return
	}

	item, err := mgr.GetMediaInfo(os.Getenv("IG_TEST_MEDIA_ID"))
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(item.SavedCollectionIds)
}
