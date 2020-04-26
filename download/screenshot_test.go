package igdl

import (
	"testing"
)

func TestGetUserDataDirChromiumSnap(t *testing.T) {
	dir, _ := GetUserDataDirChromiumSnap("instagram")
	if dir != "/home/instagram/snap/chromium/current/.config/chromium" {
		t.Error(dir)
		return
	}
}
