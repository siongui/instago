package instago

import (
	"testing"
)

func TestStripQueryString(t *testing.T) {
	u, err := StripQueryString("https://example.com/myvideo.mp4?abc=d")
	if err != nil {
		t.Error(err)
		return
	}
	if u != "https://example.com/myvideo.mp4" {
		t.Error(u)
	}

	u, err = StripQueryString("https://example.com/myvideo.mp4")
	if err != nil {
		t.Error(err)
		return
	}
	if u != "https://example.com/myvideo.mp4" {
		t.Error(u)
	}
}
