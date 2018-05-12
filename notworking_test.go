package instago

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"
	"time"
)

const urlGraphqlNoLogin = `https://www.instagram.com/graphql/query/?query_hash={{QUERYHASH}}&variables=`

// Given the HTML source code of the user profile page without logged in, return
// query_hash for Instagram GraphQL API.
func GetQueryHashNoLogin(b []byte) (qh string, err error) {
	// find JavaScript file which contains the query hash
	patternJs := regexp.MustCompile(`\/static\/bundles\/base\/ProfilePageContainer\.js\/[a-zA-Z0-9]+?\.js`)
	jsPath := string(patternJs.Find(b))
	jsUrl := "https://www.instagram.com" + jsPath
	bJs, err := getHTTPResponseNoLogin(jsUrl)
	if err != nil {
		return
	}

	patternQh := regexp.MustCompile(`e\.profilePosts\.byUserId\.get\(t\)\)\?n\.pagination:n},queryId:"([a-zA-Z0-9]+)",`)
	qhtmp := string(patternQh.Find(bJs))
	qhtmp = strings.TrimPrefix(qhtmp, `e.profilePosts.byUserId.get(t))?n.pagination:n},queryId:"`)
	qh = strings.TrimSuffix(qhtmp, `",`)
	return
}

// Given user name, return IGMedia struct of all posts of the user without login
// status. The user account must be public.
func GetAllPostMediaNoLogin(username string) (medias []IGMedia, err error) {
	ui, err := GetUserInfoNoLogin(username)
	if err != nil {
		return
	}

	for _, node := range ui.EdgeOwnerToTimelineMedia.Edges {
		medias = append(medias, node.Node)
	}

	if ui.IsPrivate {
		return
	}

	urlqh := strings.Replace(urlGraphqlNoLogin, "{{QUERYHASH}}", ui.QueryHash, 1)
	hasNextPage := ui.EdgeOwnerToTimelineMedia.PageInfo.HasNextPage
	vartmpl := strings.Replace(`{"id":"<ID>","first":50,"after":"<ENDCURSOR>"}`, "<ID>", ui.Id, 1)
	variables := strings.Replace(vartmpl, "<ENDCURSOR>", ui.EdgeOwnerToTimelineMedia.PageInfo.EndCursor, 1)

	for hasNextPage == true {
		url := urlqh + variables

		b, err := getHTTPResponseNoLogin(url)
		if err != nil {
			return medias, err
		}

		d := dataUserMedia{}
		if err = json.Unmarshal(b, &d); err != nil {
			return medias, err
		}

		for _, node := range d.Data.User.EdgeOwnerToTimelineMedia.Edges {
			medias = append(medias, node.Node)
		}
		hasNextPage = d.Data.User.EdgeOwnerToTimelineMedia.PageInfo.HasNextPage
		variables = strings.Replace(vartmpl, "<ENDCURSOR>", d.Data.User.EdgeOwnerToTimelineMedia.PageInfo.EndCursor, 1)

		printPostCount(len(medias), url)

		if hasNextPage {
			// sleep 20 seconds to prevent http 429 (too many requests)
			time.Sleep(20 * time.Second)
		}
	}
	return
}

func ExampleGetAllPostMediaNoLogin(t *testing.T) {
	medias, err := GetAllPostMediaNoLogin(os.Getenv("IG_TEST_USERNAME"))
	if err != nil {
		t.Error(err)
		return
	}

	for _, media := range medias {
		fmt.Printf("URL: https://www.instagram.com/p/%s/\n", media.Shortcode)
	}
}
