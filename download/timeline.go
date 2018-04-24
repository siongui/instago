package igdl

import (
	"fmt"
	"github.com/siongui/instago"
	"os"
	"time"
)

func printDownloadInfo(item instago.IGItem, url, filepath string) {
	fmt.Print("username: ")
	cc.Println(item.GetUsername())
	fmt.Print("time: ")
	cc.Println(formatTimestamp(item.GetTimestamp()))
	fmt.Print("post url: ")
	cc.Println(item.GetPostUrl())

	fmt.Print("Download ")
	rc.Print(url)
	fmt.Print(" to ")
	cc.Println(filepath)
}

func getTimelineItems(items []instago.IGItem) {
	for _, item := range items {
		if !item.IsRegularMedia() {
			continue
		}

		urls, err := item.GetMediaUrls()
		if err != nil {
			fmt.Println(err)
			continue
		}
		for index, url := range urls {
			filepath := getPostFilePath(
				item.GetUsername(),
				item.GetUserId(),
				url,
				item.GetTimestamp())
			if index > 0 {
				filepath = appendIndexToFilename(filepath, index)
			}

			CreateFilepathDirIfNotExist(filepath)
			// check if file exist
			if _, err := os.Stat(filepath); os.IsNotExist(err) {
				// file not exists
				printDownloadInfo(item, url, filepath)
				err = Wget(url, filepath)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}

// download timeline until page n
func DownloadTimeline(mgr *instago.IGApiManager, n int) {
	sleepInterval := 15 // seconds

	for {
		items, err := mgr.GetTimelineUntilPageN(n)
		if err != nil {
			fmt.Println(err)
		} else {
			getTimelineItems(items)
		}

		// sleep for a while
		fmt.Println("=========================")
		rc.Print(time.Now().Format(time.RFC3339))
		fmt.Print(": sleep ")
		cc.Print(sleepInterval)
		fmt.Println(" second")
		fmt.Println("=========================")
		time.Sleep(time.Duration(sleepInterval) * time.Second)
	}
}
