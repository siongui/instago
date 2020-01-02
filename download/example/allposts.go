package main

import (
	"flag"
	"fmt"

	"github.com/siongui/instago/download"
)

func main() {
	mgr, err := igdl.NewInstagramDownloadManager("auth.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	id := flag.String("id", "instagram", "id of instagram user")
	flag.Parse()

	fmt.Println("Download all posts of a single user")
	// Given username, the following method will download all posts of the
	// user.
	mgr.DownloadAllPosts(*id)
}
