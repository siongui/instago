// Package instago helps you access the private API of Instagram. For examle,
// get URLs of all posts of a specific Instagram user, media (photos and videos)
// links of posts, stories of a Instagram user, your following and followers.
package instago

import (
	"encoding/json"
	"io/ioutil"
)

type IGApiManager struct {
	cookies map[string]string
}

// NewInstagramApiManager returns a API manager of a logged-in Instagram user,
// given the JSON file of cookies of a Instagram logged-in account.
//
// The cookies, such as *ds_user_id*, *sessionid*, or *csrftoken* can be viewed
// in Chrome Developer Tools. See https://stackoverflow.com/a/44773079
//
// You can get the JSON file of cookies using chrome extension in crx-cookies/
// directory.
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

// GetSelfId returns the id of a Instagram user of the API manager.
func (m *IGApiManager) GetSelfId() string {
	id, _ := m.cookies["ds_user_id"]
	return id
}
