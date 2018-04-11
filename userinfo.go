package instago

import (
	"encoding/json"
	"strings"
)

// no need to login or cookie to access this URL. But if login to Instagram,
// private account will return private data if you are allowed to view the
// private account.
const urlUserInfo = `https://www.instagram.com/{{USERNAME}}/?__a=1`
const urlGraphql = `https://instagram.com/graphql/query/?query_id=17888483320059182&variables=`

// used to decode the JSON data
type rawUserResp struct {
	LoggingPageId         string `json:"logging_page_id"`
	ShowSuggestedProfiles bool   `json:"show_suggested_profiles"`
	GraphQL               struct {
		User UserInfo `json:"user"`
	} `json:"graphql"`
}

type UserInfo struct {
	Biography              string `json:"biography"`
	BlockedByViewer        bool   `json:"blocked_by_viewer"`
	CountryBlock           bool   `json:"country_block"`
	ExternalUrl            string `json:"external_url"`
	ExternalUrlLinkshimmed string `json:"external_url_linkshimmed"`
	EdgeFollowedBy         struct {
		Count int64 `json:"count"`
	} `json:"edge_followed_by"`
	FollowedByViewer bool `json:"followed_by_viewer"`
	EdgeFollow       struct {
		Count int64 `json:"count"`
	} `json:"edge_followe"`
	FollowsViewer      bool   `json:"follows_viewer"`
	FullName           string `json:"full_name"`
	HasBlockedViewer   bool   `json:"has_blocked_viewer"`
	HasRequestedViewer bool   `json:"has_requested_viewer"`
	Id                 string `json:"id"`
	IsPrivate          bool   `json:"is_private"`
	IsVerified         bool   `json:"is_verified"`
	MutualFollowers    struct {
		AdditionalCount int64    `json:"additional_count"`
		Usernames       []string `json:"usernames"`
	} `json:"mutual_followers"`
	ProfilePicUrl            string `json:"profile_pic_url"`
	ProfilePicUrlHd          string `json:"profile_pic_url_hd"`
	RequestedByViewer        bool   `json:"requested_by_viewer"`
	Username                 string `json:"username"`
	ConnectedFbPage          string `json:"connected_fb_page"`
	EdgeOwnerToTimelineMedia struct {
		Count    int64 `json:"count"`
		PageInfo struct {
			HasNextPage bool   `json:"has_next_page"`
			EndCursor   string `json:"end_cursor"`
		} `json:"page_info"`
		Edges []struct {
			Node IGMedia `json:"node"`
		} `json:"edges"`
	} `json:"edge_owner_to_timeline_media"`
}

type dataUserMedia struct {
	Data struct {
		User UserInfo `json:"user"`
	} `json:"data"`
}

// Given user name, return information of the user name without login.
func GetUserInfoNoLogin(username string) (ui UserInfo, err error) {
	url := strings.Replace(urlUserInfo, "{{USERNAME}}", username, 1)
	b, err := getHTTPResponseNoLogin(url)
	if err != nil {
		return
	}

	r := rawUserResp{}
	if err = json.Unmarshal(b, &r); err != nil {
		return
	}
	ui = r.GraphQL.User
	return
}

// Given user name, return information of the user name.
func (m *IGApiManager) GetUserInfo(username string) (ui UserInfo, err error) {
	url := strings.Replace(urlUserInfo, "{{USERNAME}}", username, 1)
	b, err := getHTTPResponse(url, m.dsUserId, m.sessionid, m.csrftoken)
	if err != nil {
		return
	}

	r := rawUserResp{}
	if err = json.Unmarshal(b, &r); err != nil {
		return
	}
	ui = r.GraphQL.User
	return
}

// Given user name, return codes of all posts of the user.
// TODO: add sleep at the end of forloop. If the number of posts is over 2400,
// Instagram API will return http response code 429 (Too Many Requests)
func (m *IGApiManager) GetAllPostCode(username string) (codes []string, err error) {
	ui, err := m.GetUserInfo(username)
	if err != nil {
		return
	}

	for _, node := range ui.EdgeOwnerToTimelineMedia.Edges {
		codes = append(codes, node.Node.Shortcode)
	}

	hasNextPage := ui.EdgeOwnerToTimelineMedia.PageInfo.HasNextPage
	vartmpl := strings.Replace(`{"id":"<ID>","first":300,"after":"<ENDCURSOR>"}`, "<ID>", ui.Id, 1)
	variables := strings.Replace(vartmpl, "<ENDCURSOR>", ui.EdgeOwnerToTimelineMedia.PageInfo.EndCursor, 1)

	for hasNextPage == true {
		url := urlGraphql + variables

		b, err := getHTTPResponse(url, m.dsUserId, m.sessionid, m.csrftoken)
		if err != nil {
			return codes, err
		}

		d := dataUserMedia{}
		if err = json.Unmarshal(b, &d); err != nil {
			return codes, err
		}

		for _, node := range d.Data.User.EdgeOwnerToTimelineMedia.Edges {
			codes = append(codes, node.Node.Shortcode)
		}
		hasNextPage = d.Data.User.EdgeOwnerToTimelineMedia.PageInfo.HasNextPage
		variables = strings.Replace(vartmpl, "<ENDCURSOR>", d.Data.User.EdgeOwnerToTimelineMedia.PageInfo.EndCursor, 1)

		printPostCount(len(codes), url)
	}
	return
}

// Given user name, return id of the user name.
func GetUserId(username string) (id string, err error) {
	ui, err := GetUserInfoNoLogin(username)
	if err != nil {
		return
	}
	id = ui.Id
	return
}
