package instago

// Get story highlights of a specific user

import (
	"encoding/json"
	"strings"
)

const urlUserHighlightStories = `https://i.instagram.com/api/v1/highlights/{{USERID}}/highlights_tray/`

// Used to decode JSON returned by Instagram story API.
type rawHighlightsTray struct {
	Trays          []IGStoryHighlightsTray `json:"tray"`
	ShowEmptyState bool                    `json:"show_empty_state"`
	Status         string                  `json:"status"`
}

// Used to decode JSON returned by Instagram story API.
type IGStoryHighlightsTray struct {
	Id              string  `json:"id"`
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

func (t *IGStoryHighlightsTray) GetTitle() string {
	return t.Title
}

func (t *IGStoryHighlightsTray) GetUsername() string {
	return t.User.Username
}

func (t *IGStoryHighlightsTray) GetItems() []IGItem {
	return t.Items
}

// Return story highlight trays of a specific user. Sometimes items in a
// highlight tray are empty. Call *IGApiManager.GetHighlightsReelsMedia to get
// items of the tray. See *IGApiManager.GetAllStoryHighlights
func (m *IGApiManager) GetUserStoryHighlights(id string) (trays []IGStoryHighlightsTray, err error) {
	url := strings.Replace(urlUserHighlightStories, "{{USERID}}", id, 1)
	b, err := m.getHTTPResponse(url, "GET")
	if err != nil {
		return
	}

	// for development purpose
	if saveRawJsonByte {
		SaveRawJsonByte(id+"-highlights_tray-", b)
	}

	t := rawHighlightsTray{}
	err = json.Unmarshal(b, &t)
	if err != nil {
		return
	}
	trays = t.Trays
	return
}
