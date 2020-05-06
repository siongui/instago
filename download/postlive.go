package igdl

import (
	"errors"
	"fmt"
	"os"

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

func DownloadPostLiveItem(pli instago.IGPostLiveItem) (err error) {
	for _, broadcast := range pli.GetBroadcasts() {
		urls, err := broadcast.GetBaseUrls()
		if err != nil {
			fmt.Println(err)
			return err
		}

		// total 5 or 2 urls.
		// if 5, first 4 urls are video track. last url is audio track.
		// second url seems to be best video quality track.
		vidx := 0
		if len(urls) == 2 {
			//fmt.Println("number of urls = 2")
			vidx = 0
		} else if len(urls) == 5 {
			//fmt.Println("number of urls = 5")
			vidx = 1
		} else {
			fmt.Println("error: number of urls != (5 or 2)", len(urls))
			return errors.New("error: number of urls != (5 or 2)")
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
					return err
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
	return
}

func DownloadPostLive(pl instago.IGPostLive, isDownloading map[string]bool) {
	for _, item := range pl.PostLiveItems {
		if _, ok := isDownloading[item.Pk]; ok {
			fmt.Println(item.Pk, " is downloading. ignored.")
			continue
		}
		isDownloading[item.Pk] = true
		PrintPostLiveItem(item)
		DownloadPostLiveItem(item)
		delete(isDownloading, item.Pk)
	}
}

func (m *IGDownloadManager) DownloadStoryAndPostLive() {
	// channel for waiting DownloadPostLive completed
	isDownloading := make(map[string]bool)

	sleepInterval := 30 // seconds
	count := 0
	for {
		rt, err := m.apimgr.GetReelsTray()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go DownloadPostLive(rt.PostLive, isDownloading)
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

		SleepAndPrint(sleepInterval)
	}
}
