package instago

import (
	"encoding/json"
	"strings"
)

const urlUserReelMedia = `https://i.instagram.com/api/v1/feed/user/{{USERID}}/reel_media/`

func (m *IGApiManager) GetUserReelMedia(userid string) (tray IGReelTray, err error) {
	url := strings.Replace(urlUserReelMedia, "{{USERID}}", userid, 1)
	b, err := getHTTPResponse(url, m.dsUserId, m.sessionid, m.csrftoken)
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &tray)
	return
}
