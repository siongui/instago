package instago

import (
	"os"
	"testing"
)

func ExampleTopsearch(t *testing.T) {
	mgr := NewInstagramApiManager(
		os.Getenv("IG_DS_USER_ID"),
		os.Getenv("IG_SESSIONID"),
		os.Getenv("IG_CSRFTOKEN"))
	tr, err := mgr.Topsearch("instagram")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(tr)
}
