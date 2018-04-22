package instago

import (
	"os"
	"testing"
)

func ExampleGetPostInfo(t *testing.T) {
	mgr := NewInstagramApiManager(
		os.Getenv("IG_DS_USER_ID"),
		os.Getenv("IG_SESSIONID"),
		os.Getenv("IG_CSRFTOKEN"))
	em, err := mgr.GetPostInfo(os.Getenv("IG_TEST_CODE"))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(em.GetPostUrl())
	t.Log(em.GetUsername())
	t.Log(em.GetTimestamp())
	for _, url := range em.GetMediaUrls() {
		t.Log(url)
	}
}

func ExampleGetPostInfoNoLogin(t *testing.T) {
	em, err := GetPostInfoNoLogin(os.Getenv("IG_TEST_CODE"))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(em.GetPostUrl())
	t.Log(em.GetUsername())
	t.Log(em.GetTimestamp())
	for _, url := range em.GetMediaUrls() {
		t.Log(url)
	}
}
