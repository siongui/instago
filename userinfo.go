package instago

import (
	"encoding/json"
	"errors"
	"regexp"
	"strings"
)

// no need to login or cookie to access this URL. But if login to Instagram,
// private account will return private data if you are allowed to view the
// private account.
const urlUserInfo = `https://www.instagram.com/{{USERNAME}}/?__a=1`
const urlGraphql = `https://instagram.com/graphql/query/?query_id=17888483320059182&variables=`

// used to decode the JSON data
// https://www.instagram.com/{{USERNAME}}/?__a=1
type rawUserResp struct {
	LoggingPageId         string `json:"logging_page_id"`
	ShowSuggestedProfiles bool   `json:"show_suggested_profiles"`
	GraphQL               struct {
		User UserInfo `json:"user"`
	} `json:"graphql"`
}

// used to decode the JSON data
// https://www.instagram.com/{{USERNAME}}/
type SharedData struct {
	EntryData struct {
		ProfilePage []struct {
			GraphQL struct {
				User UserInfo `json:"user"`
			} `json:"graphql"`
		} `json:"ProfilePage"`
	} `json:"entry_data"`

	RhxGis string `json:"rhx_gis"`
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

func (u UserInfo) GetUserId() string {
	return u.Id
}

func (u UserInfo) GetUsername() string {
	return u.Username
}

func (u UserInfo) IsPublic() bool {
	return !u.IsPrivate
}

type dataUserMedia struct {
	Data struct {
		User UserInfo `json:"user"`
	} `json:"data"`
}

func getJsonBytes(b []byte) []byte {
	pattern := regexp.MustCompile(`<script type="text\/javascript">window\._sharedData = (.*?);<\/script>`)
	m := string(pattern.Find(b))
	m1 := strings.TrimPrefix(m, `<script type="text/javascript">window._sharedData = `)
	return []byte(strings.TrimSuffix(m1, `;</script>`))
}

func checkSharedData(sd SharedData) (err error) {
	if len(sd.EntryData.ProfilePage) == 0 {
		err = errors.New("Looks like decoded JSON data is not correct")
	}
	return
}

// Given the HTML source code of the user profile page without logged in, return
// query_hash for Instagram GraphQL API.
func GetQueryHashNoLogin(b []byte) (qh string, err error) {
	// find JavaScript file which contains the query hash
	patternJs := regexp.MustCompile(`\/static\/bundles\/metro\/ProfilePageContainer\.js\/[a-zA-Z0-9]+?\.js`)
	jsPath := string(patternJs.Find(b))
	jsUrl := "https://www.instagram.com" + jsPath
	bJs, err := GetHTTPResponseNoLogin(jsUrl)
	if err != nil {
		return
	}

	patternQh := regexp.MustCompile(`t\.profilePosts\.byUserId\.get.+?queryId:"([a-zA-Z0-9]+)",`)
	qhtmp := string(patternQh.Find(bJs))
	//qhtmp = strings.Split(qhtmp, `queryId:"`)[1]
	//qh = strings.TrimSuffix(qhtmp, `",`)
	qhtmp2 := strings.Split(qhtmp, `queryId:"`)
	if len(qhtmp2) != 2 {
		err = errors.New("fail to get queryId")
		return
	}
	qhtmp3 := qhtmp2[1]
	qh = strings.TrimSuffix(qhtmp3, `",`)
	return
}

// Given username, get:
//
//   1. sharedData embedded in the HTML of user profile page.
//   2. query_hash (for get all codes of posts without login)
func GetSharedDataQueryHashNoLogin(username string) (sd SharedData, qh string, err error) {
	url := "https://www.instagram.com/" + username + "/"
	b, err := GetHTTPResponseNoLogin(url)
	if err != nil {
		return
	}

	err = json.Unmarshal(getJsonBytes(b), &sd)
	if err != nil {
		return
	}

	err = checkSharedData(sd)
	if err != nil {
		return
	}

	qh, err = GetQueryHashNoLogin(b)
	return
}

// GetSharedDataNoLogin returns sharedData embedded in the HTML source code of
// the profile page of given user name.
func GetSharedDataNoLogin(username string) (sd SharedData, err error) {
	url := "https://www.instagram.com/" + username + "/"
	b, err := GetHTTPResponseNoLogin(url)
	if err != nil {
		return
	}

	err = json.Unmarshal(getJsonBytes(b), &sd)
	if err != nil {
		return
	}

	err = checkSharedData(sd)
	return
}

// GetUserInfoNoLogin returns information of the given user name without login
// status.
func GetUserInfoNoLogin(username string) (ui UserInfo, err error) {
	sd, err := GetSharedDataNoLogin(username)
	if err != nil {
		return
	}

	ui = sd.EntryData.ProfilePage[0].GraphQL.User
	return
}

// GetUserInfo returns information of the given user name with logged in status.
func (m *IGApiManager) GetUserInfo(username string) (ui UserInfo, err error) {
	//url := strings.Replace(urlUserInfo, "{{USERNAME}}", username, 1)
	url := "https://www.instagram.com/" + username + "/"
	b, err := m.getHTTPResponse(url, "GET")
	if err != nil {
		return
	}

	//r := rawUserResp{}
	r := SharedData{}
	if err = json.Unmarshal(getJsonBytes(b), &r); err != nil {
		return
	}

	if err = checkSharedData(r); err != nil {
		return
	}

	ui = r.EntryData.ProfilePage[0].GraphQL.User
	return
}

// GetRecentPostCodeNoLogin returns codes of recent posts (usually 12 posts) of
// the given user name without login status.
func GetRecentPostCodeNoLogin(username string) (codes []string, err error) {
	ui, err := GetUserInfoNoLogin(username)
	if err != nil {
		return
	}

	for _, node := range ui.EdgeOwnerToTimelineMedia.Edges {
		codes = append(codes, node.Node.Shortcode)
	}
	return
}

// GetRecentPostCode returns codes of recent posts (usually 12 posts) of the
// given user name with logged in status.
func (m *IGApiManager) GetRecentPostCode(username string) (codes []string, err error) {
	ui, err := m.GetUserInfo(username)
	if err != nil {
		return
	}

	for _, node := range ui.EdgeOwnerToTimelineMedia.Edges {
		codes = append(codes, node.Node.Shortcode)
	}
	return
}

// GetUserId returns id of the given user name.
func GetUserId(username string) (id string, err error) {
	ui, err := GetUserInfoNoLogin(username)
	if err != nil {
		return
	}
	id = ui.Id
	return
}

// GetUserProfilePicUrlHd returns the URL of profile hd pic of given user name.
func GetUserProfilePicUrlHd(username string) (url string, err error) {
	ui, err := GetUserInfoNoLogin(username)
	if err != nil {
		return
	}
	url = ui.ProfilePicUrlHd
	return
}

// GetRecentPostMediaNoLogin returns IGMedia struct of recent posts (usually 12
// posts if not private) of the given user name without login status.
func GetRecentPostMediaNoLogin(username string) (medias []IGMedia, err error) {
	ui, err := GetUserInfoNoLogin(username)
	if err != nil {
		return
	}

	for _, node := range ui.EdgeOwnerToTimelineMedia.Edges {
		medias = append(medias, node.Node)
	}
	return
}

// GetRecentPostMedia returns IGMedia struct of recent posts (usually 12 posts)
// of the given user name with logged in status.
func (m *IGApiManager) GetRecentPostMedia(username string) (medias []IGMedia, err error) {
	ui, err := m.GetUserInfo(username)
	if err != nil {
		return
	}

	for _, node := range ui.EdgeOwnerToTimelineMedia.Edges {
		medias = append(medias, node.Node)
	}
	return
}
