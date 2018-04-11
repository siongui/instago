package instago

import (
	"os"
	"testing"
)

func ExampleGetUserStoryHighlights(t *testing.T) {
	mgr := NewInstagramApiManager(
		os.Getenv("IG_DS_USER_ID"),
		os.Getenv("IG_SESSIONID"),
		os.Getenv("IG_CSRFTOKEN"))
	trays, err := mgr.GetUserStoryHighlights(os.Getenv("IG_TEST_ID"))
	if err != nil {
		t.Error(err)
		return
	}

	//jsonPrettyPrint(trays)
	//t.Log(trays)
	for _, tray := range trays {
		t.Log(tray.GetTitle())
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
}
