package igdl

import (
	"fmt"
	"os"
	"strconv"

	"github.com/siongui/instago"
)

func getStoryItem(item instago.IGItem, username string) {
	if !(item.MediaType == 1 || item.MediaType == 2) {
		fmt.Println("In getStoryItem: not single photo or video!")
		return
	}

	urls, err := item.GetMediaUrls()
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(urls) != 1 {
		fmt.Println("In getStoryItem: number of download url != 1")
		return
	}
	url := urls[0]

	filepath := getStoryFilePath(
		username,
		item.GetUserId(),
		item.GetPostCode(),
		url,
		item.GetTimestamp())

	CreateFilepathDirIfNotExist(filepath)
	// check if file exist
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		// file not exists
		printDownloadInfo(item, username, url, filepath)
		err = Wget(url, filepath)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func (m *IGDownloadManager) UsernameToId(username string) (id string, err error) {
	// Try to get id without loggin
	id, err = instago.GetUserId(username)
	if err == nil {
		return
	}

	// Try to get id with loggin
	ui, err := m.apimgr.GetUserInfo(username)
	if err == nil {
		id = ui.Id
	}
	return
}

// DownloadUserStoryByName downloads unexpired stories (last 24 hours) of the
// given user name.
func (m *IGDownloadManager) DownloadUserStoryByName(username string) {
	id, err := m.UsernameToId(username)
	if err != nil {
		panic(err)
	}

	tray, err := m.apimgr.GetUserStory(id)
	if err != nil {
		panic(err)
	}
	for _, item := range tray.GetItems() {
		getStoryItem(item, tray.GetUsername())
	}
	return
}

// DownloadUserStory downloads unexpired stories (last 24 hours) of the given
// user id.
func (m *IGDownloadManager) DownloadUserStory(userId int64) (err error) {
	tray, err := m.apimgr.GetUserStory(strconv.FormatInt(userId, 10))
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, item := range tray.GetItems() {
		getStoryItem(item, tray.GetUsername())
	}
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

func (m *IGDownloadManager) getStoryItemLayer(item instago.IGItem, username string, layer int) {
	getStoryItem(item, username)
	for _, reelmention := range item.ReelMentions {
		m.DownloadUserStoryLayer(reelmention.User.Pk, layer)
	}
}

// DownloadUserStoryByNameLayer downloads unexpired stories (last 24 hours) of
// the given user name, and also stories of reel mentions.
func (m *IGDownloadManager) DownloadUserStoryByNameLayer(username string, layer int) {
	if layer < 1 {
		return
	}
	layer--

	id, err := m.UsernameToId(username)
	if err != nil {
		panic(err)
	}

	tray, err := m.apimgr.GetUserStory(id)
	if err != nil {
		panic(err)
	}
	for _, item := range tray.GetItems() {
		m.getStoryItemLayer(item, tray.GetUsername(), layer)
	}
	return
}

// DownloadUserStoryLayer is the same as DownloadUserStoryByNameLayer, except
// int64 id passed as argument.
func (m *IGDownloadManager) DownloadUserStoryLayer(userId int64, layer int) {
	if layer < 1 {
		return
	}
	layer--

	tray, err := m.apimgr.GetUserStory(strconv.FormatInt(userId, 10))
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, item := range tray.GetItems() {
		m.getStoryItemLayer(item, tray.GetUsername(), layer)
	}
	return
}
