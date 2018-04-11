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

func DownloadPost(code string, mgr *instago.IGApiManager) {
	em, err := mgr.GetPostInfo(code)
	if err != nil {
		log.Println(err)
		return
	}

	urls := em.GetMediaUrls()

	for index, url := range urls {
		filepath := GetPostFilePath(
			em.GetUsername(),
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
			}
		}
	}
}

func DownloadAllPosts(username string, mgr *instago.IGApiManager) {
	codes, err := mgr.GetAllPostCode(username)
	if err != nil {
		log.Println(err)
		return
	}

	for _, code := range codes {
		DownloadPost(code, mgr)
	}
}
