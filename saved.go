package instago

// This file decode JSON data returned by Instagram saved posts API endpoint

import (
	"encoding/json"
	"log"
	"strings"
	"time"
)

const urlSaved = `https://i.instagram.com/api/v1/feed/saved/`
const urlSavedCollection = `https://i.instagram.com/api/v1/feed/collection/{{COLLECTIONID}}/`
const urlSavedCollectionList = `https://i.instagram.com/api/v1/collections/list/`

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

type Collection struct {
	CollectionId         string `json:"collection_id"`
	CollectionName       string `json:"collection_name"`
	CollectionType       string `json:"collection_type"`
	CollectionMediaCount int64  `json:"collection_media_count"`
	//TODO: cover_media
}

type collectionsList struct {
	Collections         []Collection `json:"items"`
	MoreAvailable       bool         `json:"more_available"`
	AutoLoadMoreEnabled bool         `json:"auto_load_more_enabled"`
	NextMaxId           string       `json:"next_max_id"`
	Status              string       `json:"status"`
}

// GetSavedPosts returns your saved posts. Pass -1 will get all saved posts.
func (m *IGApiManager) GetSavedPosts(numOfItem int) (items []IGItem, err error) {
	b, err := m.getHTTPResponse(urlSaved, "GET")
	if err != nil {
		return
	}

	// for development purpose
	if saveRawJsonByte {
		SaveRawJsonByte("saved-", b)
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

// GetSavedCollection returns your saved collections.
func (m *IGApiManager) GetSavedCollection(id string) (items []IGItem, err error) {
	url := strings.Replace(urlSavedCollection, "{{COLLECTIONID}}", id, 1)
	b, err := m.getHTTPResponse(url, "GET")
	if err != nil {
		return
	}

	// for development purpose
	if saveRawJsonByte {
		SaveRawJsonByte(id+"-collection-", b)
	}

	spp := savedPostsResp{}
	err = json.Unmarshal(b, &spp)
	if err != nil {
		return
	}
	for _, item := range spp.Items {
		items = append(items, item.Item)
	}

	return
}

// GetSavedCollectionList returns your list of saved collections.
func (m *IGApiManager) GetSavedCollectionList() (c []Collection, err error) {
	b, err := m.getHTTPResponse(urlSavedCollectionList, "GET")
	if err != nil {
		return
	}

	// for development purpose
	if saveRawJsonByte {
		SaveRawJsonByte("collections-list-", b)
	}

	cl := collectionsList{}
	err = json.Unmarshal(b, &cl)
	if err != nil {
		return
	}
	c = cl.Collections
	return
}
