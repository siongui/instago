package igdl

import (
	"fmt"
	"os"
	"time"

	"github.com/siongui/instago"
)

func printPostLiveDownloadInfo(username, url, filepath string, timestamp int64) {
	fmt.Print("username: ")
	cc.Println(username)
	fmt.Print("time: ")
	cc.Println(formatTimestamp(timestamp))

	fmt.Print("Download ")
	rc.Print(url)
	fmt.Print(" to ")
	cc.Println(filepath)
}

func DownloadPostLive(pl instago.IGPostLive) {
	for _, item := range pl.PostLiveItems {
		for _, broadcast := range item.GetBroadcasts() {
			urls, err := broadcast.GetBaseUrls()
			if err != nil {
				fmt.Println(err)
				return
			}
			if len(urls) != 2 {
				fmt.Println("error: number of urls: ", len(urls))
				return
			}

			username := item.GetUsername()
			timestamp := broadcast.GetPublishedTime()
			filepath := ""
			vpath := ""
			apath := ""
			for index, url := range urls {
				if index == 0 {
					filepath = getPostLiveFilePath(
						username,
						url,
						"video",
						timestamp)
					vpath = filepath
				} else {
					filepath = getPostLiveFilePath(
						username,
						url,
						"audio",
						timestamp)
					apath = filepath
				}

				CreateFilepathDirIfNotExist(filepath)
				// check if file exist
				if _, err := os.Stat(filepath); os.IsNotExist(err) {
					// file not exists
					printPostLiveDownloadInfo(username, url, filepath, timestamp)
					err = Wget(url, filepath)
					if err != nil {
						fmt.Println(err)
						return
					}
				}
			}

			mpath := getPostLiveMergedFilePath(vpath, apath)
			// If video file is too big, and download is not yet
			// finished, next DownloadPostLive will create merged
			// file from unfinished downloaded files, which is not
			// correct.
			// check if file exist
			if _, err := os.Stat(mpath); err == nil {
				// file exists, delete it and create new one
				err = os.Remove(mpath)
				if err != nil {
					fmt.Println(err)
				}
				mergePostliveVideoAndAudio(vpath, apath)
			}
		}
	}
}

func DownloadStoryAndPostLive(mgr *instago.IGApiManager) {
	sleepInterval := 30 // seconds
	count := 0
	for {
		rt, err := mgr.GetReelsTray()
		if err != nil {
			fmt.Println(err)
		} else {
			go DownloadPostLive(rt.PostLive)
			if count == 0 {
				DownloadAllStory(rt.Trays, mgr)
				cc.Println("Download all stories finished")
			} else {
				DownloadUnreadStory(rt.Trays)
				cc.Println("Download unread stories finished")
			}
			count++
			count %= 5
		}

		rc.Print(time.Now().Format(time.RFC3339))
		fmt.Print(": sleep ")
		cc.Print(sleepInterval)
		fmt.Println(" second")
		time.Sleep(time.Duration(sleepInterval) * time.Second)
	}
}
