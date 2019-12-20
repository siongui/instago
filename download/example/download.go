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

	typ := flag.String("downloadtype", "timeline", "Download 1) timeline 2) story 3) highlight 4) saved posts")
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
	case "saved":
		fmt.Println("Download saved posts")
		mgr.DownloadSavedPosts()
	default:
		fmt.Println("You have to choose a download type")
	}
}
