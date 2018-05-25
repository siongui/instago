package instago

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

const urlGraphqlNoLogin = `https://www.instagram.com/graphql/query/?query_hash={{QUERYHASH}}&variables=`

func getGisHash(rhx_gis, variables string) string {
	h := md5.New()
	h.Write([]byte(rhx_gis + ":" + variables))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// Given user name, return IGMedia struct of all posts of the user without login
// status. The user account must be public.
func GetAllPostMediaNoLogin(username string) (medias []IGMedia, err error) {
	sd, qh, err := GetSharedDataQueryHashNoLogin(username)
	if err != nil {
		return
	}
	ui := sd.EntryData.ProfilePage[0].GraphQL.User

	for _, node := range ui.EdgeOwnerToTimelineMedia.Edges {
		medias = append(medias, node.Node)
	}

	if ui.IsPrivate {
		return
	}

	urlqh := strings.Replace(urlGraphqlNoLogin, "{{QUERYHASH}}", qh, 1)
	hasNextPage := ui.EdgeOwnerToTimelineMedia.PageInfo.HasNextPage
	vartmpl := strings.Replace(`{"id":"<ID>","first":50,"after":"<ENDCURSOR>"}`, "<ID>", ui.Id, 1)
	variables := strings.Replace(vartmpl, "<ENDCURSOR>", ui.EdgeOwnerToTimelineMedia.PageInfo.EndCursor, 1)

	for hasNextPage == true {
		url := urlqh + variables

		b, err := getHTTPResponseNoLoginWithGis(url, getGisHash(sd.RhxGis, variables))
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
