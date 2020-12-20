package instago

// Get all unexpired stories of a specific user, without postlive

import (
	"encoding/json"
	"strings"
)

const urlUserReelMedia = `https://i.instagram.com/api/v1/feed/user/{{USERID}}/reel_media/`

// GetUserReelMedia returns unexpired stories of the given user id. No postlives
// included.
func (m *IGApiManager) GetUserReelMedia(id string) (tray IGReelTray, err error) {
	url := strings.Replace(urlUserReelMedia, "{{USERID}}", id, 1)
	b, err := m.getHTTPResponse(url, "GET")
	if err != nil {
		return
	}

	// for development purpose
	if saveRawJsonByte {
		SaveRawJsonByte(id+"-reel_media-", b)
	}

	err = json.Unmarshal(b, &tray)
	return
}
