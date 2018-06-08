package instago

// This file decode JSON data returned by Instagram saved posts API endpoint

import (
	"encoding/json"
)

const urlSaved = `https://i.instagram.com/api/v1/feed/saved/`

type savedPostsResp struct {
	Items []struct {
		Item IGItem `json:"media"`
	} `json:"items"`
	NumResults          int64  `json:"num_results"`
	MoreAvailable       bool   `json:"more_available"`
	AutoLoadMoreEnabled bool   `json:"auto_load_more_enabled"`
	NextMaxId           string `json:"next_max_id"`
	Status              string `json:"status"`
}

// GetSavedPosts returns your saved posts.
func (m *IGApiManager) GetSavedPosts() (items []IGItem, err error) {
	b, err := getHTTPResponse(urlSaved, m.dsUserId, m.sessionid, m.csrftoken)
	if err != nil {
		return
	}

	spp := savedPostsResp{}
	err = json.Unmarshal(b, &spp)
	if err != nil {
		return
	}
	for _, item := range spp.Items {
		items = append(items, item.Item)
	}

	for spp.MoreAvailable {
		url := urlSaved + "?max_id=" + spp.NextMaxId
		b, err = getHTTPResponse(url, m.dsUserId, m.sessionid, m.csrftoken)
		if err != nil {
			return
		}
		spp = savedPostsResp{}
		err = json.Unmarshal(b, &spp)
		if err != nil {
			return
		}
		for _, item := range spp.Items {
			items = append(items, item.Item)
		}

	}

	return
}
