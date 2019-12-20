package instago

import (
	"testing"
)

func ExampleGetTimeline(t *testing.T) {
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
	t.Log(tl)
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
