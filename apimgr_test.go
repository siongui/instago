package instago

import (
	"testing"
)

func TestGetSelfId(t *testing.T) {
	mgr, err := NewInstagramApiManager("auth.json")

	if err != nil {
		t.Error(err)
		return
	}

	if mgr.GetSelfId() == "" {
		t.Error("no ds_user_id")
		return
	}

	if _, ok := mgr.cookies["sessionid"]; !ok {
		t.Error("no sessionid")
		return
	}
	if _, ok := mgr.cookies["csrftoken"]; !ok {
		t.Error("no csrftoken")
		return
	}

	for k, v := range mgr.cookies {
		t.Log(k, v)
	}
}
