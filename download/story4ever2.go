package igdl

import (
	"fmt"
	"log"
	"strconv"
)

var isFirstAccessReelsTray = true

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
func (m *IGDownloadManager) TwoAccountDownloadStoryForeverSecondAccountViaStoryAPI(interval1 int, interval2, interval3 int64, ignoreMuted, verbose bool) {
	if !m.IsCleanAccountSet() {
		fmt.Println("clean account not set. exit")
		return
	}

	// usually there are at most 150 trays in reels_tray.
	// double the buffer to 300. 160 or 200 may be ok as well.
	cPublicUser := make(chan TrayInfo, 300)
	cPrivateUser := make(chan TrayInfo, 300)

	go m.GetCleanAccountManager().TrayDownloaderViaStoryAPI(cPublicUser, NewTimeLimiter(interval2), true, true, verbose)
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

func (m *IGDownloadManager) TrayDownloaderViaStoryAPI(c chan TrayInfo, tl *TimeLimiter, stackMode, ignorePrivateReelMention, verbose bool) {
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
				if ti.LatestReelMedia > 0 && IsLatestReelMediaExist(ti.Username, ti.LatestReelMedia) {
					if verbose {
						PrintTrayInfoMsg(ti, "reelmedia already latest. ignore.", "len(c):", len(c), "len(queue):", len(queue))
					}
					break
				}
				queue = append(queue, ti)
				if verbose {
					PrintTrayInfoMsg(ti, "appended.", "len(c):", len(c), "len(queue):", len(queue))
				}
			}
		default:
			if len(queue) > 0 {

				if isFirstAccessReelsTray {
					// How do I reverse an array in Go?
					// https://stackoverflow.com/a/19239850
					for i, j := 0, len(queue)-1; i < j; i, j = i+1, j-1 {
						queue[i], queue[j] = queue[j], queue[i]
					}
					isFirstAccessReelsTray = false
				}

				ti := TrayInfo{}
				if stackMode {
					ti = queue[len(queue)-1]
					queue = queue[:len(queue)-1]
				} else {
					ti = queue[0]
					queue = queue[1:]
				}

				// wait at least *interval* seconds until next private API access (prevent http 429)
				tl.WaitAtLeastIntervalAfterLastTime()
				ut, err := m.GetUserStory(strconv.FormatInt(ti.Id, 10))
				tl.SetLastTimeToNow()
				if err != nil {
					log.Println(err)
					c <- ti
				} else {
					tray := ut.Reel
					ti.Username = tray.GetUsername()
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

func (m *IGDownloadManager) DownloadStoryFromUserIdFile(useridfile string, interval int64, verbose bool) (err error) {
	idstrs, err := ReadNonCommentLines(useridfile)
	if err != nil {
		return
	}

	// usually there are at most 150 trays in reels_tray.
	// double the buffer to 300. 160 or 200 may be ok as well.
	c := make(chan TrayInfo, 300)

	for _, idstr := range idstrs {
		id, err := strconv.ParseInt(idstr, 10, 64)
		if err != nil {
			return err
		}

		// layer = 2: also download reel mentions in story item
		c <- SetupTrayInfo(id, "", 2, false, 0)
	}

	m.TrayDownloaderViaStoryAPI(c, NewTimeLimiter(interval), true, true, verbose)

	return
}
