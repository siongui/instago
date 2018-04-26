package igdl

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/siongui/instago"
)

func printProfilePicDownloadInfo(username, url, filepath string, timestamp int64) {
	fmt.Print("username: ")
	cc.Println(username)
	fmt.Print("time: ")
	cc.Println(timestamp)
	fmt.Print("profile pic url: ")
	cc.Println(url)

	fmt.Print("Download ")
	rc.Print(url)
	fmt.Print(" to ")
	cc.Println(filepath)
}

// Given user name, download user profile pic hd (usually 320x320)
func DownloadUserProfilePicUrlHd(username string) {
	ui, err := instago.GetUserInfoNoLogin(username)
	if err != nil {
		log.Println(err)
		return
	}
	timestamp := time.Now().Unix()
	url := ui.ProfilePicUrlHd
	filepath := getUserProfilPicFilePath(username, ui.Id, url, timestamp)

	// check if file exist
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		// file not exists
		printProfilePicDownloadInfo(username, url, filepath, timestamp)
		CreateFilepathDirIfNotExist(filepath)
		err = Wget(url, filepath)
		if err != nil {
			log.Println(err)
		}
	}
}
