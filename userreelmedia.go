package instago

// Get a user's unexpired stories and postlive

import (
	"encoding/json"
	"strings"
)

const urlUserReelMedia = `https://i.instagram.com/api/v1/feed/user/{{USERID}}/story/`

type userTray struct {
	Reel         IGReelTray     `json:"reel"`
	PostLiveItem IGPostLiveItem `json:"post_live_item"`

	Status string `json:"status"`
}

func (m *IGApiManager) GetUserReelMedia(userid string) (ut userTray, err error) {
	url := strings.Replace(urlUserReelMedia, "{{USERID}}", userid, 1)
	b, err := m.getHTTPResponse(url, "GET")
	if err != nil {
		return
	}

	// for development purpose
	if saveRawJsonByte {
		SaveRawJsonByte(userid+"-story-", b)
	}

	err = json.Unmarshal(b, &ut)
	return
}
