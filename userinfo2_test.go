package instago

import (
	"testing"
)

func ExampleGetUserInfoEndPoint(t *testing.T) {
	mgr, err := NewInstagramApiManager("auth.json")
	if err != nil {
		t.Error(err)
		return
	}

	user, err := mgr.GetUserInfoEndPoint("25025320")
	if err != nil {
		t.Error(err)
		return
	}

	if user.Username != "instagram" {
		t.Error(user)
		return
	}
	t.Log(user)
}
