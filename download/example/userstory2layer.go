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

	username := flag.String("username", "instagram", "Instagram username")
	flag.Parse()

	fmt.Println("Download unexpired stories (last 24 hours) of the user and reel mentions")
	mgr.DownloadUserStoryByNameLayer(*username, 2, 12)
}
