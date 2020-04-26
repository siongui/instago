package igdl

import (
	"errors"
	"os"
	"strings"
)

// Given user name, get the location of user data dir of Chromium of Ubuntu snap
// version.
func GetUserDataDirChromiumSnap(username string) (userdatadir string, err error) {
	// snap version chromium user data dir
	// see https://askubuntu.com/a/1075319
	userdatadir = strings.Replace(
		"/home/<USER>/snap/chromium/current/.config/chromium",
		"<USER>", username, 1)

	fi, err := os.Stat(userdatadir)
	if err != nil {
		return
	}

	if !fi.IsDir() {
		err = errors.New("cannot find user data dir")
		return
	}

	return
}
