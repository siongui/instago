package instago

import (
	"os"
	"testing"
)

func ExampleGetUserStory(t *testing.T) {
	mgr, err := NewInstagramApiManager("auth.json")
	if err != nil {
		t.Error(err)
		return
	}
	tray, err := mgr.GetUserStory(os.Getenv("IG_TEST_ID"))
	if err != nil {
		t.Error(err)
		return
	}

	//jsonPrettyPrint(tray)
	//t.Log(tray)
	for _, item := range tray.GetItems() {
		urls, err := item.GetMediaUrls()
		if err != nil {
			t.Error(err)
			continue
		}
		for _, url := range urls {
			t.Log(url)
		}
		for _, rm := range item.ReelMentions {
			t.Log(rm.GetUsername(), rm.GetUserId())
		}
	}
}
