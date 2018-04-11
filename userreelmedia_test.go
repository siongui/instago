package instago

import (
	"os"
	"testing"
)

func ExampleGetUserReelMedia(t *testing.T) {
	mgr := NewInstagramApiManager(
		os.Getenv("IG_DS_USER_ID"),
		os.Getenv("IG_SESSIONID"),
		os.Getenv("IG_CSRFTOKEN"))
	urm, err := mgr.GetUserReelMedia(os.Getenv("IG_TEST_ID"))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(urm)
}
