package instago

import (
	"encoding/json"
	"strings"
)

const urlUsers = `https://i.instagram.com/api/v1/users/{{USERID}}/info/`

type rawUserinfoResp struct {
	User IGUser `json:"user"`
}

// FIXME: do not return IGUser
func (m *IGApiManager) GetUserInfoEndPoint(userid string) (user IGUser, err error) {
	url := strings.Replace(urlUsers, "{{USERID}}", userid, 1)
	b, err := m.getHTTPResponse(url, "GET")
	if err != nil {
		return
	}

	// for development purpose
	if saveRawJsonByte {
		SaveRawJsonByte(userid+"-info-", b)
	}

	r := rawUserinfoResp{}
	err = json.Unmarshal(b, &r)
	if err == nil {
		user = r.User
	}
	return
}
