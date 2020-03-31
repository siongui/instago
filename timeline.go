package instago

import (
	"encoding/json"
)

const urlTimeline = `https://i.instagram.com/api/v1/feed/timeline/`

type IGTimeline struct {
	Items               []IGItem `json:"items"`
	NumResults          int64    `json:"num_results"`
	MoreAvailable       bool     `json:"more_available"`
	AutoLoadMoreEnabled bool     `json:"auto_load_more_enabled"`
	IsDirectV2Enabled   bool     `json:"is_direct_v2_enabled"`
	NextMaxId           string   `json:"next_max_id"`
	Status              string   `json:"status"`
}

// Get only one page of timeline
func (m *IGApiManager) GetTimeline() (tl IGTimeline, err error) {
	b, err := m.getHTTPResponse(urlTimeline, "POST")
	if err != nil {
		return
	}

	//println(string(b))
	err = json.Unmarshal(b, &tl)
	return
}

// Get N pages of timeline
func (m *IGApiManager) GetTimelineUntilPageN(pageN int) (items []IGItem, err error) {
	tl := IGTimeline{}
	tl.MoreAvailable = true
	pageCount := 0

	for tl.MoreAvailable && pageCount < pageN {
		url := ""
		if pageCount == 0 {
			url = urlTimeline
		} else {
			url = urlTimeline + "?max_id=" + tl.NextMaxId
		}
		b, err := m.getHTTPResponse(url, "POST")
		if err != nil {
			return items, err
		}
		pageCount++
		err = json.Unmarshal(b, &tl)
		if err != nil {
			return items, err
		}
		items = append(items, tl.Items...)
		tl = IGTimeline{
			MoreAvailable: tl.MoreAvailable,
			NextMaxId:     tl.NextMaxId,
		}
	}

	return
}
