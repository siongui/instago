package instago

// Get a user's unexpired stories and postlive

import (
	"encoding/json"
	"strings"
)

const urlUserStory = `https://i.instagram.com/api/v1/feed/user/{{USERID}}/story/`

type UserTray struct {
	Reel         IGReelTray     `json:"reel"`
	PostLiveItem IGPostLiveItem `json:"post_live_item"`

	Status string `json:"status"`
}

func (m *IGApiManager) GetUserStory(userid string) (ut UserTray, err error) {
	url := strings.Replace(urlUserStory, "{{USERID}}", userid, 1)
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
