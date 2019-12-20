package instago

// This file implements Instagram discover top live API.

import (
	"encoding/json"
)

const urlToplive = `https://i.instagram.com/api/v1/discover/top_live/`

// Decode JSON data returned from Instagram top live API
type IGTopliveResp struct {
	Broadcasts []struct {
		Id                   int64   `json:"id"`
		RtmpPlaybackUrl      string  `json:"rtmp_playback_url"`
		DashPlaybackUrl      string  `json:"dash_playback_url"`
		DashAbrPlaybackUrl   string  `json:"dash_abr_playback_url"`
		BroadcastStatus      string  `json:"broadcast_status"`
		ViewerCount          float64 `json:"viewer_count"`
		InternalOnly         bool    `json:"internal_only"`
		CoverFrameUrl        string  `json:"cover_frame_url"`
		BroadcastOwner       IGUser  `json:"broadcast_owner"`
		PublishedTime        int64   `json:"published_time"`
		MediaId              string  `json:"media_id"`
		BroadcastMessage     string  `json:"broadcast_message"`
		OrganicTrackingToken string  `json:"organic_tracking_token"`
	} `json:"broadcasts"`
	MoreAvailable       bool   `json:"more_available"`
	AutoLoadMoreEnabled bool   `json:"auto_load_more_enabled"`
	NextMaxId           int64  `json:"next_max_id"`
	Status              string `json:"status"`
}

// Given a string, return the users that best matches the string. This is
// actually *topsearch* on Instagram web.
func (m *IGApiManager) Toplive() (tr IGTopliveResp, err error) {
	b, err := m.getHTTPResponse(urlToplive, "GET")
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &tr)
	return
}
