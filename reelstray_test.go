package instago

import (
	"os"
	"testing"
)

func ExampleGetReelsTray(t *testing.T) {
	mgr := NewInstagramApiManager(
		os.Getenv("IG_DS_USER_ID"),
		os.Getenv("IG_SESSIONID"),
		os.Getenv("IG_CSRFTOKEN"))
	rt, err := mgr.GetReelsTray()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(rt)
}
