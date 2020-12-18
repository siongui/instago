package main

import (
	"fmt"

	"github.com/siongui/instago/download"
)

func main() {
	mgr, err := igdl.NewInstagramDownloadManager("auth.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Download all story highlights of all following users")
	// This method will run once and download all stories highlights of your
	// following users.
	mgr.DownloadStoryHighlights()
}
