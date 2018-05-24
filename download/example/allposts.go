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

	fmt.Println("Download all posts of a single user")
	// Given username, the following method will download all posts of the
	// user.
	mgr.DownloadAllPosts("instagram")
}
