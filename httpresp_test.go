package instago

import (
	"testing"
)

func TestSetUserAgent(t *testing.T) {
	SetUserAgent("hello world")
	if userAgent != "hello world" {
		t.Error(userAgent)
	}
}
