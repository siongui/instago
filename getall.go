package instago

import (
	"encoding/json"
	"strings"
	"time"
)

// Given user name, return IGMedia struct of all posts of the user with logged in status.
func (m *IGApiManager) GetAllPostMedia(username string) (medias []IGMedia, err error) {
	ui, err := m.GetUserInfo(username)
	if err != nil {
		return
	}

	for _, node := range ui.EdgeOwnerToTimelineMedia.Edges {
		medias = append(medias, node.Node)
	}

	hasNextPage := ui.EdgeOwnerToTimelineMedia.PageInfo.HasNextPage
	// "first" cannot be 300 now. cannot be 100 either. 50 is ok.
	vartmpl := strings.Replace(`{"id":"<ID>","first":50,"after":"<ENDCURSOR>"}`, "<ID>", ui.Id, 1)
	variables := strings.Replace(vartmpl, "<ENDCURSOR>", ui.EdgeOwnerToTimelineMedia.PageInfo.EndCursor, 1)

	for hasNextPage == true {
		url := urlGraphql + variables

		b, err := getHTTPResponse(url, m.dsUserId, m.sessionid, m.csrftoken)
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

// Given user name, return codes of all posts of the user with logged in status.
func (m *IGApiManager) GetAllPostCode(username string) (codes []string, err error) {
	medias, err := m.GetAllPostMedia(username)
	if err != nil {
		return
	}

	for _, media := range medias {
		codes = append(codes, media.Shortcode)
	}
	return
}

// Given a user id, return all story highlights of the user.
func (m *IGApiManager) GetAllStoryHighlights(userid string) (trays []IGStoryHighlightsTray, err error) {
	subtrays, err := m.GetUserStoryHighlights(userid)
	if err != nil {
		return
	}
	for _, tray := range subtrays {
		if len(tray.GetItems()) == 0 {
			tt, err := m.GetReelsMedia(tray.Id)
			if err != nil {
				return trays, err
			}
			trays = append(trays, tt)
		} else {
			trays = append(trays, tray)
		}
	}
	return
}
