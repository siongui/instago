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

	fmt.Println("Download unexpired stories (last 24 hours) of the user", *id)
	// Given user id, the following method will download unexpired stories
	// of the user.
	mgr.DownloadUserStory(*id)

	// mgr.DownloadUserStoryByName can also be used to download user
	// unexpired stories by username.
}
