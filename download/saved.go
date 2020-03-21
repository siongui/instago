package igdl

import (
	"log"

	"github.com/siongui/instago"
)

// DownloadSavedPost downloads your saved posts.
// -1 means download all saved posts.
// downloadStory flag will also download unexpired stories of the post user.
func (m *IGDownloadManager) DownloadSavedPosts(numOfItem int, downloadStory bool) {
	items, err := m.apimgr.GetSavedPosts(numOfItem)
	if err != nil {
		log.Println(err)
		return
	}

	// The following code is not good because
	// 1. Cannot download saved posts of non-following users
	// 2. not best quality of saved posts
	//getTimelineItems(items)

	username := make(map[string]bool)
	for _, item := range items {
		isDownloaded := m.DownloadPost(item.Code)
		if isDownloaded && downloadStory {
			u := item.GetUsername()
			if _, ok := username[u]; !ok {
				// Pk here is user id
				m.DownloadUserStory(item.User.Pk)
				username[u] = true
			}
		}
	}
}

func (m *IGDownloadManager) IsInCollection(item instago.IGItem, name string) bool {
	for _, id := range item.SavedCollectionIds {
		for _, collection := range m.collections {
			if collection.CollectionId == id && name == collection.CollectionName {
				return true
			}
		}
	}
	return false
}
