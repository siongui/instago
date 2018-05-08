package instago

import (
	"encoding/json"
	"errors"
)

const urlReelsMedia = `https://i.instagram.com/api/v1/feed/reels_media/`

type highlightMedia struct {
	Trays []IGStoryHighlightsTray `json:"reels_media"`
}

func (m *IGApiManager) GetReelsMedia(id string) (tray IGStoryHighlightsTray, err error) {
	url := urlReelsMedia + "?user_ids=" + id
	b, err := getHTTPResponse(url, m.dsUserId, m.sessionid, m.csrftoken)
	if err != nil {
		return
	}

	h := highlightMedia{}
	err = json.Unmarshal(b, &h)
	if err != nil {
		return
	}
	if len(h.Trays) == 0 {
		err = errors.New("reels_media has no highlight tray!")
		return
	}
	tray = h.Trays[0]
	return
}
