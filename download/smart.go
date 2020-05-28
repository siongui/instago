package igdl

import (
	"log"

	"github.com/siongui/instago"
)

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

func (m *IGDownloadManager) SmartDownloadAllPosts(username string) (err error) {
	err = DownloadAllPostsNoLogin(username)
	if err == nil {
		return
	}

	if m.mgr2 != nil {
		err = m.mgr2.DownloadAllPosts(username)
		//err = m.mgr2.DownloadAllPostsNoLogin(username)
	}
	return
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
