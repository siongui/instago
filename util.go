package instago

import (
	"net/url"
	"time"
)

func formatTimestamp(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format(time.RFC3339)
}

// Remove query string in the URL
func stripQueryString(inputUrl string) (su string, err error) {
	u, err := url.Parse(inputUrl)
	if err != nil {
		return
	}
	u.RawQuery = ""
	su = u.String()
	return
}
