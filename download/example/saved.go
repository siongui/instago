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

	num := flag.Int("num", 20, "number of saved posts to be downloaded")
	flag.Parse()

	fmt.Println("Download saved post:", *num)
	// The following method will download given number of saved posts.
	// -1 will download all saved posts.
	mgr.DownloadSavedPosts(*num)
}
