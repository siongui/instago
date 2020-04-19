package instago

// Get following and followers

import (
	"encoding/json"
	"strconv"
	"strings"
)

const urlFollowers = `https://i.instagram.com/api/v1/friendships/{{USERID}}/followers/`
const urlFollowing = `https://i.instagram.com/api/v1/friendships/{{USERID}}/following/`

type rawFollow struct {
	//sections
	//global_blacklist_sample
	Users     []IGFollowUser `json:"users"`
	BigList   bool           `json:"big_list"`    // if false, no next_max_id in response
	NextMaxId int64          `json:"next_max_id"` // used for pagination if list is too big
	PageSize  int64          `json:"page_size"`
	Status    string         `json:"status"`
}

// GetFollowers returns all followers of the given user id.
func (m *IGApiManager) GetFollowers(id string) (users []IGFollowUser, err error) {
	url := strings.Replace(urlFollowers, "{{USERID}}", id, 1)
	return m.getFollow(url)
}

// GetFollowing returns all following users of the given user id.
func (m *IGApiManager) GetFollowing(id string) (users []IGFollowUser, err error) {
	url := strings.Replace(urlFollowing, "{{USERID}}", id, 1)
	return m.getFollow(url)
}

func (m *IGApiManager) getFollow(url string) (users []IGFollowUser, err error) {
	rf, err := m.getFollowResponse(url)
	if err != nil {
		return
	}
	users = append(users, rf.Users...)

	// If the list is too big and next_max_id is not ""
	for rf.NextMaxId != 0 {
		urln := url + "?max_id=" + strconv.FormatInt(rf.NextMaxId, 10)
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

	// for development purpose
	if saveRawJsonByte {
		SaveRawJsonByte("follow-", b)
	}

	err = json.Unmarshal(b, &rf)
	return
}
