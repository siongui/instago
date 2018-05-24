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

	fmt.Println("Download unexpired stories (last 24 hours) of the user")
	// Given username, the following method will download unexpired stories
	// of the user.
	mgr.DownloadUserStoryByName("instagram")
}
