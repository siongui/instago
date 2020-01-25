package instago

// This file decode JSON data returned by Instagram saved posts API endpoint

import (
	"encoding/json"
	"log"
	"time"
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

// GetSavedPosts returns your saved posts. Pass -1 will get all saved posts.
func (m *IGApiManager) GetSavedPosts(numOfItem int) (items []IGItem, err error) {
	b, err := m.getHTTPResponse(urlSaved, "GET")
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
		if numOfItem > 0 && len(items) > numOfItem {
			break
		}

		url := urlSaved + "?max_id=" + spp.NextMaxId
		b, err = m.getHTTPResponse(url, "GET")
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
		log.Println("fetched", len(items), "items")
		// sleep 500ms to prevent http 429
		time.Sleep(500 * time.Millisecond)
	}

	return
}
