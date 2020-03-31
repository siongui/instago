package igdl

import (
	"fmt"
	"log"
	"os"

	"github.com/siongui/instago"
)

func printDownloadInfo(pi instago.PostItem, url, filepath string) {
	fmt.Print("username: ")
	// FIXME: username of story item is missing
	cc.Println(pi.GetUsername())
	fmt.Print("time: ")
	cc.Println(formatTimestamp(pi.GetTimestamp()))
	fmt.Print("post url: ")
	cc.Println(pi.GetPostUrl())

	fmt.Print("Download ")
	rc.Print(url)
	fmt.Print(" to ")
	cc.Println(filepath)
}

func DownloadPostNoLogin(code string) (isDownloaded bool, err error) {
	em, err := instago.GetPostInfoNoLogin(code)
	if err != nil {
		log.Println(err)
		return
	}

	return DownloadPostItem(&em)
}

func (m *IGDownloadManager) DownloadPost(code string) (isDownloaded bool, err error) {
	em, err := m.apimgr.GetPostInfo(code)
	if err != nil {
		log.Println(err)
		return
	}

	return DownloadPostItem(&em)
}

// TODO: try to merge getPostItem and DownloadPostItem
//
// DownloadPostItem downloads photos/videos in the post.
// IGItem (items in timeline or saved posts) or IGMedia (read from
// https://www.instagram.com/p/{{CODE}}/?__a=1) can be argument in this method.
func DownloadPostItem(pi instago.PostItem) (isDownloaded bool, err error) {
	urls, err := pi.GetMediaUrls()
	if err != nil {
		log.Println(err)
		return
	}

	for index, url := range urls {
		filepath := getPostFilePath(
			pi.GetUsername(),
			pi.GetUserId(),
			pi.GetPostCode(),
			url,
			pi.GetTimestamp())
		if index > 0 {
			filepath = appendIndexToFilename(filepath, index)
		}

		CreateFilepathDirIfNotExist(filepath)
		// check if file exist
		if _, err := os.Stat(filepath); os.IsNotExist(err) {
			// file not exists
			printDownloadInfo(pi, url, filepath)
			err = Wget(url, filepath)
			if err != nil {
				log.Println(err)
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
