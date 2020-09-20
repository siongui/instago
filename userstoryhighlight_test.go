package instago

import (
	"testing"
)

func ExampleGetUserStoryHighlights(t *testing.T) {
	mgr, err := NewInstagramApiManager("auth.json")
	trays, err := mgr.GetUserStoryHighlights("25025320")
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
