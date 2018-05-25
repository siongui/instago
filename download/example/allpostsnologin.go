package main

import (
	"fmt"

	"github.com/siongui/instago/download"
)

func main() {
	fmt.Println("Download all posts of a single user without login")
	// Given username, the following method will download all posts of the
	// user without login. (The user account must be public)
	igdl.DownloadAllPostsNoLogin("instagram")
}
