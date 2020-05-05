package igdl

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/siongui/instago"
)

func getStoryItem(item instago.IGItem, username string) (err error) {
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
		if err != nil {
			fmt.Println(err)
			return err
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
		getStoryItem(item, tray.GetUsername())
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

// DownloadUserStoryPostLive downloads unexpired stories (last 24 hours) and
// postlive of the given user id.
func (m *IGDownloadManager) DownloadUserStoryPostlive(userId int64) (err error) {
	ut, err := m.apimgr.GetUserReelMedia(strconv.FormatInt(userId, 10))
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, item := range ut.Reel.GetItems() {
		getStoryItem(item, ut.Reel.GetUsername())
	}
	DownloadPostLiveItem(ut.PostLiveItem)

	return
}

// DownloadUserStoryPostLiveByName is the same as DownloadUserStoryPostlive,
// except username is given as argument.
func (m *IGDownloadManager) DownloadUserStoryPostliveByName(username string) (err error) {
	id, err := m.UsernameToId(username)
	if err != nil {
		fmt.Println(err)
		return
	}

	ut, err := m.apimgr.GetUserReelMedia(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, item := range ut.Reel.GetItems() {
		getStoryItem(item, ut.Reel.GetUsername())
	}
	DownloadPostLiveItem(ut.PostLiveItem)

	return
}

// DownloadUnreadStory downloads all available stories in IGReelTray.
func DownloadUnreadStory(trays []instago.IGReelTray) {
	for _, tray := range trays {
		//fmt.Println(tray.GetUsername())
		for _, item := range tray.GetItems() {
			getStoryItem(item, tray.GetUsername())
		}
	}
}

func (m *IGDownloadManager) fetchUserStory(userId int64, username string, c chan int) {
	defer func() { c <- 1 }()

	err := m.DownloadUserStory(userId)
	if err != nil {
		fmt.Println("In fetchUserStorie: fail to fetch " + username)
		fmt.Println(err)
	}
}

// DownloadAllStory downloads all unexpired stories of all users in IGReelTray.
func (m *IGDownloadManager) DownloadAllStory(trays []instago.IGReelTray) {
	c := make(chan int)
	numOfStoryUser := 0
	for _, tray := range trays {
		items := tray.GetItems()
		if len(items) == 0 {
			numOfStoryUser++
			go m.fetchUserStory(tray.Id, tray.GetUsername(), c)
		} else {
			for _, item := range items {
				getStoryItem(item, tray.GetUsername())
			}
		}
	}

	// wait all goroutines to finish
	for i := 0; i < numOfStoryUser; i++ {
		<-c
	}
}

func (m *IGDownloadManager) getStoryItemLayer(item instago.IGItem, username string, layer int, isdone map[string]string) {
	getStoryItem(item, username)
	for _, reelmention := range item.ReelMentions {
		// Pk is user id
		id := strconv.FormatInt(reelmention.User.Pk, 10)
		m.downloadUserStoryLayer(id, layer, isdone)
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

// DownloadUserStoryByNameLayer downloads unexpired stories (last 24 hours) of
// the given user name, and also stories of reel mentions.
func (m *IGDownloadManager) DownloadUserStoryByNameLayer(username string, layer int) (err error) {
	id, err := m.UsernameToId(username)
	if err != nil {
		return
	}

	isdone := make(map[string]string)
	return m.downloadUserStoryLayer(id, layer, isdone)
}

// DownloadUserStoryLayer is the same as DownloadUserStoryByNameLayer, except
// int64 id passed as argument.
func (m *IGDownloadManager) DownloadUserStoryLayer(userId int64, layer int) (err error) {
	isdone := make(map[string]string)
	return m.downloadUserStoryLayer(strconv.FormatInt(userId, 10), layer, isdone)
}
