package main

import (
	"fmt"
	"os"

	"github.com/siongui/instago/download"
)

func main() {
	mgr, err := igdl.NewInstagramDownloadManager(
		os.Getenv("IG_DS_USER_ID"),
		os.Getenv("IG_SESSIONID"),
		os.Getenv("IG_CSRFTOKEN"))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Download stories and postlives of following users")
	// This method will run forever and donwload stories and postlives of
	// your following users every 30 seconds.
	mgr.DownloadStoryAndPostLive()
}
