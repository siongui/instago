package instago

// Get following and followers

import (
	"encoding/json"
	"strings"
)

const urlFollowers = `https://i.instagram.com/api/v1/friendships/{{USERID}}/followers/`
const urlFollowing = `https://i.instagram.com/api/v1/friendships/{{USERID}}/following/`

type rawFollow struct {
	Users     []IGUser `json:"users"`
	BigList   bool     `json:"big_list"` // if false, no next_max_id in response
	PageSize  int64    `json:"page_size"`
	Status    string   `json:"status"`
	NextMaxId string   `json:"next_max_id"` // used for pagination if list is too big
}

// GetFollowers returns all followers of the given user id.
func (m *IGApiManager) GetFollowers(id string) (users []IGUser, err error) {
	url := strings.Replace(urlFollowers, "{{USERID}}", id, 1)
	return m.getFollow(url)
}

// GetFollowing returns all following users of the given user id.
func (m *IGApiManager) GetFollowing(id string) (users []IGUser, err error) {
	url := strings.Replace(urlFollowing, "{{USERID}}", id, 1)
	return m.getFollow(url)
}

func (m *IGApiManager) getFollow(url string) (users []IGUser, err error) {
	rf, err := m.getFollowResponse(url)
	if err != nil {
		return
	}
	users = append(users, rf.Users...)

	// If the list is too big and next_max_id is not ""
	for rf.NextMaxId != "" {
		urln := url + "?max_id=" + rf.NextMaxId
		rfn, err := m.getFollowResponse(urln)
		if err != nil {
			return users, err
		}
		users = append(users, rfn.Users...)
		rf.NextMaxId = rfn.NextMaxId
	}
	return
}

func (m *IGApiManager) getFollowResponse(url string) (rf rawFollow, err error) {
	b, err := m.getHTTPResponse(url, "GET")
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &rf)
	return
}
