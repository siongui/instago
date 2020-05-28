package igdl

import (
	"github.com/siongui/instago"
)

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
		isDownloaded, err = m.DownloadPost(item.GetPostCode())
		return
	}

	isDownloaded, err = DownloadPostNoLogin(item.GetPostCode())
	if err == nil {
		return
	}

	if m.mgr2 != nil {
		isDownloaded, err = m.mgr2.DownloadPost(item.GetPostCode())
	}
	return
}
