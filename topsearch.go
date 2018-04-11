package instago

// This file implements topsearch of Instagram web.

import (
	"encoding/json"
)

const urlTopsearch = `https://www.instagram.com/web/search/topsearch/?query=`

// Decode JSON data returned from Instagram topsearch API
type IGTopsearchResp struct {
	Users []struct {
		Position int64  `json:"position"`
		User     IGUser `json:"user"`
	} `json:"users"`

	//TODO: Places ... `json:"places"`

	Hashtags []struct {
		Position int64 `json:"position"`
		Hashtag  struct {
			Name       string `json:"name"`
			Id         int64  `json:"id"`
			MediaCount int64  `json:"media_count"`
		} `json:"hashtag"`
	} `json:"hashtags"`
	HasMore          bool   `json:"has_more"`
	RankToken        string `json:"rank_token"`
	ClearClientCache bool   `json:"clear_client_cache"`
	Status           string `json:"status"`
}

// Given a string, return the users that best matches the string. This is
// actually *topsearch* on Instagram web.
func (m *IGApiManager) Topsearch(str string) (tr IGTopsearchResp, err error) {
	url := urlTopsearch + str
	b, err := getHTTPResponse(url, m.dsUserId, m.sessionid, m.csrftoken)
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &tr)
	return
}
