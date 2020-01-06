package main

import (
	"fmt"

	"github.com/siongui/instago/download"
)

func main() {
	mgr, err := igdl.NewInstagramDownloadManager("auth.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	users:= []string{"instagram", "google"}

	for _, user := range users {
		fmt.Println("Download all posts of", user)
		mgr.DownloadAllPosts(user)
	}
}
