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

	fmt.Println("Download story highlights of the user", *id)
	// Given username, the following method will download story highlights
	// of the user.
	mgr.DownloadUserStoryHighlightsByName(*id)
}
