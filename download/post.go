package igdl

import (
	"fmt"
	"github.com/siongui/instago"
	"log"
	"os"
)

func printPostDownloadInfo(em instago.IGMedia, url, filepath string) {
	fmt.Print("username: ")
	cc.Println(em.GetUsername())
	fmt.Print("time: ")
	cc.Println(em.GetTimestamp())
	fmt.Print("post url: ")
	cc.Println(em.GetPostUrl())

	fmt.Print("Download ")
	rc.Print(url)
	fmt.Print(" to ")
	cc.Println(filepath)
}

func DownloadPostNoLogin(code string) {
	em, err := instago.GetPostInfoNoLogin(code)
	if err != nil {
		log.Println(err)
		return
	}
	DownloadIGMedia(em)
}

func (m *IGDownloadManager) DownloadPost(code string) (isDownloaded bool) {
	em, err := m.apimgr.GetPostInfo(code)
	if err != nil {
		log.Println(err)
		return
	}
	isDownloaded = DownloadIGMedia(em)
	return
}

// TODO: try to merge getPostItem and DownloadIGMedia
func DownloadIGMedia(em instago.IGMedia) (isDownloaded bool) {
	urls := em.GetMediaUrls()

	for index, url := range urls {
		filepath := getPostFilePath(
			em.GetUsername(),
			em.GetUserId(),
			em.GetPostCode(),
			url,
			em.GetTimestamp())
		if index > 0 {
			filepath = appendIndexToFilename(filepath, index)
		}

		CreateFilepathDirIfNotExist(filepath)
		// check if file exist
		if _, err := os.Stat(filepath); os.IsNotExist(err) {
			// file not exists
			printPostDownloadInfo(em, url, filepath)
			err = Wget(url, filepath)
			if err != nil {
				fmt.Println(err)
			} else {
				isDownloaded = true
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
