package instago

import (
	"testing"
)

func TestGetTimeline(t *testing.T) {
	mgr, err := NewInstagramApiManager("auth.json")
	if err != nil {
		t.Error(err)
		return
	}

	tl, err := mgr.GetTimeline()
	if err != nil {
		t.Error(err)
		return
	}

	for _, item := range tl.Items {
		if item.EndOfFeedDemarcator.Title != "" {
			t.Log(item.EndOfFeedDemarcator)
			if item.IsRegularMedia() == true {
				t.Error("IsRegularMedia() should be false")
			}
			continue
		}

		if item.Injected.Label != "" {
			t.Log(item.Injected.AdTitle)
			if item.IsRegularMedia() == true {
				t.Error("IsRegularMedia() should be false")
			}
			continue
		}

		if item.Type == 2 {
			t.Log(item.Suggestions)
			if item.IsRegularMedia() == true {
				t.Error("IsRegularMedia() should be false")
			}
			continue
		}

		t.Log(item.MediaType)
		if item.IsRegularMedia() == false {
			t.Error("IsRegularMedia() should be true")
		}
	}
}

func ExampleGetTimelineUntilPageN(t *testing.T) {
	mgr, err := NewInstagramApiManager("auth.json")
	if err != nil {
		t.Error(err)
		return
	}
	items, err := mgr.GetTimelineUntilPageN(5)
	if err != nil {
		t.Error(err)
		return
	}
	for _, item := range items {
		t.Log(item)
	}
}
