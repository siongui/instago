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

	fmt.Println("Download stories and postlives of following users")
	// This method will run forever and download stories and postlives of
	// your following users every 30 seconds.
	mgr.DownloadStoryAndPostLive()
}
