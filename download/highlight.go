package igdl

import (
	"fmt"
	"strconv"
	"time"
)

// Download story highlights of all following users
func (m *IGDownloadManager) DownloadStoryHighlights() {
	users, err := m.apimgr.GetFollowing(m.apimgr.GetSelfId())
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, user := range users {
		fmt.Println("Fetching", user.Username, user.Pk, "stroy highlights ...")
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
		fmt.Println("Fetching", user.Username, user.Pk, "stroy highlights done!")

		// sleep 3 seconds to prevent http 429
		time.Sleep(3 * time.Second)
	}
}
