package igdl

import (
	"log"
)

// DownloadSavedPost downloads all your saved posts.
func (m *IGDownloadManager) DownloadSavedPosts() {
	items, err := m.apimgr.GetSavedPosts()
	if err != nil {
		log.Println(err)
		return
	}
	getTimelineItems(items)
}
