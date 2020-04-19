package igdl

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/siongui/instago"
)

func (m *IGDownloadManager) GetSelfFollowers() (users []instago.IGFollowUser, err error) {
	return m.apimgr.GetFollowers(m.apimgr.GetSelfId())
}

func (m *IGDownloadManager) SaveSelfFollowers(filepath string) (err error) {
	users, err := m.GetSelfFollowers()
	if err != nil {
		log.Println(err)
		return
	}

	b, err := json.Marshal(users)
	if err != nil {
		log.Println(err)
		return
	}

	return ioutil.WriteFile(filepath, b, 0644)
}
