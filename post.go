package instago

// This file decode JSON data returned by Instagram post API

import (
	"encoding/json"
	"strings"
)

const urlPost = `https://www.instagram.com/p/{{CODE}}/?__a=1`

// Decode JSON data returned by Instagram post API
type postInfo struct {
	GraphQL struct {
		ShortcodeMedia IGMedia `json:"shortcode_media"`
	} `json:"graphql"`
}

// Given the code of the post, return url of the post.
func CodeToUrl(code string) string {
	return strings.Replace(`https://www.instagram.com/p/{{CODE}}/`, "{{CODE}}", code, 1)
}

// Given code of post, return information of the post with login status.
func (m *IGApiManager) GetPostInfo(code string) (em IGMedia, err error) {
	url := strings.Replace(urlPost, "{{CODE}}", code, 1)
	b, err := m.getHTTPResponse(url, "GET")
	if err != nil {
		return
	}

	// for development purpose
	if saveRawJsonByte {
		SaveRawJsonByte("post-"+code+"-with-login-", b)
	}

	pi := postInfo{}
	err = json.Unmarshal(b, &pi)
	if err != nil {
		return
	}
	em = pi.GraphQL.ShortcodeMedia
	return
}

// Given code of post, return information of the post without login status.
func GetPostInfoNoLogin(code string) (em IGMedia, err error) {
	url := strings.Replace(urlPost, "{{CODE}}", code, 1)
	b, err := getHTTPResponseNoLogin(url)
	if err != nil {
		return
	}

	// for development purpose
	if saveRawJsonByte {
		SaveRawJsonByte("post-"+code+"-no-login-", b)
	}

	pi := postInfo{}
	err = json.Unmarshal(b, &pi)
	if err != nil {
		return
	}
	em = pi.GraphQL.ShortcodeMedia
	return
}
