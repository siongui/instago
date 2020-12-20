package igdl

import (
	"log"
	"strconv"

	"github.com/siongui/instago"
)

func (m *IGDownloadManager) downloadUserStoryPostlive(id string) (err error) {
	ut, err := m.GetUserStory(id)
	if err != nil {
		log.Println(err)
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
		log.Println(err)
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
	return m.downloadUserStoryPostliveLayer(user, layer, isdone)
}

func (m *IGDownloadManager) DownloadUserStoryPostliveByNameLayerIfPublic(username string, layer int) (err error) {
	user, err := m.UsernameToUser(username)
	if err != nil {
		return
	}

	if user.IsPublic() {
		isdone := make(map[string]string)
		// FIXME: if reel_mention is private, do not download
		return m.downloadUserStoryPostliveLayer(user, layer, isdone)
	}
	PrintUserInfo(user)
	return
}
