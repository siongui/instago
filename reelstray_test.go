package instago

import (
	"testing"
)

func ExampleGetReelsTray(t *testing.T) {
	mgr, err := NewInstagramApiManager("auth.json")
	if err != nil {
		t.Error(err)
		return
	}
	rt, err := mgr.GetReelsTray()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(rt)
}
