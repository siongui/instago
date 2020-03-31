package igdl

import (
	"fmt"
	"github.com/siongui/instago"
	"time"
)

func printDownloadInfo(item instago.IGItem, username string, url, filepath string) {
	fmt.Print("username: ")
	cc.Println(username)
	fmt.Print("time: ")
	cc.Println(formatTimestamp(item.GetTimestamp()))
	fmt.Print("post url: ")
	cc.Println(item.GetPostUrl())

	fmt.Print("Download ")
	rc.Print(url)
	fmt.Print(" to ")
	cc.Println(filepath)
}

// download timeline until page n
func (m *IGDownloadManager) DownloadTimeline(n int) {
	sleepInterval := 12 // seconds

	for {
		items, err := m.apimgr.GetTimelineUntilPageN(n)
		if err != nil {
			fmt.Println(err)
		} else {
			for idx, item := range items {
				printSavedInfo(idx, &item)
				if !item.IsRegularMedia() {
					fmt.Println("seems to be ads. download ignored")
					continue
				}
				m.getPostItem(item)
			}
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
