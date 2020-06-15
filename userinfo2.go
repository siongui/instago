package instago

import (
	"encoding/json"
	"strconv"
	"strings"
)

const urlUsers = `https://i.instagram.com/api/v1/users/{{USERID}}/info/`

type UserInfoEndPoint struct {
	Pk                         int64  `json:"pk"`
	Username                   string `json:"username"`
	FullName                   string `json:"full_name"`
	IsPrivate                  bool   `json:"is_private"`
	ProfilePicUrl              string `json:"profile_pic_url"`
	ProfilePicId               string `json:"profile_pic_id"`
	IsVerified                 bool   `json:"is_verified"`
	HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
	MediaCount                 int64  `json:"media_count"`
	GeoMediaCount              int64  `json:"geo_media_count"`
	FollowerCount              int64  `json:"follower_count"`
	FollowingCount             int64  `json:"following_count"`
	FollowingTagCount          int64  `json"following_tag_count"`
	Biography                  string `json:"biography"`
	BiographyWithEntities      struct {
		RawText string `json:"raw_text"`
		//Entities	???	`json:"entities"`
	} `json:"biography_with_entities"`
	ExternalUrl             string `json:"external_url"`
	TotalIgtvVideos         int64  `json:"total_igtv_videos"`
	TotalClipsCount         int64  `json:"total_clips_count"`
	TotalArEffects          int64  `json:"total_ar_effects"`
	UsertagsCount           int64  `json:"usertags_count"`
	IsFavorite              bool   `json:"is_favorite"`
	IsFavoriteForStories    bool   `json:"is_favorite_for_stories"`
	IsFavoriteForIgtv       bool   `json:"is_favorite_for_igtv"`
	IsFavoriteForHighlights bool   `json:"is_favorite_for_highlights"`
	LiveSubscriptionStatus  string `json:"live_subscription_status"`
	IsInterestAccount       bool   `json:"is_interest_account"`
	HasChaining             bool   `json:"has_chaining"`
	HdProfilePicVersions    []struct {
		Width  int64  `json:"width"`
		Height int64  `json:"height"`
		Url    string `json:"url"`
	} `json:"hd_profile_pic_versions"`
}

func (u UserInfoEndPoint) GetUserId() string {
	return strconv.FormatInt(u.Pk, 10)
}

func (u UserInfoEndPoint) GetUsername() string {
	return u.Username
}

func (u UserInfoEndPoint) IsPublic() bool {
	return !u.IsPrivate
}

type rawUserinfoResp struct {
	User UserInfoEndPoint `json:"user"`
}

// FIXME: do not return IGUser
func (m *IGApiManager) GetUserInfoEndPoint(userid string) (user UserInfoEndPoint, err error) {
	url := strings.Replace(urlUsers, "{{USERID}}", userid, 1)
	b, err := m.getHTTPResponse(url, "GET")
	if err != nil {
		return
	}

	// for development purpose
	if saveRawJsonByte {
		SaveRawJsonByte(userid+"-info-", b)
	}

	r := rawUserinfoResp{}
	err = json.Unmarshal(b, &r)
	if err == nil {
		user = r.User
	}
	return
}
