package igdl

import (
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

func (m *IGDownloadManager) SmartDownloadStory(user instago.IGUser) (err error) {
	// Pk here is user id
	if user.IsPrivate {
		return m.DownloadUserStoryPostlive(user.Pk)
	}

	// in case main account is blocked by some users, we use clean account
	// (account not blocked) to download public user account
	if m.mgr2 != nil {
		return m.mgr2.DownloadUserStoryPostlive(user.Pk)
	}

	return m.DownloadUserStoryPostlive(user.Pk)
}

func (m *IGDownloadManager) SmartDownloadPost(item instago.IGItem) (isDownloaded bool, err error) {
	if item.User.IsPrivate {
		return m.DownloadPost(item.GetPostCode())
	}

	return m.DownloadPostNoLoginIfPossible(item.GetPostCode())
}

func (m *IGDownloadManager) smartGetStoryItemLayer(item instago.IGItem, username, id string, layer int, isdone map[string]string) {
	getStoryItem(item, username)
	for _, rm := range item.ReelMentions {
		PrintReelMentionInfo(rm)
		if !rm.User.IsPublic() {
			//UsernameIdColorPrint(rm.GetUsername(), rm.GetUserId())
			//log.Println("is private. ignored.")
			continue
		}
		m.smartDownloadUserStoryPostliveLayer(rm.GetUsername(), rm.GetUserId(), layer, isdone)
	}
}

func (m *IGDownloadManager) smartDownloadUserStoryPostliveLayer(username, id string, layer int, isdone map[string]string) (err error) {
	if layer < 1 {
		return
	}
	layer--

	if username2, ok := isdone[id]; ok {
		UsernameIdColorPrint(username2, id)
		log.Println("'s stories and postlives are already fetched")
		return
	} else {
		log.Print("fetching stories and postlives of")
		UsernameIdColorPrint(username, id)
		log.Println("")
	}

	ut, err := m.GetUserReelMedia(id)
	if err != nil {
		log.Println(err)
		return
	}
	tray := ut.Reel

	isdone[id] = tray.GetUsername()
	UsernameIdColorPrint(tray.GetUsername(), id)
	log.Println("'s metadata of stories and postlives are fetched successfully")

	for _, item := range tray.GetItems() {
		m.smartGetStoryItemLayer(item, tray.GetUsername(), id, layer, isdone)
	}

	return DownloadPostLiveItem(ut.PostLiveItem)
}

func (m *IGDownloadManager) SmartDownloadUserStoryPostliveLayer(user instago.User, layer int) (err error) {
	isdone := make(map[string]string)
	if user.IsPublic() {
		if m.mgr2 != nil {
			return m.mgr2.smartDownloadUserStoryPostliveLayer(user.GetUsername(), user.GetUserId(), layer, isdone)
		}
	}
	return m.smartDownloadUserStoryPostliveLayer(user.GetUsername(), user.GetUserId(), layer, isdone)
}
