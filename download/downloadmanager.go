// Package igdl helps you download posts, stories and story highlights of
// Instagram users.
package igdl

import (
	"errors"

	"github.com/siongui/instago"
)

type IGDownloadManager struct {
	apimgr      *instago.IGApiManager
	collections []instago.Collection
	idusernames map[string][]string

	// manager of clean (non-blocked) account
	mgr2 *IGDownloadManager
}

// The arguments here is the same as the NewInstagramApiManager of instago.
// See README of instago for more informantion
func NewInstagramDownloadManager(authFilePath string) (*IGDownloadManager, error) {
	var m IGDownloadManager
	var err error
	if !IsCommandAvailable("wget") {
		err = errors.New("Please install wget")
		return &m, err
	}
	if !IsCommandAvailable("ffmpeg") {
		err = errors.New("Please install ffmpeg")
		return &m, err
	}

	m.apimgr, err = instago.NewInstagramApiManager(authFilePath)

	// TODO: check m.apimgr here?

	return &m, err
}

func (m *IGDownloadManager) GetSelfId() string {
	return m.apimgr.GetSelfId()
}

func (m *IGDownloadManager) LoadCollectionList() (err error) {
	m.collections, err = m.apimgr.GetSavedCollectionList()
	return
}

func (m *IGDownloadManager) GetReelsTray() (instago.IGReelsTray, error) {
	return m.apimgr.GetReelsTray()
}

func (m *IGDownloadManager) GetUserInfoEndPoint(id string) (instago.UserInfoEndPoint, error) {
	return m.apimgr.GetUserInfoEndPoint(id)
}

func (m *IGDownloadManager) GetPostInfo(code string) (instago.IGMedia, error) {
	return m.apimgr.GetPostInfo(code)
}

func (m *IGDownloadManager) GetAllPostMedia(username string) ([]instago.IGMedia, error) {
	return m.apimgr.GetAllPostMedia(username)
}

func (m *IGDownloadManager) GetUserInfo(username string) (instago.UserInfo, error) {
	return m.apimgr.GetUserInfo(username)
}

func (m *IGDownloadManager) GetUserStory(id string) (instago.UserTray, error) {
	return m.apimgr.GetUserStory(id)
}

func (m *IGDownloadManager) GetSavedPosts(numOfItem int) ([]instago.IGItem, error) {
	return m.apimgr.GetSavedPosts(numOfItem)
}

// Optional. Load another account which is not blocked by any other account.
func (m *IGDownloadManager) LoadCleanDownloadManager(authFilePath string) (err error) {
	m2, err := NewInstagramDownloadManager(authFilePath)
	if err == nil {
		m.mgr2 = m2
	}
	return
}

func (m *IGDownloadManager) IsCleanAccountSet() bool {
	return m.mgr2 != nil
}

func (m *IGDownloadManager) GetCleanAccountManager() *IGDownloadManager {
	return m.mgr2
}

func (m *IGDownloadManager) LoadIdUsernameFromDataDir() (err error) {
	iu, err := BuildIdUsernameMapFromLocalData(GetIdUsernameDir())
	if err == nil {
		m.idusernames = iu
	}
	return
}
