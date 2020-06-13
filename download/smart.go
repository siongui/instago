package igdl

import (
	"fmt"
	"log"

	"github.com/siongui/instago"
)

func (m *IGDownloadManager) GetAllPostMediaNoLoginIfPossible(username string) (medias []instago.IGMedia, err error) {
	medias, err = instago.GetAllPostMediaNoLogin(username)
	if err == nil {
		return
	}

	log.Println(err)
	log.Println(username, "cannot download without login")

	if m.mgr2 != nil {
		log.Println("try to get all media infos using clean account")
		medias, err = m.mgr2.GetAllPostMedia(username)
		if err == nil {
			return
		}
		log.Println(err)
		log.Println("fail to get all media infos using clean account")
	}

	// Cannot get without login. also cannot get using clean account.
	// try to download using main account
	return m.GetAllPostMedia(username)
}

// FIXME: only works for public account now
func (m *IGDownloadManager) DownloadAllPostsNoLoginIfPossible(username string) (err error) {
	medias, err := m.GetAllPostMediaNoLoginIfPossible(username)
	if err != nil {
		log.Println(err)
		return
	}

	for _, media := range medias {
		_, err = m.DownloadPostNoLoginIfPossible(media.Shortcode)
		if err != nil {
			// try again
			log.Println(err)
			// check err again?
			m.DownloadPostNoLoginIfPossible(media.Shortcode)
		}
	}
	return
}

func (m *IGDownloadManager) DownloadPostNoLoginIfPossible(code string) (isDownloaded bool, err error) {
	em, err := instago.GetPostInfoNoLogin(code)
	if err == nil {
		return DownloadIGMedia(em)
	}

	log.Println(err)
	log.Println(code, "cannot download without login")

	if m.mgr2 != nil {
		log.Println("try to download using clean account")
		em, err = m.mgr2.GetPostInfo(code)
		if err == nil {
			return DownloadIGMedia(em)
		}
		log.Println(err)
		log.Println("fail to download using clean account")
	}

	// Cannot download without login. also cannot download using clean
	// account. try to download using main account
	em, err = m.GetPostInfo(code)
	if err == nil {
		return DownloadIGMedia(em)
	}
	log.Println(err)
	return
}

func (m *IGDownloadManager) SmartDownloadAllPosts(item instago.IGItem) (err error) {
	if item.User.IsPrivate {
		return m.DownloadAllPosts(item.GetUsername())
	}

	return m.DownloadAllPostsNoLoginIfPossible(item.GetUsername())
}

func (m *IGDownloadManager) SmartDownloadHighlights(item instago.IGItem) (err error) {
	// in case main account is blocked by some users, we use clean account
	// (account not blocked) to download public user account
	if m.mgr2 != nil && !item.User.IsPrivate {
		m.mgr2.DownloadUserStoryHighlights(item.GetUserId())
		return
	}

	m.DownloadUserStoryHighlights(item.GetUserId())
	return
}

func (m *IGDownloadManager) SmartDownloadStory(user instago.User) (err error) {
	// in case main account is blocked by some users, we use clean account
	// (account not blocked) to download public user account
	if m.mgr2 != nil && user.IsPublic() {
		return m.mgr2.downloadUserStoryPostlive(user.GetUserId())
	}

	return m.downloadUserStoryPostlive(user.GetUserId())
}

func (m *IGDownloadManager) SmartDownloadPost(item instago.IGItem) (isDownloaded bool, err error) {
	if item.User.IsPrivate {
		return m.DownloadPost(item.GetPostCode())
	}

	return m.DownloadPostNoLoginIfPossible(item.GetPostCode())
}

func (m *IGDownloadManager) smartGetStoryItemLayer(item instago.IGItem, user instago.User, layer int, isdone map[string]string) {
	getStoryItem(item, user.GetUsername())
	for _, rm := range item.ReelMentions {
		fmt.Print("  ")
		PrintReelMentionInfo(rm)
		if !rm.User.IsPublic() {
			//UsernameIdColorPrint(rm.GetUsername(), rm.GetUserId())
			//log.Println("is private. ignored.")
			continue
		}
		m.smartDownloadUserStoryPostliveLayer(rm, layer, isdone)
	}
}

func (m *IGDownloadManager) smartDownloadUserStoryPostliveLayer(user instago.User, layer int, isdone map[string]string) (err error) {
	if layer < 1 {
		return
	}
	layer--

	id := user.GetUserId()
	if username, ok := isdone[id]; ok {
		UsernameIdColorPrint("  "+username, id)
		fmt.Println("'s stories and postlives are already fetched. ignored")
		return
	} else {
		fmt.Print("  Try to fetch metadata of stories and postlives of ")
		UsernameIdColorPrint(user.GetUsername(), id)
		fmt.Println("")
	}

	ut := instago.UserTray{}
	if user.IsPublic() && m.mgr2 != nil {
		ut, err = m.mgr2.GetUserReelMedia(id)
	} else {
		ut, err = m.GetUserReelMedia(id)
	}

	if err != nil {
		log.Println(err)
		return
	}
	tray := ut.Reel

	// FIXME: tray.GetUsername() is ""
	//isdone[id] = tray.GetUsername()
	isdone[id] = user.GetUsername()
	UsernameIdColorPrint("  "+user.GetUsername(), id)
	fmt.Println("'s metadata of stories and postlives are fetched successfully")

	for _, item := range tray.GetItems() {
		m.smartGetStoryItemLayer(item, user, layer, isdone)
	}

	return DownloadPostLiveItem(ut.PostLiveItem)
}

func (m *IGDownloadManager) SmartDownloadUserStoryPostliveLayer(user instago.User, layer int) (err error) {
	PrintUserInfo(user)
	isdone := make(map[string]string)
	return m.smartDownloadUserStoryPostliveLayer(user, layer, isdone)
}

func (m *IGDownloadManager) SmartGetUserReelMedia(user instago.User) (instago.UserTray, error) {
	if user.IsPublic() && m.mgr2 != nil {
		return m.mgr2.GetUserReelMedia(user.GetUserId())
	}
	return m.GetUserReelMedia(user.GetUserId())
}
