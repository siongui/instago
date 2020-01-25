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

	// The following code is not good because
	// 1. Cannot download saved posts of non-following users
	// 2. not best quality of saved posts
	//getTimelineItems(items)

	for _, item := range items {
		m.DownloadPost(item.Code)
	}
}
