package instago

import (
	"regexp"
	"strconv"
)

// Live videos that users share to their stories
type IGPostLive struct {
	PostLiveItems []IGPostLiveItem `json:"post_live_items"`
}

type IGPostLiveItem struct {
	Pk                  string        `json:"pk"`
	User                IGUser        `json:"user"`
	Broadcasts          []IGBroadcast `json:"broadcasts"`
	LastSeenBroadcastTs float64       `json:"last_seen_broadcast_ts"`
	RankedPosition      int64         `json:"ranked_position"`
	SeenRankedPosition  int64         `json:"seen_ranked_position"`
	Muted               bool          `json:"muted"`
	CanReply            bool          `json:"can_reply"`
	CanReshare          bool          `json:"can_reshare"`
}

func (i *IGPostLiveItem) GetUsername() string {
	return i.User.Username
}

func (i *IGPostLiveItem) GetUserId() string {
	return strconv.FormatInt(i.User.Pk, 10)
}

func (i *IGPostLiveItem) GetBroadcasts() []IGBroadcast {
	return i.Broadcasts
}

type IGBroadcast struct {
	Id                   int64  `json:"id"`
	BroadcastStatus      string `json:"broadcast_status"`
	DashManifest         string `json:"dash_manifest"`
	ExpireAt             int64  `json:"expire_at"`
	EncodingTag          string `json:"encoding_tag"`
	InternalOnly         bool   `json:"internal_only"`
	NumberOfQualities    int64  `json:"number_of_qualities"`
	CoverFrameUrl        string `json:"cover_frame_url"`
	BroadcastOwner       IGUser `json:"broadcast_owner"`
	PublishedTime        int64  `json:"published_time"`
	MediaId              string `json:"media_id"`
	BroadcastMessage     string `json:"broadcast_message"`
	OrganicTrackingToken string `json:"organic_tracking_token"`
}

func (b *IGBroadcast) GetDashManifest() string {
	return b.DashManifest
}

func (b *IGBroadcast) GetBaseUrls() (urls []string, err error) {
	reBaseUrls, err := regexp.Compile(`<BaseURL>(.+?)<\/BaseURL>`)
	if err != nil {
		return
	}

	matches := reBaseUrls.FindAllStringSubmatch(b.GetDashManifest(), -1)
	for _, match := range matches {
		urls = append(urls, match[1])
	}
	return
}

func (b *IGBroadcast) GetPublishedTime() int64 {
	return b.PublishedTime
}
