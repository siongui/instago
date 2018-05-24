package main

import (
	"fmt"

	"github.com/siongui/instago/download"
)

func main() {
	mgr, err := igdl.NewInstagramDownloadManager(
		"IG_DS_USER_ID",
		"IG_SESSIONID",
		"IG_CSRFTOKEN")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Download all story highlights of all following users")
	// This method will run once and donwload all stories highlights of your
	// following users.
	mgr.DownloadStoryHighlights()
}
