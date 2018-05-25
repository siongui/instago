package instago

// Get all unexpired stories of a specific user

import (
	"encoding/json"
	"strings"
)

const urlUserStory = `https://i.instagram.com/api/v1/feed/user/{{USERID}}/reel_media/`

// GetUserStory returns unexpired stories of the given user id.
func (m *IGApiManager) GetUserStory(id string) (tray IGReelTray, err error) {
	url := strings.Replace(urlUserStory, "{{USERID}}", id, 1)
	b, err := getHTTPResponse(url, m.dsUserId, m.sessionid, m.csrftoken)
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &tray)
	return
}
