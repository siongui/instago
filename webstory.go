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

type IGReelMediaItem struct {
	Audience              string `json:"audience"`
	EdgeStoryMediaViewers struct {
		Count    int64 `json:"count"`
		PageInfo struct {
			HasNextPage bool `json:"has_next_page"`
			//end_cursor
		} `json:"page_info"`
		//edges
	} `json:"edge_story_media_viewers"`
	Typename   string `json:"__typename"`
	Id         string `json:"id"`
	Dimensions struct {
		Height int64 `json:"height"`
		Width  int64 `json:"width"`
	} `json:"dimensions"`
	DisplayResources []struct {
		Src          string `json:"src"`
		ConfigWidth  int64  `json:"config_width"`
		ConfigHeight int64  `json:"config_height"`
	} `json:"display_resources"`
	DisplayUrl   string `json:"display_url"`
	MediaPreview string `json:"media_preview"`
	//gating_info
	//fact_check_overall_rating
	//fact_check_information
	//media_overlay_info
	//sensitivity_friction_info
	TakenAtTimestamp    int64 `json:"taken_at_timestamp"`
	ExpiringAtTimestamp int64 `json:"expiring_at_timestamp"`
	//story_cta_url
	//story_view_count
	IsVideo         bool            `json:"is_video"`
	Owner           IGReelMediaUser `json:"owner"`
	TrackingToken   string          `json:"tracking_token"`
	TappableObjects []struct {
		Typename string  `json:"__typename"`
		X        float64 `json:"x"`
		Y        float64 `json:"y"`
		Width    float64 `json:"width"`
		Height   float64 `json:"height"`
		Rotation float64 `json:"rotation"`
		//custom_title
		//attribution
		TappableType string `json:"tappable_type"`
		Username     string `json:"username"`
		FullName     string `json:"full_name"`
		IsPrivate    bool   `json:"is_private"`
	} `json:"tappable_objects"`
	//story_app_attribution
	//edge_media_to_sponsor_user
	//muting_info
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
