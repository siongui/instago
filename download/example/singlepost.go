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

	code := flag.String("code", "B6qsBPtgSFx", "code of the post")
	flag.Parse()

	fmt.Println("Download single post")
	// Given username, the following method will download all posts of the
	// user.
	mgr.DownloadPost(*code)
}
