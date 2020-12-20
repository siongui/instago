package igdl

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/siongui/instago"
)

func (m *IGDownloadManager) GetStoryItemAndReelMentions2(item instago.IGItem, username string, interval int, getTime map[string]time.Time) (err error) {
	isDownloaded, err := getStoryItem(item, username)
	if err != nil {
		return
	}

	if item.Audience == "besties" {
		PrintBestiesItemInfo(item, username)
	}

	fetched := make(map[string]bool)
	if isDownloaded {
		for _, rm := range item.ReelMentions {
			PrintReelMentionInfo(rm)
			if rm.GetUsername() == username {
				fmt.Println("reel mention self. download ignored.")
				continue
			}

			if rm.IsPublic() {
				if _, ok := fetched[rm.GetUsername()]; !ok {
					d := time.Now().Sub(getTime["last"])
					for d < time.Duration(interval)*time.Second {
						time.Sleep(time.Duration(interval)*time.Second - d)
						d = time.Now().Sub(getTime["last"])
					}
					m.downloadUserReelMedia(rm.GetUserId())
					getTime["last"] = time.Now()
					fetched[rm.GetUsername()] = true
				}
			}
		}
	}
	return
}

func (m *IGDownloadManager) DownloadZeroItemUsers2(c chan instago.IGReelTray, interval int, getTime map[string]time.Time, verbose bool) {
	queue := []instago.IGReelTray{}
	for {
		select {
		case tray := <-c:
			// append to queue if not exist
			id := tray.Id
			username := tray.GetUsername()
			if verbose {
				UsernameIdColorPrint(username, id)
				fmt.Println("legnth of channel:", len(c))
			}
			if isTrayInQueue(queue, tray) {
				if verbose {
					PrintUsernameIdMsg(username, id, "exist. ignore.")
				}
			} else {
				queue = append(queue, tray)
				if verbose {
					PrintUsernameIdMsg(username, id, "appended")
				}
			}
		default:
			if len(queue) > 0 {
				tray := queue[0]
				queue = queue[1:]

				id := strconv.FormatInt(tray.Id, 10)
				username := tray.GetUsername()
				if verbose {
					PrintUsernameIdMsg(username, id, " downloading...")
				}

				go func() {
					// FIXME: take besties into account
					d := time.Now().Sub(getTime["last"])
					for d < time.Duration(interval)*time.Second {
						time.Sleep(time.Duration(interval)*time.Second - d)
						d = time.Now().Sub(getTime["last"])
					}
					ut, err := m.GetUserReelMedia(id)
					getTime["last"] = time.Now()
					if err != nil {
						PrintUsernameIdMsg(username, id, err)
						queue = append(queue, tray)
						return
					}

					for _, item := range ut.GetItems() {
						err = m.GetStoryItemAndReelMentions2(item, ut.GetUsername(), interval, getTime)
						if err != nil {
							PrintUsernameIdMsg(username, id, err)
							queue = append(queue, tray)
							return
						}
					}
				}()
			}

			if verbose {
				fmt.Println("current queue length: ", len(queue))
				PrintMsgSleep(interval, "DownloadZeroItemUsers: ")
			} else {
				SleepSecond(interval)
			}
		}
	}
}

// Try (25, 15, true, false). If http 429 happens, increase the number.
func (m *IGDownloadManager) DownloadStoryAndPostLiveForever2(interval1, interval2 int, ignoreMuted, verbose bool) {
	// channel for waiting DownloadPostLive completed
	isDownloading := make(map[string]bool)

	// usually there are at most 150 trays in reels_tray.
	// double the buffer to 300. 160 or 200 may be ok as well.
	c := make(chan instago.IGReelTray, 300)

	getTime := make(map[string]time.Time)
	getTime["last"] = time.Now()
	go m.DownloadZeroItemUsers2(c, interval2, getTime, verbose)
	for {
		rt, err := m.GetReelsTray()
		if err != nil {
			log.Println(err)
			continue
		}

		// TODO: use channel for postlive?
		go DownloadPostLive(rt.PostLive, isDownloading)
		go PrintLiveBroadcasts(rt.Broadcasts)

		//for index, tray := range rt.Trays {
		for _, tray := range rt.Trays {
			username := tray.GetUsername()
			id := tray.Id
			//items := tray.GetItems()

			if ignoreMuted && tray.Muted {
				if verbose {
					PrintUsernameIdMsg(username, id, " is muted && ignoreMuted set. no download")
				}
				continue
			}

			if isLatestReelMediaDownloaded(username, tray.LatestReelMedia) {
				if verbose {
					PrintUsernameIdMsg(username, id, " all downloaded")
				}
				continue
			}

			if tray.HasBestiesMedia {
				PrintUsernameIdMsg(username, id, "has close friend (besties) story item(s)")
			}

			if verbose {
				UsernameIdColorPrint(username, id)
				fmt.Println(" has undownloaded items")
			}
			c <- tray
		}

		PrintMsgSleep(interval1, "DownloadStoryAndPostLiveForever: ")
	}
}
