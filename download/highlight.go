package igdl

import (
	"fmt"
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
		fmt.Println(user.Username, ": ", user.Pk)
	}
}
