package instago

import (
	"testing"
)

func ExampleGetMultipleReelsMedia(t *testing.T) {
	mgr, err := NewInstagramApiManager("auth.json")
	if err != nil {
		t.Error(err)
		return
	}

	userids := []string{"25025320", "1067259270"}
	trays, err := mgr.GetMultipleReelsMedia(userids)
	if err != nil {
		t.Error(err)
		return
	}

	for _, tray := range trays {
		t.Log(tray.Id)
		t.Log(tray.Items)
	}
}
