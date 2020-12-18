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

	fmt.Println("Download timeline")
	// This method will run forever and download posts in your timeline
	// every 15 seconds. The argument `1` means download only 1 page. You
	// can download more pages by changing this argument.
	mgr.DownloadTimeline(1)
}
