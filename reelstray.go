package instago

// reels_tray contains the following information:
//
//   1. unread stories
//   2. users with unexpired stories
//   3. post-live video (if User-Agent is correctly set)
//
// The response will return all users with unexpired stories,
// but only stories of users of unread stories will be returned.

import (
	"encoding/json"
	"regexp"
)

const urlReelsTray = `https://i.instagram.com/api/v1/feed/reels_tray/`

// Used to decode JSON returned by Instagram reels tray feed API.
type IGReelsTray struct {
	Trays             []IGReelTray  `json:"tray"`
	Broadcasts        []IGBroadcast `json:"broadcasts"`
	PostLive          IGPostLive    `json:"post_live"`
	StoryRankingToken string        `json:"story_ranking_token"`

	// "broadcasts"

	StickerVersion       int64  `json:"sticker_version"`
	FaceFilterNuxVersion int64  `json:"face_filter_nux_version"`
	HasNewNuxStory       bool   `json:"has_new_nux_story"`
	Status               string `json:"status"`
}

type IGReelTray struct {
	Id                 int64   `json:"id"`
	LatestReelMedia    int64   `json:"latest_reel_media"`
	ExpiringAt         float64 `json:"expiring_at"`
	Seen               float64 `json:"seen"`
	CanReply           bool    `json:"can_reply"`
	CanReshare         bool    `json:"can_reshare"`
	ReelType           string  `json:"reel_type"`
	User               IGUser  `json:"user"`
	RankedPosition     int64   `json:"ranked_position"`
	SeenRankedPosition int64   `json:"seen_ranked_position"`
	Muted              bool    `json:"muted"`
	PrefetchCount      int64   `json:"prefetch_count"`

	// close friend
	HasBestiesMedia bool `json:"has_besties_media"`

	Items []IGItem `json:"items"`
}

func (t *IGReelTray) GetUsername() string {
	return t.User.Username
}

func (t *IGReelTray) GetItems() []IGItem {
	return t.Items
}

func removePride(b []byte) []byte {
	pattern := regexp.MustCompile(`{"id": "election:pride:.+?"hide_from_feed_unit": true},`)
	return []byte(pattern.ReplaceAllString(string(b), ""))
}

func removeLunar(b []byte) []byte {
	pattern := regexp.MustCompile(`{"id":"election:lunar_new_year:.+?"hide_from_feed_unit":true},`)
	return []byte(pattern.ReplaceAllString(string(b), ""))
}

func (m *IGApiManager) GetReelsTray() (r IGReelsTray, err error) {
	b, err := m.getHTTPResponse(urlReelsTray, "GET")
	if err != nil {
		return
	}

	// for development purpose
	if saveRawJsonByte {
		SaveRawJsonByte("reels_tray-", b)
	}

	err = json.Unmarshal(removeLunar(b), &r)
	return
}
