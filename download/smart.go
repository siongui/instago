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
		// check error?
		m.DownloadPostNoLoginIfPossible(media.Shortcode)
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
