package igdl

import (
	"fmt"
	"time"
)

// download timeline until page n
func (m *IGDownloadManager) DownloadTimeline(n int) {
	sleepInterval := 12 // seconds

	for {
		items, err := m.apimgr.GetTimelineUntilPageN(n)
		if err != nil {
			fmt.Println(err)
		} else {
			for idx, item := range items {
				printTimelineItemInfo(idx, item)
				if !item.IsRegularMedia() {
					continue
				}
				m.GetPostItem(item)
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
