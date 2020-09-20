package instago

import (
	"testing"
)

func ExampleTopsearch(t *testing.T) {
	mgr, err := NewInstagramApiManager("auth.json")
	tr, err := mgr.Topsearch("instagram")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(tr)
}
