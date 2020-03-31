package igdl

import (
	"fmt"
	"strconv"
	"time"
)

// Given username, download story highlights of the user.
func (m *IGDownloadManager) DownloadUserStoryHighlightsByName(username string) {
	id, err := m.UsernameToId(username)
	if err != nil {
		panic(err)
	}

	m.DownloadUserStoryHighlights(id)
}

// Download story highlights of a single user.
func (m *IGDownloadManager) DownloadUserStoryHighlights(userid string) {
	trays, err := m.apimgr.GetAllStoryHighlights(userid)
	if err != nil {
		panic(err)
	}
	for _, tray := range trays {
		for _, item := range tray.GetItems() {
			getStoryItem(item, tray.GetUsername())
		}
	}
}

// Download story highlights of all following users
func (m *IGDownloadManager) DownloadStoryHighlights() {
	users, err := m.apimgr.GetFollowing(m.apimgr.GetSelfId())
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, user := range users {
		fmt.Println("Fetching", user.Username, user.Pk, "story highlights ...")
		userid := strconv.FormatInt(user.Pk, 10)
		m.DownloadUserStoryHighlights(userid)
		fmt.Println("Fetching", user.Username, user.Pk, "story highlights done!")

		// sleep 3 seconds to prevent http 429
		time.Sleep(3 * time.Second)
	}
}
