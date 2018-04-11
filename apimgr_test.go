package instago

import (
	"os"
	"testing"
)

func TestGetSelfId(t *testing.T) {
	mgr := NewInstagramApiManager(
		os.Getenv("IG_DS_USER_ID"),
		os.Getenv("IG_SESSIONID"),
		os.Getenv("IG_CSRFTOKEN"))

	if mgr.GetSelfId() != os.Getenv("IG_DS_USER_ID") {
		t.Error(mgr.GetSelfId())
	}
}
