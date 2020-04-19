package instago

// This file decode JSON data returned by Instagram media API endpoint

import (
	"encoding/json"
	"errors"
	"strings"
)

const urlMedia = `https://i.instagram.com/api/v1/media/{{ID}}/info/`

type mediaPostResp struct {
	Items               []IGItem `json:"items"`
	NumResults          int64    `json:"num_results"`
	MoreAvailable       bool     `json:"more_available"`
	AutoLoadMoreEnabled bool     `json:"auto_load_more_enabled"`
	Status              string   `json:"status"`
}

// GetMediaInfo returns information of post via API endpoint.
func (m *IGApiManager) GetMediaInfo(id string) (item IGItem, err error) {
	url := strings.Replace(urlMedia, "{{ID}}", id, 1)
	b, err := m.getHTTPResponse(url, "GET")
	if err != nil {
		return
	}

	// for development purpose
	if saveRawJsonByte {
		SaveRawJsonByte("media-"+id+"-info-", b)
	}

	mpp := mediaPostResp{}
	err = json.Unmarshal(b, &mpp)
	if err != nil {
		return
	}

	if len(mpp.Items) != 1 {
		err = errors.New("number of item != 1")
		return
	}
	item = mpp.Items[0]
	return
}
