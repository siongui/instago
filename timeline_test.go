package instago

import (
	"os"
	"testing"
)

func ExampleGetTimeline(t *testing.T) {
	mgr := NewInstagramApiManager(
		os.Getenv("IG_DS_USER_ID"),
		os.Getenv("IG_SESSIONID"),
		os.Getenv("IG_CSRFTOKEN"))
	tl, err := mgr.GetTimeline()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(tl)
}

func ExampleGetTimelineUntilPageN(t *testing.T) {
	mgr := NewInstagramApiManager(
		os.Getenv("IG_DS_USER_ID"),
		os.Getenv("IG_SESSIONID"),
		os.Getenv("IG_CSRFTOKEN"))
	items, err := mgr.GetTimelineUntilPageN(5)
	if err != nil {
		t.Error(err)
		return
	}
	for _, item := range items {
		t.Log(item)
	}
}
