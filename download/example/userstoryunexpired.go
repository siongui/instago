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

	fmt.Println("Download all unexpired stories")
	err = mgr.DownloadUnexpiredStoryOfAllFollowingUsers(2)
	if err != nil {
		panic(err)
	}
}
