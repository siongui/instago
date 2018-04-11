package main

import (
	"flag"
	"fmt"
	"github.com/siongui/instago"
	"github.com/siongui/instago/download"
	"os"
)

func main() {
	if !igdl.IsCommandAvailable("wget") {
		fmt.Println("Please install wget")
		return
	}
	if !igdl.IsCommandAvailable("ffmpeg") {
		fmt.Println("Please install ffmpeg")
		return
	}

	typ := flag.String("downloadtype", "timeline", "Download 1) timeline 2) story 3) highlight")
	flag.Parse()

	mgr := instago.NewInstagramApiManager(
		os.Getenv("IG_DS_USER_ID"),
		os.Getenv("IG_SESSIONID"),
		os.Getenv("IG_CSRFTOKEN"))

	switch *typ {
	case "timeline":
		fmt.Println("Download timeline")
		igdl.DownloadTimeline(mgr, 1)
	case "story":
		fmt.Println("Download Stories and Post lives")
		igdl.DownloadStoryAndPostLive(mgr)
	case "highlight":
		fmt.Println("Download all story highlights of all following users")
		igdl.DownloadStoryHighlights(mgr)
	default:
		fmt.Println("You have to choose a download type")
	}
}
