// Package instago helps you get all URLs of posts of a specific Instagram user,
// media (photos and videos) links of posts, stories of following user,
// following and followers.
package instago

import (
	"encoding/json"
	"io/ioutil"
)

type IGApiManager struct {
	cookies map[string]string
}

// After login to Instagram, you can get the cookies of *ds_user_id*,
// *sessionid*, *csrftoken* in Chrome Developer Tools.
// See https://stackoverflow.com/a/44773079
// or
// https://github.com/hoschiCZ/instastories-backup#obtain-cookies
func NewInstagramApiManager(authFilePath string) (*IGApiManager, error) {
	var m IGApiManager
	b, err := ioutil.ReadFile(authFilePath)
	if err != nil {
		return &m, err
	}

	m.cookies = make(map[string]string)
	err = json.Unmarshal(b, &m.cookies)
	return &m, err
}

func (m *IGApiManager) GetSelfId() string {
	id, _ := m.cookies["ds_user_id"]
	return id
}
