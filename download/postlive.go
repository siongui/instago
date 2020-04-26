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

func DownloadPostLiveItem(pli instago.IGPostLiveItem) {
	for _, broadcast := range pli.GetBroadcasts() {
		urls, err := broadcast.GetBaseUrls()
		if err != nil {
			fmt.Println(err)
			return
		}

		// total 5 or 2 urls.
		// if 5, first 4 urls are video track. last url is audio track.
		// second url seems to be best video quality track.
		vidx := 0
		if len(urls) == 2 {
			fmt.Println("number of urls = 2")
			vidx = 0
		} else if len(urls) == 5 {
			fmt.Println("number of urls = 5")
			vidx = 1
		} else {
			fmt.Println("error: number of urls != (5 or 2)", len(urls))
			return
		}

		username := pli.GetUsername()
		id := pli.GetUserId()
		timestamp := broadcast.GetPublishedTime()
		filepath := ""
		vpath := ""
		apath := ""
		for index, url := range urls {
			if index == vidx {
				filepath = getPostLiveFilePath(
					username,
					id,
					url,
					"video",
					timestamp)
				vpath = filepath
			} else if index == (len(urls) - 1) {
				filepath = getPostLiveFilePath(
					username,
					id,
					url,
					"audio",
					timestamp)
				apath = filepath
			} else {
				continue
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
		// check if file exist
		if _, err := os.Stat(mpath); os.IsNotExist(err) {
			// file not exists
			mergePostliveVideoAndAudio(vpath, apath)
		}
	}
}

func DownloadPostLive(pl instago.IGPostLive, cpl chan int) {
	defer func() { cpl <- 1 }()

	for _, item := range pl.PostLiveItems {
		go PrintPostLiveItem(item)
		DownloadPostLiveItem(item)
	}
}

func (m *IGDownloadManager) DownloadStoryAndPostLive() {
	// channel for waiting DownloadPostLive completed
	cpl := make(chan int)

	sleepInterval := 30 // seconds
	count := 0
	for {
		rt, err := m.apimgr.GetReelsTray()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go DownloadPostLive(rt.PostLive, cpl)
		go PrintLiveBroadcasts(rt.Broadcasts)
		if count == 0 {
			m.DownloadAllStory(rt.Trays)
			cc.Println("Download all stories finished")
		} else {
			DownloadUnreadStory(rt.Trays)
			cc.Println("Download unread stories finished")
		}
		count++
		count %= 5

		rc.Print(time.Now().Format(time.RFC3339))
		fmt.Print(": sleep ")
		cc.Print(sleepInterval)
		fmt.Println(" second")
		time.Sleep(time.Duration(sleepInterval) * time.Second)

		// wait DownloadPostLive completed
		<-cpl
	}
}
