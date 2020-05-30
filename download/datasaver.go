package igdl

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/siongui/instago"
)

var saveData = false

// Default is false (not save data).
func SetSaveData(b bool) {
	saveData = b
}

// Given id, get username from mobile API endpoint
func (m *IGDownloadManager) IdToUsername(id string) (username string, err error) {
	user, err := m.GetUserInfoEndPoint(id)
	if err == nil {
		username = user.Username
	}
	return
}

func (m *IGDownloadManager) UsernameToId(username string) (id string, err error) {
	// Try to get id from local saved data
	if m.idusernames != nil {
		id = FindIdFromUsernameInMap(m.idusernames, username)
		if id != "" {
			// double check in case someone change username
			u, err := m.IdToUsername(id)
			if err == nil && u == username {
				return id, err
			}
		}
		log.Println("fail to get id from local saved data")
	}

	// Try to get id without loggin
	id, err = instago.GetUserId(username)
	if err == nil {
		if saveData {
			saveIdUsername(id, username)
		}
		return
	}

	// Try to get id with loggin
	ui, err := m.GetUserInfo(username)
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
	p := GetIdUsernamePath(id, username)
	return saveEmpty(p)
}

func saveReelMentions(rms []instago.ItemReelMention) (err error) {
	for _, rm := range rms {
		saveIdUsername(rm.GetUserId(), rm.GetUsername())
		p := GetReelMentionsPath(rm.GetUserId(), rm.GetUsername())
		err = saveEmpty(p)
	}
	// DISCUSS: err returned here seems useless
	return
}

func saveTaggedUsers(taggedusers [][2]string) (err error) {
	for _, user := range taggedusers {
		err = saveIdUsername(user[0], user[1])
	}
	return
}

func BuildIdUsernameMapFromLocalData(idusernamedir string) (m map[string][]string, err error) {
	m = make(map[string][]string)

	infos, err := ioutil.ReadDir(idusernamedir)
	if err != nil {
		return
	}

	for _, info := range infos {
		//fmt.Println(info.Name())

		a := strings.Split(info.Name(), "-")
		if len(a) != 2 {
			panic(info.Name())
		}
		id := a[0]
		username := a[1]

		if usernames, ok := m[id]; ok {
			m[id] = append(usernames, username)
			//fmt.Println(id, m[id])
		} else {
			m[id] = []string{username}
		}
	}

	return
}

func FindIdFromUsernameInMap(m map[string][]string, username string) (id string) {
	for i, usernames := range m {
		for _, u := range usernames {
			if u == username {
				id = i
				return
			}
		}
	}
	return
}

func (m *IGDownloadManager) FindNewUsernameFromOldName(idUsernameMap map[string][]string, oldname string) (username string) {
	id := FindIdFromUsernameInMap(idUsernameMap, oldname)
	if id == "" {
		fmt.Println("fail to look up id locally")
		return
	}
	fmt.Println("old username:", oldname, "id:", id)

	name, err := m.IdToUsername(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	username = name
	fmt.Println("old username:", oldname, "id:", id, "new username:", username)
	return
}
