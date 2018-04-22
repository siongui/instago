package igdl

import (
	"github.com/siongui/instago"
	"os"
	"testing"
)

func ExampleDownloadPost(t *testing.T) {
	mgr := instago.NewInstagramApiManager(
		os.Getenv("IG_DS_USER_ID"),
		os.Getenv("IG_SESSIONID"),
		os.Getenv("IG_CSRFTOKEN"))
	DownloadPost(os.Getenv("IG_TEST_CODE"), mgr)
}

func ExampleDownloadPostNoLogin(t *testing.T) {
	DownloadPostNoLogin(os.Getenv("IG_TEST_CODE"))
}
