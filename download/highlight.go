package igdl

import (
	"fmt"
	"strconv"

	"github.com/siongui/instago"
)

// Download story highlights of all following users
func DownloadStoryHighlights(mgr *instago.IGApiManager) {
	users, err := mgr.GetFollowing(mgr.GetSelfId())
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, user := range users {
		//fmt.Println(user.Username, ": ", user.Pk)
		userid := strconv.FormatInt(user.Pk, 10)
		trays, err := mgr.GetAllStoryHighlights(userid)
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
