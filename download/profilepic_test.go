package igdl

import (
	"os"
	"testing"
)

func ExampleDownloadUserProfilePicUrlHd(t *testing.T) {
	DownloadUserProfilePicUrlHd(os.Getenv("IG_TEST_USERNAME"))
}
