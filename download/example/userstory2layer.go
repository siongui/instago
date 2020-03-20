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

	fmt.Println("Download unexpired stories (last 24 hours) of the user and reel mentions")
	mgr.DownloadUserStoryByNameLayer(*id, 2)
}
