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

	usernames := []string{"google", "facebook", "instagram"}

	for _, username := range usernames {
		fmt.Println("Download unexpired stories of", username)
		err = mgr.DownloadUserStoryByName(username)
		if err != nil {
			fmt.Println(username)
			panic(err)
		}
	}
}
