package igdl

import (
	"errors"
	"log"
	"os"

	"github.com/siongui/instago"
)

// GetStoryItem downloads the given story item. If the story item already exists
// on local machine, story download will be ignored and isDownloaded is set to
// false.
func GetStoryItem(item instago.IGItem, username string) (isDownloaded bool, err error) {
	return getStoryItem(item, username)
}

func getStoryItem(item instago.IGItem, username string) (isDownloaded bool, err error) {
	if !(item.MediaType == 1 || item.MediaType == 2) {
		err = errors.New("In getStoryItem: not single photo or video!")
		log.Println(err)
		return
	}

	urls, err := item.GetMediaUrls()
	if err != nil {
		log.Println(err)
		return
	}
	if len(urls) != 1 {
		err = errors.New("In getStoryItem: number of download url != 1")
		log.Println(err)
		return
	}
	url := urls[0]

	if saveData {
		saveIdUsername(item.GetUserId(), username)
		saveReelMentions(item.ReelMentions)
	}

	// fix missing username issue while print download info
	item.User.Username = username

	filepath := GetStoryFilePath(
		username,
		item.GetUserId(),
		item.GetPostCode(),
		url,
		item.GetTimestamp(),
		item.ReelMentions)

	CreateFilepathDirIfNotExist(filepath)
	// check if file exist
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		// file not exists
		printDownloadInfo(&item, url, filepath)
		for _, rm := range item.ReelMentions {
			PrintReelMentionInfo(rm)
		}
		err = Wget(url, filepath)
		if err == nil {
			isDownloaded = true
		} else {
			log.Println(err)
			return isDownloaded, err
		}
	} else {
		if err != nil {
			log.Println(err)
		}
	}
	return
}
