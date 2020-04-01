package igdl

import (
	"log"
	"os"

	"github.com/siongui/instago"
)

// getPostItem downloads media (photo/video) item in the post.
// TODO: try to merge getPostItem and DownloadPostItem
func (m *IGDownloadManager) getPostItem(item instago.IGItem) (isDownloaded bool, err error) {
	// FIXME: Use item.IsRegularMedia() to check validity of item?

	urls, err := item.GetMediaUrls()
	if err != nil {
		log.Println(err)
		return
	}

	for index, url := range urls {
		filepath := getPostFilePath(
			item.GetUsername(),
			item.GetUserId(),
			item.GetPostCode(),
			url,
			item.GetTimestamp())
		if index > 0 {
			filepath = appendIndexToFilename(filepath, index)
		}

		CreateFilepathDirIfNotExist(filepath)
		// check if file exist
		if _, err := os.Stat(filepath); os.IsNotExist(err) {
			// file not exists
			if false {
				// never run here. Old way of implementation
				// only for reference. Old way is not good
				// because not best quality of posts
				printDownloadInfo(&item, url, filepath)
				err = Wget(url, filepath)
				if err != nil {
					log.Println(err)
				} else {
					isDownloaded = true
				}
			} else {
				// always run here.
				return m.DownloadPost(item.GetPostCode())
			}
		}
	}
	return
}

// DownloadSavedPost downloads your saved posts.
// -1 means download all saved posts.
// downloadStory flag will also download unexpired stories of the post user.
func (m *IGDownloadManager) DownloadSavedPosts(numOfItem int, downloadStory bool) {
	items, err := m.apimgr.GetSavedPosts(numOfItem)
	if err != nil {
		log.Println(err)
		return
	}

	username := make(map[string]bool)
	for idx, item := range items {
		// FIXME: check err
		printItemInfo(idx, &item)
		isDownloaded, _ := m.getPostItem(item)
		if isDownloaded && downloadStory {
			u := item.GetUsername()
			if _, ok := username[u]; !ok {
				// Pk here is user id
				m.DownloadUserStoryPostlive(item.User.Pk)
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
