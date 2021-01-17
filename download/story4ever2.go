package igdl

import (
	"fmt"
	"log"
	"strconv"
)

func (m *IGDownloadManager) AccessReelsTrayOnce2Chan(cPublicUser, cPrivateUser chan TrayInfo, ignoreMuted, verbose bool) (err error) {
	rt, err := m.GetReelsTray()
	if err != nil {
		log.Println(err)
		return
	}

	go PrintLiveBroadcasts(rt.Broadcasts)

	for index, tray := range rt.Trays {
		fmt.Print(index, ":")

		// 2: also download reel mentions in story item
		if tray.User.IsPrivate {
			ProcessTray(cPrivateUser, tray, 2, ignoreMuted, verbose)
		} else {
			ProcessTray(cPublicUser, tray, 2, ignoreMuted, verbose)
		}
	}

	return
}

// Under test.
func (m *IGDownloadManager) TwoAccountDownloadStoryForever(interval1 int, interval2, interval3 int64, ignoreMuted, verbose bool) {
	if !m.IsCleanAccountSet() {
		fmt.Println("clean account not set. exit")
		return
	}

	// usually there are at most 150 trays in reels_tray.
	// double the buffer to 300. 160 or 200 may be ok as well.
	cPublicUser := make(chan TrayInfo, 300)
	cPrivateUser := make(chan TrayInfo, 300)

	go m.GetCleanAccountManager().TrayDownloader(cPublicUser, NewTimeLimiter(interval2), true, verbose)
	go m.TrayDownloader2ChanPrivate(cPublicUser, cPrivateUser, NewTimeLimiter(interval3), true, verbose)

	for {
		err := m.AccessReelsTrayOnce2Chan(cPublicUser, cPrivateUser, ignoreMuted, verbose)
		if err != nil {
			log.Println(err)
		}
		PrintMsgSleep(interval1, "TwoAccountDownloadStoryForever: ")
	}
}

// Under test.
func (m *IGDownloadManager) TwoAccountDownloadStoryForeverSecondAccountViaStoryAPI(interval1 int, interval2, interval3 int64, ignoreMuted, verbose bool) {
	if !m.IsCleanAccountSet() {
		fmt.Println("clean account not set. exit")
		return
	}

	// usually there are at most 150 trays in reels_tray.
	// double the buffer to 300. 160 or 200 may be ok as well.
	cPublicUser := make(chan TrayInfo, 300)
	cPrivateUser := make(chan TrayInfo, 300)

	go m.GetCleanAccountManager().TrayDownloaderViaStoryAPI(cPublicUser, NewTimeLimiter(interval2), true, verbose)
	go m.TrayDownloader2ChanPrivate(cPublicUser, cPrivateUser, NewTimeLimiter(interval3), true, verbose)

	for {
		err := m.AccessReelsTrayOnce2Chan(cPublicUser, cPrivateUser, ignoreMuted, verbose)
		if err != nil {
			log.Println(err)
		}
		PrintMsgSleep(interval1, "TwoAccountDownloadStoryForever: ")
	}
}

func (m *IGDownloadManager) TrayDownloader2ChanPrivate(cPublicUser, cPrivateUser chan TrayInfo, tl *TimeLimiter, ignorePrivate, verbose bool) {
	maxReelsMediaIds := 20
	queue := []TrayInfo{}
	for {
		select {
		case ti := <-cPrivateUser:
			// append to queue if not exist
			if IsTrayInfoInQueue(queue, ti) {
				if verbose {
					PrintTrayInfoMsg(ti, "exist. ignore.", "len(cPrivateUser):", len(cPrivateUser), "len(queue):", len(queue))
				}
			} else {
				queue = append(queue, ti)
				if verbose {
					PrintTrayInfoMsg(ti, "appended.", "len(cPrivateUser):", len(cPrivateUser), "len(queue):", len(queue))
				}
			}
		default:
			tis := []TrayInfo{}
			for len(queue) > 0 {
				ti := queue[0]
				queue = queue[1:]

				/*
					if ignorePrivate && ti.IsPrivate {
						continue
					}
				*/

				tis = append(tis, ti)

				if len(tis) == maxReelsMediaIds {
					break
				}
			}

			// delay download to reduce API access
			if len(tis) < maxReelsMediaIds {
				if len(tis) > 0 && tl.IsOverNIntervalAfterLastTime(2) {
					m.DownloadTrayInfos(tis, cPublicUser, tl, false, ignorePrivate, verbose)
				} else {
					queue = append(queue, tis...)
				}
			} else {
				m.DownloadTrayInfos(tis, cPublicUser, tl, false, ignorePrivate, verbose)
			}

			restInterval := 1
			if verbose {
				log.Println("TrayDownloader2ChanPrivate: current queue length is ", len(queue))
			}
			SleepSecond(restInterval)
		}
	}
}

func (m *IGDownloadManager) TrayDownloaderViaStoryAPI(c chan TrayInfo, tl *TimeLimiter, ignorePrivateReelMention, verbose bool) {
	queue := []TrayInfo{}
	for {
		select {
		case ti := <-c:
			// append to queue if not exist
			if IsTrayInfoInQueue(queue, ti) {
				if verbose {
					PrintTrayInfoMsg(ti, "exist. ignore.", "len(c):", len(c), "len(queue):", len(queue))
				}
			} else {
				queue = append(queue, ti)
				if verbose {
					PrintTrayInfoMsg(ti, "appended.", "len(c):", len(c), "len(queue):", len(queue))
				}
			}
		default:
			if len(queue) > 0 {
				ti := queue[0]
				queue = queue[1:]

				// wait at least *interval* seconds until next private API access (prevent http 429)
				tl.WaitAtLeastIntervalAfterLastTime()
				ut, err := m.GetUserStory(strconv.FormatInt(ti.Id, 10))
				tl.SetLastTimeToNow()
				if err != nil {
					log.Println(err)
					c <- ti
				} else {
					tray := ut.Reel
					for _, item := range tray.GetItems() {
						err = ProcessTrayItem(c, item, ti, true, ignorePrivateReelMention, verbose)
						if err != nil {
							PrintUsernameIdMsg(ti.Username, ti.Id, err)
						}
					}
				}
			}

			restInterval := 1
			if verbose {
				log.Println("TrayDownloaderViaStoryAPI: current queue length is ", len(queue))
			}
			SleepSecond(restInterval)
		}
	}
}
