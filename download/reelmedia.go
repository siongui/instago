package igdl

import (
	"log"
	"time"

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

func (m *IGDownloadManager) getReelMediaItemLayer(item instago.IGItem, username string, layer, interval int64, isdone map[string]string, getTime map[string]time.Time) {
	getStoryItem(item, username)
	for _, reelmention := range item.ReelMentions {
		PrintReelMentionInfo(reelmention)
		m.downloadUserReelMediaLayer(reelmention.GetUserId(), layer, interval, isdone, getTime)
	}
}

func (m *IGDownloadManager) downloadUserReelMediaLayer(id string, layer, interval int64, isdone map[string]string, getTime map[string]time.Time) (err error) {
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

	d := time.Now().Sub(getTime["last"])
	if d < time.Duration(interval)*time.Second {
		time.Sleep(time.Duration(interval)*time.Second - d)
	}
	tray, err := m.GetUserReelMedia(id)
	getTime["last"] = time.Now()
	if err != nil {
		return
	}

	isdone[id] = tray.GetUsername()
	log.Println("fetch story of", tray.GetUsername(), id, "success")

	for _, item := range tray.GetItems() {
		m.getReelMediaItemLayer(item, tray.GetUsername(), layer, interval, isdone, getTime)
	}
	return
}

func (m *IGDownloadManager) DownloadUserReelMediaLayer(id string, layer, interval int64) (err error) {
	isdone := make(map[string]string)
	t := time.Unix(time.Now().Unix()-interval-1, 0)
	getTime := make(map[string]time.Time)
	getTime["last"] = t
	return m.downloadUserReelMediaLayer(id, layer, interval, isdone, getTime)
}
