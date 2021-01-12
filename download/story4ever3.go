package igdl

import (
	"log"
)

func (m *IGDownloadManager) TrayDownloader2(c chan TrayInfo, tl *TimeLimiter, verbose bool) {
	maxReelsMediaIds := 20
	queue := []TrayInfo{}
	for {
		select {
		case ti := <-c:
			// append to queue if not exist
			if IsTrayInfoInQueue(queue, ti) {
				if verbose {
					PrintTrayInfoMsg(ti, "exist. ignore.", "len(channel):", len(c), "len(queue):", len(queue))
				}
			} else {
				queue = append(queue, ti)
				if verbose {
					PrintTrayInfoMsg(ti, "appended.", "len(channel):", len(c), "len(queue):", len(queue))
				}
			}
		default:
			tis := []TrayInfo{}
			for len(queue) > 0 {
				ti := queue[0]
				queue = queue[1:]

				tis = append(tis, ti)

				if len(tis) == maxReelsMediaIds {
					break
				}
			}

			// delay download to reduce API access
			if len(tis) < maxReelsMediaIds {
				if len(tis) > 0 && tl.IsOverNIntervalAfterLastTime(2) {
					m.DownloadTrayInfos(tis, c, tl, true, true, verbose)
				} else {
					queue = append(queue, tis...)
				}
			} else {
				m.DownloadTrayInfos(tis, c, tl, true, true, verbose)
			}

			restInterval := 1
			if verbose {
				log.Println("TrayDownloader: current queue length is ", len(queue))
			}
			SleepSecond(restInterval)
		}
	}
}

func (m *IGDownloadManager) DownloadStoryForeverPublicReelMentions(interval1 int, interval2 int64, ignoreMuted, verbose bool) {
	// usually there are at most 150 trays in reels_tray.
	// double the buffer to 300. 160 or 200 may be ok as well.
	c := make(chan TrayInfo, 300)

	tl := NewTimeLimiter(interval2)
	go m.TrayDownloader2(c, tl, verbose)

	for {
		err := m.AccessReelsTrayOnce(c, ignoreMuted, verbose)
		if err != nil {
			log.Println(err)
		}
		PrintMsgSleep(interval1, "DownloadStoryForever: ")
	}
}
