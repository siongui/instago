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

func saveIdUsername(id, username string) (err error) {
	p := getIdUsernamePath(id, username)
	CreateFilepathDirIfNotExist(p)
	// check if file exist
	if _, err := os.Stat(p); os.IsNotExist(err) {
		// file not exists
		err = ioutil.WriteFile(p, []byte(""), 0644)
		if err == nil {
			fmt.Println(id, username, "saved")
		}
	}
	return
}
