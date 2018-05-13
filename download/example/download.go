package main

import (
	"flag"
	"fmt"
	"github.com/siongui/instago/download"
	"os"
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

	typ := flag.String("downloadtype", "timeline", "Download 1) timeline 2) story 3) highlight")
	flag.Parse()

	switch *typ {
	case "timeline":
		fmt.Println("Download timeline")
		mgr.DownloadTimeline(1)
	case "story":
		fmt.Println("Download Stories and Post lives")
		mgr.DownloadStoryAndPostLive()
	case "highlight":
		fmt.Println("Download all story highlights of all following users")
		mgr.DownloadStoryHighlights()
	default:
		fmt.Println("You have to choose a download type")
	}
}
