package igdl

import (
	"fmt"
	"testing"
)

func TestDownloadPostLives(t *testing.T) {
	cpl := make(chan int)
	mgr, err := NewInstagramDownloadManager("auth.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	rt, err := mgr.apimgr.GetReelsTray()
	DownloadPostLive(rt.PostLive, cpl)
}
