package instago

import (
	"encoding/json"
	"errors"
	"strings"
)

type WebStoryInfo struct {
	User struct {
		Id            string `json:"id"`
		ProfilePicUrl string `json:"profile_pic_url"`
		Username      string `json:"username"`
	} `json:"user"`
	Highlight struct {
		Id    int64  `json:"id"`
		Title string `json:"title"`
	} `json:"highlight"`
}

type WebStoryQueryResponse struct {
	Data struct {
		ReelsMedia []IGReelMedia `json:"reels_media"`
	} `json:"data"`
	Status string `json:"status"`
}

type IGReelMediaUser struct {
	Typename          string `json:"__typename"`
	Id                string `json:"id"`
	ProfilePicUrl     string `json:"profile_pic_url"`
	Username          string `json:"username"`
	FollowedByViewer  bool   `json:"followed_by_viewer"`
	RequestedByViewer bool   `json:"requested_by_viewer"`
}

// IGReelMedia represent story info of a user. web version of type IGReelTray
// This struct is returned from web GraphQL query.
type IGReelMedia struct {
	Typename        string            `json:"__typename"`
	Id              string            `json:"id"`
	LatestReelMedia int64             `json:"latest_reel_media"`
	CanReply        bool              `json:"can_reply"`
	Owner           IGReelMediaUser   `json:"owner"`
	CanReshare      bool              `json:"can_reshare"`
	ExpiringAt      int64             `json:"expiring_at"`
	HasBestiesMedia bool              `json:"has_besties_media"`
	HasPrideMedia   bool              `json:"has_pride_media"`
	Seen            int64             `json:"seen"`
	User            IGReelMediaUser   `json:"user"`
	Items           []IGReelMediaItem `json:"items"`
}

func (m *IGApiManager) GetInfoFromWebStoryUrl(url string) (user WebStoryInfo, err error) {
	if !IsWebStoryUrl(url) {
		err = errors.New(url + " is not a valid web story url")
		return
	}

	jsonurl := url + "?__a=1"
	b, err := m.getHTTPResponse(jsonurl, "GET")
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &user)
	return
}

func (m *IGApiManager) GetIdFromWebStoryUrl(url string) (id string, err error) {
	user, err := m.GetInfoFromWebStoryUrl(url)
	id = user.User.Id
	return
}

func (m *IGApiManager) GetWebGraphqlStoriesJson(reelIds []string, storyQueryHash string) (b []byte, err error) {
	if len(reelIds) == 0 {
		err = errors.New("no reel_ids is given")
		return
	}
	url := `https://www.instagram.com/graphql/query/?query_hash={{QueryHash}}&variables={"reel_ids":[{{ReelIds}}],"tag_names":[],"location_ids":[],"highlight_reel_ids":[],"precomposed_overlay":false,"show_story_viewer_list":true,"story_viewer_fetch_count":50,"story_viewer_cursor":"","stories_video_dash_manifest":false}`
	url = strings.Replace(url, "{{QueryHash}}", storyQueryHash, 1)

	rids := ""
	for _, reelId := range reelIds {
		rids += `"` + reelId + `",`
	}
	rids = strings.TrimSuffix(rids, ",")
	url = strings.Replace(url, "{{ReelIds}}", rids, 1)

	b, err = m.getHTTPResponse(url, "GET")
	return
}

func (m *IGApiManager) GetUserStoryByWebGraphql(id, storyQueryHash string) (rm IGReelMedia, err error) {
	b, err := m.GetWebGraphqlStoriesJson([]string{id}, storyQueryHash)
	if err != nil {
		return
	}

	rsp := WebStoryQueryResponse{}
	err = json.Unmarshal(b, &rsp)
	if err != nil {
		return
	}
	if rsp.Status != "ok" {
		err = errors.New("response status is not ok")
		return
	}
	if len(rsp.Data.ReelsMedia) != 1 {
		err = errors.New("response reels_media length is not 1")
		return
	}
	rm = rsp.Data.ReelsMedia[0]
	return
}
