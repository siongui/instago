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

func GetInfoFromWebStoryUrl(url string) (user WebStoryInfo, err error) {
	if !IsWebStoryUrl(url) {
		err = errors.New(url + " is not a valid web story url")
		return
	}

	jsonurl := url + "?__a=1"
	b, err := GetHTTPResponseNoLogin(jsonurl)
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &user)
	return
}

func GetIdFromWebStoryUrl(url string) (id string, err error) {
	user, err := GetInfoFromWebStoryUrl(url)
	id = user.User.Id
	return
}

func GetWebGraphqlStoriesJson(reelIds []string, storyQueryHash string) (b []byte, err error) {
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

	b, err = GetHTTPResponseNoLogin(url)
	return
}
