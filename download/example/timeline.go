package main

import (
	"fmt"
	"os"

	"github.com/siongui/instago/download"
)

func main() {
	mgr, err := igdl.NewInstagramDownloadManager(
		os.Getenv("IG_DS_USER_ID"),
		os.Getenv("IG_SESSIONID"),
		os.Getenv("IG_CSRFTOKEN"))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Download timeline")
	// This method will run forever and donwload posts in your timeline
	// every 15 seconds. The argument `1` means download only 1 page. If you
	// can download more pages by changing this argument.
	mgr.DownloadTimeline(1)
}
