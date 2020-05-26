package igdl

import (
	"log"
	"os"

	"github.com/siongui/instago"
)

func (m *IGDownloadManager) SmartDownloadAllPosts(username string) (err error) {
	err = DownloadAllPostsNoLogin(username)
	if err == nil {
		return
	}

	if m.mgr2 != nil {
		err = m.mgr2.DownloadAllPosts(username)
		//err = m.mgr2.DownloadAllPostsNoLogin(username)
	}
	return
}

func (m *IGDownloadManager) SmartDownloadStory(user instago.IGUser) (err error) {
	// Pk here is user id
	if user.IsPrivate {
		return m.DownloadUserStoryPostlive(user.Pk)
	}

	// in case main account is blocked by some users, we use clean account
	// (account not blocked) to download public user account
	if m.mgr2 != nil {
		return m.mgr2.DownloadUserStoryPostlive(user.Pk)
	}

	return m.DownloadUserStoryPostlive(user.Pk)
}

func (m *IGDownloadManager) SmartDownloadPost(item instago.IGItem) (err error) {
	if item.User.IsPrivate {
		_, err = m.DownloadPost(item.GetPostCode())
		return
	}

	_, err = DownloadPostNoLogin(item.GetPostCode())
	if err == nil {
		return
	}

	if m.mgr2 != nil {
		_, err = m.mgr2.DownloadPost(item.GetPostCode())
	}
	return
}

// GetPostItem downloads media (photo/video) item in the post.
func (m *IGDownloadManager) GetPostItem(item instago.IGItem) (isDownloaded bool, err error) {
	if saveData {
		saveIdUsername(item.GetUserId(), item.GetUsername())
	}

	// FIXME: Use item.IsRegularMedia() to check validity of item?

	urls, err := item.GetMediaUrls()
	if err != nil {
		log.Println(err)
		return
	}

	for index, url := range urls {
		taggedusers := [][2]string{}
		if len(urls) == 1 {
			taggedusers = item.Usertags.GetIdUsernamePairs()
		} else {
			taggedusers = item.CarouselMedia[index].Usertags.GetIdUsernamePairs()
		}

		if saveData {
			saveTaggedUsers(taggedusers)
		}

		filepath := getPostFilePath2(
			item.GetUsername(),
			item.GetUserId(),
			item.GetPostCode(),
			url,
			item.GetTimestamp(),
			taggedusers)
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
				m.SmartDownloadPost(item)
			}
		} else {
			if err != nil {
				log.Println(err)
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
		PrintItemInfo(idx, &item)
		isDownloaded, _ := m.GetPostItem(item)
		if isDownloaded && downloadStory {
			u := item.GetUsername()
			if _, ok := username[u]; !ok {
				go m.SmartDownloadStory(item.User)
				username[u] = true
			}
		}
	}
}

// DO NOT USE. Test now. Used with DownloadDependOnCollectionName
func (m *IGDownloadManager) DownloadSavedPostsAndSendItemInCollectionToChannel(numOfItem int, downloadStory bool, c chan instago.IGItem) (err error) {
	items, err := m.apimgr.GetSavedPosts(numOfItem)
	if err != nil {
		log.Println(err)
		return
	}

	username := make(map[string]bool)
	for idx, item := range items {
		// FIXME: check err
		PrintItemInfo(idx, &item)
		isDownloaded, _ := m.GetPostItem(item)
		if isDownloaded && downloadStory {
			u := item.GetUsername()
			if _, ok := username[u]; !ok {
				go m.SmartDownloadStory(item.User)
				username[u] = true
			}
		}

		// if item in collection, send to channel
		if len(item.SavedCollectionIds) > 0 {
			c <- item
		}
	}

	return
}

// DO NOT USE. Test now. Used with DownloadDependOnCollectionName
func (m *IGDownloadManager) DownloadSavedCollectionPostsAndSendItemInCollectionToChannel(collectionName string, downloadStory bool, c chan instago.IGItem) (err error) {
	cid := m.CollectionName2Id(collectionName)
	if cid == "" {
		log.Println("fail to get collection id of ", collectionName)
	}

	items, err := m.apimgr.GetSavedCollection(cid)
	if err != nil {
		log.Println(err)
		return
	}

	username := make(map[string]bool)
	for idx, item := range items {
		// FIXME: check err
		PrintItemInfo(idx, &item)
		isDownloaded, _ := m.GetPostItem(item)
		if isDownloaded && downloadStory {
			u := item.GetUsername()
			if _, ok := username[u]; !ok {
				go m.SmartDownloadStory(item.User)
				username[u] = true
			}
		}

		// if item in collection, send to channel
		if len(item.SavedCollectionIds) > 0 {
			c <- item
		}
	}

	return
}

// DO NOT USE. Test now. Used with DownloadSavedPostsAndSendItemInCollectionToChannel
func (m *IGDownloadManager) DownloadDependOnCollectionName(name2layer, nameAllpost, nameHighlight string, c chan instago.IGItem) {
	map2layer := make(map[string]bool)
	mapAllpost := make(map[string]bool)
	mapHighlight := make(map[string]bool)
	for {
		item := <-c
		//log.Println(item.GetUsername())
		for _, cid := range item.SavedCollectionIds {

			name := m.CollectionId2Name(cid)
			//log.Println(name)

			iddone := cid + "-" + item.GetUsername() + "-" + item.GetUserId() + "-" + item.GetPostCode()
			//log.Println(iddone)

			if name == name2layer {
				if _, isDone := map2layer[iddone]; !isDone {
					log.Println(item.GetUsername(), "download 2layer", iddone)
					// TODO: smart download 2layer
					go m.DownloadUserStoryLayer(item.User.Pk, 2)
					map2layer[iddone] = true
				}
			}

			if name == nameAllpost {
				if _, isDone := mapAllpost[iddone]; !isDone {
					log.Println(item.GetUsername(), "download all post (no login)", iddone)
					go m.SmartDownloadAllPosts(item.GetUsername())
					mapAllpost[iddone] = true
				}
			}

			if name == nameHighlight {
				if _, isDone := mapHighlight[iddone]; !isDone {
					log.Println(item.GetUsername(), "download highlight", iddone)
					// TODO: smart download highlights
					go m.DownloadUserStoryHighlights(item.GetUserId())
					mapHighlight[iddone] = true
				}
			}
		}
	}
}

func (m *IGDownloadManager) GetSavedPosts(numOfItem int) (items []instago.IGItem, err error) {
	return m.apimgr.GetSavedPosts(numOfItem)
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

func (m *IGDownloadManager) CollectionId2Name(id string) string {
	for _, collection := range m.collections {
		if collection.CollectionId == id {
			return collection.CollectionName
		}
	}
	return ""
}

func (m *IGDownloadManager) CollectionName2Id(name string) string {
	for _, collection := range m.collections {
		if collection.CollectionName == name {
			return collection.CollectionId
		}
	}
	return ""
}
