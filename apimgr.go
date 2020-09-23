// Package instago helps you access the private API of Instagram. For examle,
// get URLs of all posts of a specific Instagram user, media (photos and videos)
// links of posts, stories of a Instagram user, your following and followers.
package instago

import (
	"encoding/json"
	"io/ioutil"
)

type IGApiManager struct {
	cookies map[string]string
	headers map[string]string
}

// Google Search: useragent ios instagram 134.00
// https://developers.whatismybrowser.com/useragents/parse/31984417-instagram-ios-iphone-11-pro-max-webkit
var appUserAgent = "Mozilla/5.0 (iPhone; CPU iPhone OS 13_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 Instagram 134.0.0.25.116 (iPhone12,5; iOS 13_3_1; en_US; en-US; scale=3.00; 1242x2688; 204771128)"

// SetUserAgent let you set User-Agent header in HTTP requests.
func SetUserAgent(s string) {
	appUserAgent = s
}

// NewInstagramApiManager returns a API manager of a logged-in Instagram user,
// given the JSON file of cookies of a Instagram logged-in account.
//
// The cookies, such as *ds_user_id*, *sessionid*, or *csrftoken* can be viewed
// in Chrome Developer Tools. See https://stackoverflow.com/a/44773079
//
// You can get the JSON format cookies using chrome extension in crx-cookies/
// directory.
func NewInstagramApiManager(authFilePath string) (*IGApiManager, error) {
	b, err := ioutil.ReadFile(authFilePath)
	if err != nil {
		return nil, err
	}

	cookies := make(map[string]string)
	err = json.Unmarshal(b, &cookies)
	if err != nil {
		return nil, err
	}

	headers := make(map[string]string)
	headers["User-Agent"] = appUserAgent

	return NewApiManager(cookies, headers), nil
}

// NewApiManager is a more low-level initialization function for API manager.
// The original goal of this method is used to create API manager used in Chrome
// extension.
func NewApiManager(cookies, headers map[string]string) *IGApiManager {
	var m IGApiManager
	m.cookies = cookies
	m.headers = headers
	return &m
}

// GetSelfId returns the id of a Instagram user of the API manager.
func (m *IGApiManager) GetSelfId() string {
	id, _ := m.cookies["ds_user_id"]
	return id
}
