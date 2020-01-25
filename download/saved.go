package igdl

import (
	"log"
)

// DownloadSavedPost downloads all your saved posts.
func (m *IGDownloadManager) DownloadSavedPosts(numOfItem int) {
	items, err := m.apimgr.GetSavedPosts(numOfItem)
	if err != nil {
		log.Println(err)
		return
	}
	getTimelineItems(items)
}
