package instago

import (
	"encoding/json"
	"errors"
	"strings"
)

const urlReelsMedia = `https://i.instagram.com/api/v1/feed/reels_media/`

type highlightMedia struct {
	Reels struct {
		ReelsMedia IGStoryHighlightsTray `json:"reels_media"`
	} `json:"reels"`
	Status string `json:"status"`
}

// GetHighlightsReelsMedia returns the content of the highlight tray, which
// contains metadata of story highlights of a specific title.
func (m *IGApiManager) GetHighlightsReelsMedia(id string) (tray IGStoryHighlightsTray, err error) {
	url := urlReelsMedia + "?user_ids=" + id
	b, err := m.getHTTPResponse(url, "GET")
	if err != nil {
		return
	}

	// for development purpose
	if saveRawJsonByte {
		SaveRawJsonByte("reels_media-", b)
	}

	// The name of json field is the id of the highlight tray, which is only
	// known in run-time, not compile-time. So we need to replace the id of
	// the highlight tray with *reels_media*, which can be decoded by Go
	// standard encoding/json package.
	bb := []byte(strings.Replace(string(b), id, "reels_media", 1))

	h := highlightMedia{}
	err = json.Unmarshal(bb, &h)
	if err != nil {
		return
	}

	// Check validity
	if h.Reels.ReelsMedia.Id != id {
		err = errors.New("Returned highlight tray seems weird")
		return
	}
	tray = h.Reels.ReelsMedia
	return
}

type ReelsMedia struct {
	Reels  map[string]IGReelsMediaTray `json:"reels"`
	Status string                      `json:"status"`
}

// Used to decode JSON returned by Instagram story API.
type IGReelsMediaTray struct {
	Id              int64   `json:"id"`
	LatestReelMedia int64   `json:"latest_reel_media"`
	Seen            float64 `json:"seen"`
	CanReply        bool    `json:"can_reply"`
	CanReshare      bool    `json:"can_reshare"`
	ReelType        string  `json:"reel_type"`

	CoverMedia struct {
		CroppedImageVersion struct {
			Width  int64  `json:"width"`
			Height int64  `json:"height"`
			Url    string `json:"url"`
		} `json:"cropped_image_version"`
		MediaId  string    `json:"media_id"`
		CropRect []float64 `json:"crop_rect"`
	} `json:"cover_media"`

	User  IGUser   `json:"user"`
	Items []IGItem `json:"items"`

	RankedPosition     int64  `json:"ranked_position"`
	Title              string `json:"title"`
	SeenRankedPosition int64  `json:"seen_ranked_position"`
	PrefetchCount      int64  `json:"prefetch_count"`
}

func (m *IGApiManager) GetMultipleReelsMedia(userids []string) (trays []IGReelsMediaTray, err error) {
	if len(userids) == 0 {
		err = errors.New("length of user ids is 0")
		return
	}

	query := ""
	for _, userid := range userids {
		query += "user_ids=" + userid + "&"
	}
	query = strings.TrimSuffix(query, "&")
	url := urlReelsMedia + "?" + query
	//println(url)

	b, err := m.getHTTPResponse(url, "GET")
	if err != nil {
		return
	}

	// for development purpose
	if saveRawJsonByte {
		SaveRawJsonByte("reels_media-", b)
	}

	// The name of json field is the id of the highlight tray, which is only
	// known in run-time, not compile-time. So we need to replace the id of
	// the highlight tray with *reels_media*, which can be decoded by Go
	// standard encoding/json package.
	/*
		bb := b
		for _, userid := range userids {
			str := `"` + userid + `": {"id":`
			//str2 := `"reel_media": {"id":`
			str2 := `{"id":`
			bb = []byte(strings.Replace(string(bb), str, str2, 1))
		}
		bb = []byte(strings.Replace(string(bb), `}, "status": "ok"}`, `], "status": "ok"}`, 1))
		bb = []byte(strings.Replace(string(bb), `{"reels": {`, `{"reels": [`, 1))
	*/

	rsm := ReelsMedia{}
	//err = json.Unmarshal(bb, &rsm)
	err = json.Unmarshal(b, &rsm)
	if err != nil {
		return
	}
	//trays = rsm.Reels
	for _, tray := range rsm.Reels {
		trays = append(trays, tray)
	}

	return
}
