package instago

import (
	"testing"
)

func TestTopsearch(t *testing.T) {
	mgr, err := NewInstagramApiManager("auth.json")
	tr, err := mgr.Topsearch("instagram")
	if err != nil {
		t.Error(err)
		return
	}
	for _, user := range tr.Users {
		t.Log(user)
	}
	for _, hashtag := range tr.Hashtags {
		t.Log(hashtag)
	}
	//t.Log(tr)
}
