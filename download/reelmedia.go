package igdl

import (
	"log"

	"github.com/siongui/instago"
)

func (m *IGDownloadManager) downloadUserReelMedia(id string) (err error) {
	tray, err := m.GetUserReelMedia(id)
	if err != nil {
		log.Println(err)
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

// DownloadUserReelMediaByName downloads unexpired stories (last 24 hours) of
// the given user name. No postlives included.
func (m *IGDownloadManager) DownloadUserReelMediaByName(username string) (err error) {
	id, err := m.UsernameToId(username)
	if err != nil {
		log.Println(err)
		return
	}

	return m.downloadUserReelMedia(id)
}

// DownloadUserReelMedia downloads unexpired stories (last 24 hours) of the
// given user id. No postlives included.
func (m *IGDownloadManager) DownloadUserReelMedia(id string) (err error) {
	return m.downloadUserReelMedia(id)
}

func (m *IGDownloadManager) getReelMediaItemLayer(item instago.IGItem, username string, layer int64, isdone map[string]string, tl *TimeLimiter) {
	getStoryItem(item, username)
	for _, reelmention := range item.ReelMentions {
		PrintReelMentionInfo(reelmention)
		m.downloadUserReelMediaLayer(reelmention.GetUserId(), layer, isdone, tl)
	}
}

func (m *IGDownloadManager) downloadUserReelMediaLayer(id string, layer int64, isdone map[string]string, tl *TimeLimiter) (err error) {
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

	tl.WaitAtLeastIntervalAfterLastTime()
	tray, err := m.GetUserReelMedia(id)
	tl.SetLastTimeToNow()
	if err != nil {
		return
	}

	isdone[id] = tray.GetUsername()
	log.Println("fetch story of", tray.GetUsername(), id, "success")

	for _, item := range tray.GetItems() {
		m.getReelMediaItemLayer(item, tray.GetUsername(), layer, isdone, tl)
	}
	return
}

// DownloadUserReelMediaLayer downloads reel mentions in the stories as well.
// Parameter layer 1 is actually the same as DownloadUserReelMedia.
func (m *IGDownloadManager) DownloadUserReelMediaLayer(id string, layer, interval int64) (err error) {
	isdone := make(map[string]string)
	tl := NewTimeLimiter(interval)
	return m.downloadUserReelMediaLayer(id, layer, isdone, tl)
}
