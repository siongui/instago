package instago

import (
	"testing"
)

func ExampleToplive(t *testing.T) {
	mgr, err := NewInstagramApiManager("auth.json")
	tl, err := mgr.Toplive()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(tl)
}
