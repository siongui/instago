package igdl

import (
	"fmt"
	"strconv"
)

// Download story highlights of all following users
func (m *IGDownloadManager) DownloadStoryHighlights() {
	users, err := m.apimgr.GetFollowing(m.apimgr.GetSelfId())
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, user := range users {
		//fmt.Println(user.Username, ": ", user.Pk)
		userid := strconv.FormatInt(user.Pk, 10)
		trays, err := m.apimgr.GetAllStoryHighlights(userid)
		if err != nil {
			panic(err)
		}
		for _, tray := range trays {
			for _, item := range tray.GetItems() {
				getStoryItem(item)
			}
		}
	}
}
