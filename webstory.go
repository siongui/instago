package instago

import (
	"encoding/json"
	"errors"
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
