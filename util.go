package instago

import (
	"net/url"
	"time"
)

func FormatTimestamp(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format(time.RFC3339)
}

// StripQueryString removes query string in the URL
func StripQueryString(inputUrl string) (su string, err error) {
	u, err := url.Parse(inputUrl)
	if err != nil {
		return
	}
	u.RawQuery = ""
	su = u.String()
	return
}
