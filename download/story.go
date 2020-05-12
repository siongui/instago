package igdl

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/siongui/instago"
)

func GetStoryItem(item instago.IGItem, username string) (isDownloaded bool, err error) {
	return getStoryItem(item, username)
}

func getStoryItem(item instago.IGItem, username string) (isDownloaded bool, err error) {
	if !(item.MediaType == 1 || item.MediaType == 2) {
		err = errors.New("In getStoryItem: not single photo or video!")
		fmt.Println(err)
		return
	}

	urls, err := item.GetMediaUrls()
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(urls) != 1 {
		err = errors.New("In getStoryItem: number of download url != 1")
		fmt.Println(err)
		return
	}
	url := urls[0]

	if saveData {
		saveIdUsername(item.GetUserId(), username)
		saveReelMentions(item.ReelMentions)
	}

	// fix missing username issue while print download info
	item.User.Username = username

	filepath := getStoryFilePath2(
		username,
		item.GetUserId(),
		item.GetPostCode(),
		url,
		item.GetTimestamp(),
		item.ReelMentions)

	CreateFilepathDirIfNotExist(filepath)
	// check if file exist
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		// file not exists
		printDownloadInfo(&item, url, filepath)
		err = Wget(url, filepath)
		if err == nil {
			isDownloaded = true
		} else {
			log.Println(err)
			return isDownloaded, err
		}
	} else {
		if err != nil {
			log.Println(err)
		}
	}
	return
}

func (m *IGDownloadManager) downloadUserStory(id string) (err error) {
	tray, err := m.apimgr.GetUserStory(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, item := range tray.GetItems() {
		_, err := getStoryItem(item, tray.GetUsername())
		if err != nil {
			log.Println(err)
			//return
		}
	}
	return
}

// DownloadUserStoryByName downloads unexpired stories (last 24 hours) of the
// given user name.
func (m *IGDownloadManager) DownloadUserStoryByName(username string) (err error) {
	id, err := m.UsernameToId(username)
	if err != nil {
		fmt.Println(err)
		return
	}

	return m.downloadUserStory(id)
}

// DownloadUserStory downloads unexpired stories (last 24 hours) of the given
// user id.
func (m *IGDownloadManager) DownloadUserStory(userId int64) (err error) {
	return m.downloadUserStory(strconv.FormatInt(userId, 10))
}

func (m *IGDownloadManager) downloadUserStoryPostlive(id string) (err error) {
	ut, err := m.apimgr.GetUserReelMedia(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, item := range ut.Reel.GetItems() {
		_, err = getStoryItem(item, ut.Reel.GetUsername())
		if err != nil {
			log.Println(err)
			//return
		}
	}
	return DownloadPostLiveItem(ut.PostLiveItem)
}

// DownloadUserStoryPostLive downloads unexpired stories (last 24 hours) and
// postlive of the given user id.
func (m *IGDownloadManager) DownloadUserStoryPostlive(userId int64) (err error) {
	return m.downloadUserStoryPostlive(strconv.FormatInt(userId, 10))
}

// DownloadUserStoryPostLiveByName is the same as DownloadUserStoryPostlive,
// except username is given as argument.
func (m *IGDownloadManager) DownloadUserStoryPostliveByName(username string) (err error) {
	id, err := m.UsernameToId(username)
	if err != nil {
		fmt.Println(err)
		return
	}

	return m.downloadUserStoryPostlive(id)
}

func (m *IGDownloadManager) getStoryItemLayer(item instago.IGItem, username string, layer int, isdone map[string]string) {
	getStoryItem(item, username)
	for _, reelmention := range item.ReelMentions {
		// Pk is user id
		id := strconv.FormatInt(reelmention.User.Pk, 10)
		//m.downloadUserStoryLayer(id, layer, isdone)
		m.downloadUserStoryPostliveLayer(id, layer, isdone)
	}
}

func (m *IGDownloadManager) downloadUserStoryLayer(id string, layer int, isdone map[string]string) (err error) {
	if layer < 1 {
		return
	}
	layer--

	if username, ok := isdone[id]; ok {
		log.Println(username, id, "already fetched")
		return
	} else {
		log.Println("fetching story of", id)
	}

	tray, err := m.apimgr.GetUserStory(id)
	if err != nil {
		return
	}
	isdone[id] = tray.GetUsername()
	log.Println("fetch story of", tray.GetUsername(), id, "success")

	for _, item := range tray.GetItems() {
		m.getStoryItemLayer(item, tray.GetUsername(), layer, isdone)
	}
	return
}

func (m *IGDownloadManager) downloadUserStoryPostliveLayer(id string, layer int, isdone map[string]string) (err error) {
	if layer < 1 {
		return
	}
	layer--

	if username, ok := isdone[id]; ok {
		log.Println(username, id, "already fetched")
		return
	} else {
		log.Println("fetching story of", id)
	}

	ut, err := m.apimgr.GetUserReelMedia(id)
	if err != nil {
		return
	}
	tray := ut.Reel

	isdone[id] = tray.GetUsername()
	log.Println("fetch story of", tray.GetUsername(), id, "success")

	for _, item := range tray.GetItems() {
		m.getStoryItemLayer(item, tray.GetUsername(), layer, isdone)
	}

	return DownloadPostLiveItem(ut.PostLiveItem)
}

// DownloadUserStoryByNameLayer downloads unexpired stories (last 24 hours) of
// the given user name, and also stories of reel mentions.
func (m *IGDownloadManager) DownloadUserStoryByNameLayer(username string, layer int) (err error) {
	id, err := m.UsernameToId(username)
	if err != nil {
		return
	}

	isdone := make(map[string]string)
	//return m.downloadUserStoryLayer(id, layer, isdone)
	return m.downloadUserStoryPostliveLayer(id, layer, isdone)
}

// DownloadUserStoryLayer is the same as DownloadUserStoryByNameLayer, except
// int64 id passed as argument.
func (m *IGDownloadManager) DownloadUserStoryLayer(userId int64, layer int) (err error) {
	isdone := make(map[string]string)
	//return m.downloadUserStoryLayer(strconv.FormatInt(userId, 10), layer, isdone)
	return m.downloadUserStoryPostliveLayer(strconv.FormatInt(userId, 10), layer, isdone)
}

func isTrayInQueue(queue []instago.IGReelTray, tray instago.IGReelTray) bool {
	for _, t := range queue {
		if t.Id == tray.Id {
			return true
		}
	}
	return false
}

func (m *IGDownloadManager) GetStoryItemAndReelMentions(item instago.IGItem, username string) (err error) {
	isDownloaded, err := getStoryItem(item, username)
	if err != nil {
		return
	}

	fetched := make(map[string]bool)
	if isDownloaded {
		for _, rm := range item.ReelMentions {
			PrintReelMentionInfo(rm)
			if rm.GetUsername() == username {
				fmt.Println("reel mention self. download ignored.")
				continue
			}

			if !rm.User.IsPrivate {
				if _, ok := fetched[rm.GetUsername()]; !ok {
					m.downloadUserStoryPostlive(rm.GetUserId())
					fetched[rm.GetUsername()] = true
				}
				// handle err of m.downloadUserStoryPostlive(rm.GetUserId()) ?
				/*
					err = m.downloadUserStoryPostlive(rm.GetUserId())
					if err != nil {
						PrintUsernameIdMsg(rm.GetUsername(), rm.GetUserId(), err)
						return
					}
				*/
			}
		}
	}
	return
}

func (m *IGDownloadManager) DownloadZeroItemUsers(c chan instago.IGReelTray, interval int, verbose bool) {
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
					ut, err := m.apimgr.GetUserReelMedia(id)
					if err != nil {
						PrintUsernameIdMsg(username, id, err)
						queue = append(queue, tray)
						return
					}

					for _, item := range ut.Reel.GetItems() {
						err = m.GetStoryItemAndReelMentions(item, ut.Reel.GetUsername())
						if err != nil {
							PrintUsernameIdMsg(username, id, err)
							queue = append(queue, tray)
							return
						}
					}

					err = DownloadPostLiveItem(ut.PostLiveItem)
					if err == nil {
						if verbose {
							PrintUsernameIdMsg(username, id, " Download Success.")
						}
					} else {
						PrintUsernameIdMsg(username, id, err)
						queue = append(queue, tray)
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

func isLatestReelMediaDownloaded(username string, latestReelMedia int64) bool {
	utimes, err := GetReelMediaUnixTimesInUserStoryDir(username)
	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Println("In isLatestReelMediaDownloaded", err)
		}
		return false
	}

	lrm := strconv.FormatInt(latestReelMedia, 10)
	for _, utime := range utimes {
		if lrm == utime {
			return true
		}
	}
	return false
}

// Use (25, 2, false) as arguments is good. will not cause http 429
func (m *IGDownloadManager) DownloadStoryAndPostLiveForever(interval1, interval2 int, verbose bool) {
	// channel for waiting DownloadPostLive completed
	isDownloading := make(map[string]bool)

	// usually there are at most 150 trays in reels_tray.
	// double the buffer to 300. 160 or 200 may be ok as well.
	c := make(chan instago.IGReelTray, 300)

	go m.DownloadZeroItemUsers(c, interval2, verbose)
	for {
		rt, err := m.GetReelsTray()
		if err != nil {
			log.Println(err)
			continue
		}

		// TODO: use channel for postlive?
		go DownloadPostLive(rt.PostLive, isDownloading)
		go PrintLiveBroadcasts(rt.Broadcasts)

		for index, tray := range rt.Trays {
			username := tray.GetUsername()
			id := tray.Id
			items := tray.GetItems()

			if isLatestReelMediaDownloaded(username, tray.LatestReelMedia) {
				if verbose {
					PrintUsernameIdMsg(username, id, " all downloaded")
				}
				continue
			}

			if verbose {
				UsernameIdColorPrint(username, id)
				fmt.Println(" has undownloaded items")
			}

			if len(items) == 0 {
				if verbose {
					UsernameIdColorPrint(username, id)
					fmt.Println("#", index, " send to channel")
				}
				c <- tray
			} else {
				for _, item := range items {
					err = m.GetStoryItemAndReelMentions(item, username)
					if err != nil {
						PrintUsernameIdMsg(username, id, err)
						c <- tray
						break
					}
				}
				// is there postlive items in tray here?
			}
		}

		PrintMsgSleep(interval1, "DownloadStoryAndPostLiveForever: ")
	}
}
