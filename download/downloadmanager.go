// Package igdl helps you download posts, stories and story highlights of
// Instagram users.
package igdl

import (
	"errors"

	"github.com/siongui/instago"
)

type IGDownloadManager struct {
	apimgr *instago.IGApiManager
}

// The arguments here is the same as the NewInstagramApiManager of instago.
// See README of instago for more informantion
func NewInstagramDownloadManager(ds_user_id, sessionid, csrftoken string) (mgr *IGDownloadManager, err error) {
	if !IsCommandAvailable("wget") {
		err = errors.New("Please install wget")
		return
	}
	if !IsCommandAvailable("ffmpeg") {
		err = errors.New("Please install ffmpeg")
		return
	}

	mgr = &IGDownloadManager{
		apimgr: instago.NewInstagramApiManager(ds_user_id, sessionid, csrftoken),
	}
	return
}
