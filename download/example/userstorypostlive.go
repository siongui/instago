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

	fmt.Println("Download unexpired (last 24 hours) stories and postlives of the user")
	// Given username, the following method will download unexpired stories
	// of the user.
	mgr.DownloadUserStoryPostliveByName(*id)
}
