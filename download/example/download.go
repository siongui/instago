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
	outputdir := flag.String("outputdir", "Instagram", "dir to save post and story")
	datadir := flag.String("datadir", "Data", "dir to save data")
	flag.Parse()

	igdl.SetOutputDir(*outputdir)
	igdl.SetDataDir(*datadir)

	switch *typ {
	case "timeline":
		fmt.Println("Download timeline")
		mgr.DownloadTimeline(1)
	case "story":
		fmt.Println("Download Stories and Post lives")
		mgr.DownloadStoryAndPostLiveForever(25, 2, false)
	case "highlight":
		fmt.Println("Download all story highlights of all following users")
		mgr.DownloadStoryHighlights()
	case "saved":
		fmt.Println("Download all saved posts")
		mgr.DownloadSavedPosts(-1, false)
	default:
		fmt.Println("You have to choose a download type")
	}
}
