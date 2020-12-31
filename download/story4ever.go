package igdl

import (
	"fmt"
	"log"
	"strconv"

	"github.com/siongui/instago"
)

func ProcessTrayItem(c chan TrayInfo, item instago.IGItem, ti TrayInfo, ignoreReelMentionsIfStoryItemExist, ignorePrivateReelMention, verbose bool) (err error) {
	isDownloaded, err := getStoryItem(item, ti.Username)
	if err != nil {
		return
	}

	if ti.Layer-1 < 1 {
		return
	}

	if !isDownloaded && ignoreReelMentionsIfStoryItemExist {
		return
	}

	for _, rm := range item.ReelMentions {
		PrintReelMentionInfo(rm)
		if rm.User.Pk == ti.Id {
			continue
		}
		if rm.User.IsPrivate && ignorePrivateReelMention {
			continue
		}
		c <- setupTrayInfo(rm.User.Pk, rm.GetUsername(), ti.Layer-1, rm.User.IsPrivate)
		if verbose {
			PrintUsernameIdMsg(rm.GetUsername(), rm.User.Pk, "sent to channel (reel mention)")
		}
	}

	return
}

func ProcessTray(c chan TrayInfo, tray instago.IGReelTray, layer int64, ignoreMuted, verbose bool) {
	username := tray.GetUsername()
	id := tray.Id
	//items := tray.GetItems()

	if ignoreMuted && tray.Muted {
		if verbose {
			PrintUsernameIdMsg(username, id, " is muted && ignoreMuted set. no download")
		}
		return
	}

	if IsLatestReelMediaExist(username, tray.LatestReelMedia) {
		if verbose {
			PrintUsernameIdMsg(username, id, " all downloaded")
		}
		return
	}

	if tray.HasBestiesMedia {
		PrintUsernameIdMsg(username, id, " has close friend (besties) story item(s)")
	}

	if verbose {
		PrintUsernameIdMsg(username, id, " has undownloaded items")
	}

	// layer = 2: also download reel mentions in story item
	c <- setupTrayInfo(id, username, layer, tray.User.IsPrivate)
	/*
		items := tray.GetItems()
		if len(items) > 0 {
			for _, item := range items {
				_, err = GetStoryItem(item, username)
				if err != nil {
					PrintUsernameIdMsg(username, id, err)
				}
			}
		} else {
			multipleIds = append(multipleIds, strconv.FormatInt(id, 10))
		}

		if len(multipleIds) > maxReelsMediaIds {
			break
		}
	*/
}

func (m *IGDownloadManager) DownloadTrayInfos(tis []TrayInfo, c chan TrayInfo, tl *TimeLimiter, ignoreReelMentionsIfStoryItemExist, ignorePrivateReelMention, verbose bool) {
	downloadIds := []string{}
	for _, ti := range tis {
		id := strconv.FormatInt(ti.Id, 10)
		username := ti.Username
		if verbose {
			PrintUsernameIdMsg(username, id, " to be batch downloaded")
		}
		downloadIds = append(downloadIds, id)
	}

	// wait at least *interval* seconds until next private API access
	tl.WaitAtLeastIntervalAfterLastTime()
	// get stories of multiple users at one API access
	trays, err := m.GetMultipleReelsMedia(downloadIds)
	tl.SetLastTimeToNow()
	if err != nil {
		log.Println(err)
		// sent back to channel to re-download
		for _, ti := range tis {
			if verbose {
				PrintUsernameIdMsg(ti.Username, ti.Id, " sent back to channel to re-download")
			}
			c <- ti
		}
		return
	}

	// we have story trays of multiple users now. start to download
	for _, tray := range trays {
		ti, ok := GetTrayInfoFromQueue(tis, tray.Id)
		if !ok {
			log.Println("cannot find", tray.Id, "in queue (impossible to happen)")
			continue
		}

		for _, item := range tray.Items {
			err = ProcessTrayItem(c, item, ti, ignoreReelMentionsIfStoryItemExist, ignorePrivateReelMention, verbose)
			if err != nil {
				PrintUsernameIdMsg(ti.Username, ti.Id, err)
			}
		}
	}
}

func (m *IGDownloadManager) TrayDownloader(c chan TrayInfo, tl *TimeLimiter, ignorePrivate, verbose bool) {
	maxReelsMediaIds := 20
	queue := []TrayInfo{}
	for {
		select {
		case ti := <-c:
			// append to queue if not exist
			id := ti.Id
			username := ti.Username
			if IsTrayInfoInQueue(queue, ti) {
				if verbose {
					PrintUsernameIdMsg(username, id, "exist. ignore.", "len(channel):", len(c), "len(queue):", len(queue))
				}
			} else {
				queue = append(queue, ti)
				if verbose {
					PrintUsernameIdMsg(username, id, "appended.", "len(channel):", len(c), "len(queue):", len(queue))
				}
			}
		default:
			tis := []TrayInfo{}
			for len(queue) > 0 {
				ti := queue[0]
				queue = queue[1:]

				if ignorePrivate && ti.IsPrivate {
					continue
				}

				tis = append(tis, ti)

				if len(tis) == maxReelsMediaIds {
					break
				}
			}

			// delay download to reduce API access
			if len(tis) < maxReelsMediaIds {
				if len(tis) > 0 && tl.IsOverNIntervalAfterLastTime(2) {
					m.DownloadTrayInfos(tis, c, tl, false, ignorePrivate, verbose)
				} else {
					queue = append(queue, tis...)
				}
			} else {
				m.DownloadTrayInfos(tis, c, tl, false, ignorePrivate, verbose)
			}

			restInterval := 1
			if verbose {
				log.Println("TrayDownloader: current queue length is ", len(queue))
			}
			SleepSecond(restInterval)
		}
	}
}

func (m *IGDownloadManager) AccessReelsTrayOnce(c chan TrayInfo, ignoreMuted, verbose bool) (err error) {
	rt, err := m.GetReelsTray()
	if err != nil {
		log.Println(err)
		return
	}

	go PrintLiveBroadcasts(rt.Broadcasts)

	for index, tray := range rt.Trays {
		fmt.Print(index, ":")
		// layer = 2: also download reel mentions in story item
		ProcessTray(c, tray, 2, ignoreMuted, verbose)
	}

	return
}

// DownloadStoryForever downloads reels tray periodically. interval1 is the
// interval for access to reels tray API. interval2 is the interval for fetching
// user stories. ignoreMute will ignore stories of muted users if true. verbose
// will print more info if true. If not sure, try (90, 60, true, true). If http
// 429 happens, try to use longer interval.
func (m *IGDownloadManager) DownloadStoryForever(interval1 int, interval2 int64, ignoreMuted, verbose bool) {
	// usually there are at most 150 trays in reels_tray.
	// double the buffer to 300. 160 or 200 may be ok as well.
	c := make(chan TrayInfo, 300)

	tl := NewTimeLimiter(interval2)
	go m.TrayDownloader(c, tl, false, verbose)

	for {
		err := m.AccessReelsTrayOnce(c, ignoreMuted, verbose)
		if err != nil {
			log.Println(err)
		}
		PrintMsgSleep(interval1, "DownloadStoryForever: ")
	}
}

func (m *IGDownloadManager) DownloadStoryForeverViaCleanAccount(interval1 int, interval2 int64, ignoreMuted, verbose bool) {
	if !m.IsCleanAccountSet() {
		fmt.Println("clean account not set. exit")
		return
	}

	// usually there are at most 150 trays in reels_tray.
	// double the buffer to 300. 160 or 200 may be ok as well.
	c := make(chan TrayInfo, 300)

	tl := NewTimeLimiter(interval2)
	go m.GetCleanAccountManager().TrayDownloader(c, tl, true, verbose)

	for {
		err := m.AccessReelsTrayOnce(c, ignoreMuted, verbose)
		if err != nil {
			log.Println(err)
		}
		PrintMsgSleep(interval1, "DownloadStoryForeverViaCleanAccount: ")
	}
}

// DO NOT USE. Due to Instagram changes the rate limit of private API, use of
// this method will cause HTTP 429. Will be removed soon.
func (m *IGDownloadManager) DownloadUnexpiredStoryOfUser(user instago.User) (err error) {
	// In case user change privacy, read user info via mobile api endpoint
	// first.
	u, err := m.SmartGetUserInfoEndPoint(user.GetUserId())
	if err == nil {
		err = m.SmartDownloadUserStoryPostliveLayer(u, 2)
	} else {
		log.Println(err)
		log.Println("Fail to fetch user info via mobile endpoint. use old user info data to fetch")
		err = m.SmartDownloadUserStoryPostliveLayer(user, 2)
	}
	return
}

// DO NOT USE. Due to Instagram changes the rate limit of private API, use of
// this method will cause HTTP 429. Will be removed soon.
func (m *IGDownloadManager) DownloadUnexpiredStoryOfFollowUsers(users []instago.IGFollowUser, interval int) (err error) {
	for _, user := range users {
		err = m.DownloadUnexpiredStoryOfUser(user)
		if err != nil {
			log.Println(err)
		}
		SleepLog(interval)
	}
	return
}

// DO NOT USE. Due to Instagram changes the rate limit of private API, use of
// this method will cause HTTP 429. Will be removed soon.
func (m *IGDownloadManager) DownloadUnexpiredStoryOfAllFollowingUsers(interval int) (err error) {
	log.Println("Load following users from data dir first")
	users, err := LoadLatestFollowingUsers()
	if err != nil {
		log.Println(err)
		log.Println("Fail to load users from data dir. Try to load following users from Instagram")
		users, err = m.GetSelfFollowing()
		if err != nil {
			log.Println(err)
			return
		}
	}

	return m.DownloadUnexpiredStoryOfFollowUsers(users, interval)
}
