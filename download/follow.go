package igdl

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/siongui/instago"
)

func (m *IGDownloadManager) GetSelfFollowers() ([]instago.IGFollowUser, error) {
	return m.apimgr.GetFollowers(m.apimgr.GetSelfId())
}

func (m *IGDownloadManager) GetSelfFollowing() ([]instago.IGFollowUser, error) {
	return m.apimgr.GetFollowing(m.apimgr.GetSelfId())
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

func (m *IGDownloadManager) SaveSelfFollowing(filepath string) (err error) {
	users, err := m.GetSelfFollowing()
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

// Load []instago.IGFollowUser from JSON file, which is created by
// SaveSelfFollowers or SaveSelfFollowing methods.
func LoadFollowUsers(filepath string) (users []instago.IGFollowUser, err error) {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Println(err)
		return
	}

	err = json.Unmarshal(b, &users)
	return
}
