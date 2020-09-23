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

	filepath := GetStoryFilePath(
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
	tray, err := m.apimgr.GetUserReelMedia(id)
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
	ut, err := m.GetUserStory(id)
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
		//m.downloadUserStoryLayer(reelmention, layer, isdone)
		m.downloadUserStoryPostliveLayer(reelmention, layer, isdone)
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

	tray, err := m.apimgr.GetUserReelMedia(id)
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

func (m *IGDownloadManager) downloadUserStoryPostliveLayer(user instago.User, layer int, isdone map[string]string) (err error) {
	if layer < 1 {
		return
	}
	layer--

	PrintUserInfo(user)
	id := user.GetUserId()
	if username, ok := isdone[id]; ok {
		log.Println(username, id, "already fetched")
		return
	} else {
		log.Println("fetching story of", user.GetUsername(), id)
	}

	ut, err := m.GetUserStory(id)
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
	user, err := m.UsernameToUser(username)
	if err != nil {
		return
	}

	isdone := make(map[string]string)
	//return m.downloadUserStoryLayer(user, layer, isdone)
	return m.downloadUserStoryPostliveLayer(user, layer, isdone)
}

// DownloadUserStoryLayer is the same as DownloadUserStoryByNameLayer, except
// int64 id passed as argument.
func (m *IGDownloadManager) DownloadUserStoryLayer(userId int64, layer int) (err error) {
	user, err := m.GetUserInfoEndPoint(strconv.FormatInt(userId, 10))
	if err != nil {
		return
	}

	isdone := make(map[string]string)
	//return m.downloadUserStoryLayer(user, layer, isdone)
	return m.downloadUserStoryPostliveLayer(user, layer, isdone)
}

func (m *IGDownloadManager) DownloadUserStoryPostliveByNameLayerIfPublic(username string, layer int) (err error) {
	user, err := m.UsernameToUser(username)
	if err != nil {
		return
	}

	if user.IsPublic() {
		isdone := make(map[string]string)
		//return m.downloadUserStoryLayer(user, layer, isdone)
		// FIXME: if reel_mention is private, do not download
		return m.downloadUserStoryPostliveLayer(user, layer, isdone)
	}
	PrintUserInfo(user)
	return
}
