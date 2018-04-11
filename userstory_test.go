package instago

import (
	"os"
	"testing"
)

func ExampleGetUserStory(t *testing.T) {
	mgr := NewInstagramApiManager(
		os.Getenv("IG_DS_USER_ID"),
		os.Getenv("IG_SESSIONID"),
		os.Getenv("IG_CSRFTOKEN"))
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
	}
}
