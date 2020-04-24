package igdl

import (
	"log"
	"os"

	"github.com/siongui/instago"
)

func DownloadPostNoLogin(code string) (isDownloaded bool, err error) {
	em, err := instago.GetPostInfoNoLogin(code)
	if err != nil {
		log.Println(err)
		return
	}

	return DownloadIGMedia(em)
}

func (m *IGDownloadManager) DownloadPost(code string) (isDownloaded bool, err error) {
	em, err := m.apimgr.GetPostInfo(code)
	if err != nil {
		log.Println(err)
		return
	}

	return DownloadIGMedia(em)
}

// DownloadIGMedia downloads photos/videos in the post.
// IGItem (items in timeline or saved posts) or IGMedia (read from
// https://www.instagram.com/p/{{CODE}}/?__a=1) can be argument in this method.
func DownloadIGMedia(em instago.IGMedia) (isDownloaded bool, err error) {
	if saveData {
		saveIdUsername(em.GetUserId(), em.GetUsername())
	}

	urls, err := em.GetMediaUrls()
	if err != nil {
		log.Println(err)
		return
	}

	for index, url := range urls {
		taggedusers := [][2]string{}
		if len(urls) == 1 {
			taggedusers = em.EdgeMediaToTaggedUser.GetIdUsernamePairs()
		} else {
			taggedusers = em.EdgeSidecarToChildren.Edges[index].Node.EdgeMediaToTaggedUser.GetIdUsernamePairs()
		}

		if saveData {
			saveTaggedUsers(taggedusers)
		}

		filepath := getPostFilePath2(
			em.GetUsername(),
			em.GetUserId(),
			em.GetPostCode(),
			url,
			em.GetTimestamp(),
			taggedusers)
		if index > 0 {
			filepath = appendIndexToFilename(filepath, index)
		}

		CreateFilepathDirIfNotExist(filepath)
		// check if file exist
		if _, err := os.Stat(filepath); os.IsNotExist(err) {
			// file not exists
			printDownloadInfo(&em, url, filepath)
			err = Wget(url, filepath)
			if err != nil {
				log.Println(err)
			} else {
				isDownloaded = true
			}
		} else {
			if err != nil {
				log.Println(err)
			}
		}
	}
	return
}

// Given username, download all posts of the user without login.
func DownloadAllPostsNoLogin(username string) {
	medias, err := instago.GetAllPostMediaNoLogin(username)
	if err != nil {
		log.Println(err)
		return
	}

	for _, media := range medias {
		DownloadPostNoLogin(media.Shortcode)
	}
}

// Given username, download all posts of the user with logged in status.
func (m *IGDownloadManager) DownloadAllPosts(username string) {
	codes, err := m.apimgr.GetAllPostCode(username)
	if err != nil {
		log.Println(err)
		return
	}

	for _, code := range codes {
		m.DownloadPost(code)
	}
}

// Given user name, download recent posts (usually 12 posts) of the user without
// login status.
func DownloadRecentPostsNoLogin(username string) {
	codes, err := instago.GetRecentPostCodeNoLogin(username)
	if err != nil {
		log.Println(err)
		return
	}

	for _, code := range codes {
		DownloadPostNoLogin(code)
	}
}
