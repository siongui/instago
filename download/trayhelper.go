package igdl

import (
	"log"
	"os"
	"strconv"
)

type TrayInfo struct {
	Id              int64
	Username        string
	Layer           int64
	IsPrivate       bool
	LatestReelMedia int64
	//Tray     instago.IGReelTray
}

func SetupTrayInfo(id int64, username string, layer int64, isPrivate bool, latestReelMedia int64) (ti TrayInfo) {
	return TrayInfo{
		Id:              id,
		Username:        username,
		Layer:           layer,
		IsPrivate:       isPrivate,
		LatestReelMedia: latestReelMedia,
	}
}

func IsTrayInfoInQueue(queue []TrayInfo, ti TrayInfo) bool {
	for _, t := range queue {
		if t.Id == ti.Id {
			return true
		}
	}
	return false
}

func GetTrayInfoFromQueue(queue []TrayInfo, id int64) (ti TrayInfo, ok bool) {
	for _, t := range queue {
		if t.Id == id {
			ti = t
			ok = true
			return
		}
	}
	return
}

func IsLatestReelMediaExist(username string, latestReelMedia int64) bool {
	utimes, err := GetReelMediaUnixTimesInUserStoryDir(username)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Println("In IsLatestReelMediaExist", err)
		}
		return false
	}

	lrm := strconv.FormatInt(latestReelMedia, 10)
	for _, utime := range utimes {
		if lrm == utime {
			return true
		}
	}
	return false
}
