package igdl

import (
	"log"
)

// DownloadStoryOfMultipleId downloads stories of multiple user id via reels
// media API. If a newly created account use this method to download stories,
// it may cause Instagram to block your account. Be careful before use. Also
// note that the max length of multipleIds allowed in the API call  is between
// 20 to 30.
func (m *IGDownloadManager) DownloadStoryOfMultipleId(multipleIds []string) (err error) {
	trays, err := m.GetMultipleReelsMedia(multipleIds)
	if err != nil {
		log.Println(err)
		return
	}

	for _, tray := range trays {
		for _, item := range tray.Items {
			username := tray.User.GetUsername()
			id := tray.User.GetUserId()
			_, err = GetStoryItem(item, username)
			if err != nil {
				PrintUsernameIdMsg(username, id, err)
				return
			}
		}
	}

	return
}
