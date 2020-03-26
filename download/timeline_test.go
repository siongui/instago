package igdl

import (
	"fmt"
	"testing"
)

func ExampleDownloadTimeline(t *testing.T) {
	mgr, err := NewInstagramDownloadManager("auth.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	mgr.DownloadTimeline(1)
}
