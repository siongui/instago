package igdl

import (
	"log"
	"strconv"

	"github.com/siongui/instago"
)

func (m *IGDownloadManager) downloadUserStory(id string) (err error) {
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

	// FIXME: it seems PostLiveItem does not exist anymore? remove it?
	return DownloadPostLiveItem(ut.PostLiveItem)
}

func (m *IGDownloadManager) DownloadUserStory(id string) (err error) {
	return m.downloadUserStory(id)
}

func (m *IGDownloadManager) DownloadUserStoryByName(username string) (err error) {
	id, err := m.UsernameToId(username)
	if err != nil {
		log.Println(err)
		return
	}

	return m.downloadUserStory(id)
}

func (m *IGDownloadManager) getStoryItemLayer(item instago.IGItem, username string, layer int, isdone map[string]string, tl *TimeLimiter) {
	getStoryItem(item, username)
	for _, reelmention := range item.ReelMentions {
		m.downloadUserStoryLayer(reelmention, layer, isdone, tl)
	}
}

func (m *IGDownloadManager) downloadUserStoryLayer(user instago.User, layer int, isdone map[string]string, tl *TimeLimiter) (err error) {
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

	tl.WaitAtLeastIntervalAfterLastTime()
	ut, err := m.GetUserStory(id)
	tl.SetLastTimeToNow()
	if err != nil {
		return
	}
	tray := ut.Reel

	isdone[id] = tray.GetUsername()
	log.Println("fetch story of", tray.GetUsername(), id, "success")

	for _, item := range tray.GetItems() {
		m.getStoryItemLayer(item, tray.GetUsername(), layer, isdone, tl)
	}

	return DownloadPostLiveItem(ut.PostLiveItem)
}

func (m *IGDownloadManager) DownloadUserStoryByNameLayer(username string, layer int) (err error) {
	user, err := m.UsernameToUser(username)
	if err != nil {
		return
	}

	isdone := make(map[string]string)
	tl := NewTimeLimiter(12)
	return m.downloadUserStoryLayer(user, layer, isdone, tl)
}

func (m *IGDownloadManager) DownloadUserStoryLayer(userId int64, layer int) (err error) {
	user, err := m.GetUserInfoEndPoint(strconv.FormatInt(userId, 10))
	if err != nil {
		return
	}

	isdone := make(map[string]string)
	tl := NewTimeLimiter(12)
	return m.downloadUserStoryLayer(user, layer, isdone, tl)
}
