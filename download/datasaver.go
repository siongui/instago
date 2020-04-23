package igdl

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/siongui/instago"
)

var saveData = false

// Default is false (not save data).
func SetSaveData(b bool) {
	saveData = b
}

func (m *IGDownloadManager) UsernameToId(username string) (id string, err error) {
	// Try to get id without loggin
	id, err = instago.GetUserId(username)
	if err == nil {
		if saveData {
			saveIdUsername(id, username)
		}
		return
	}

	// Try to get id with loggin
	ui, err := m.apimgr.GetUserInfo(username)
	if err == nil {
		id = ui.Id
		if saveData {
			saveIdUsername(id, username)
		}
	}
	return
}

func saveEmpty(p string) (err error) {
	CreateFilepathDirIfNotExist(p)
	// check if file exist
	if _, err := os.Stat(p); os.IsNotExist(err) {
		// file not exists
		err = ioutil.WriteFile(p, []byte(""), 0644)
		if err == nil {
			fmt.Println(p, "saved")
		}
	}
	return
}

func saveIdUsername(id, username string) (err error) {
	p := getIdUsernamePath(id, username)
	return saveEmpty(p)
}

func saveReelMentions(rms []instago.ItemReelMention) (err error) {
	for _, rm := range rms {
		saveIdUsername(rm.GetUserId(), rm.GetUsername())
		p := getReelMentionsPath(rm.GetUserId(), rm.GetUsername())
		err = saveEmpty(p)
	}
	// DISCUSS: err returned here seems useless
	return
}
