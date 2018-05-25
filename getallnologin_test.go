package instago

import (
	"fmt"
	"os"
	"testing"
)

func ExampleGetAllPostMediaNoLogin(t *testing.T) {
	medias, err := GetAllPostMediaNoLogin(os.Getenv("IG_TEST_USERNAME"))
	if err != nil {
		t.Error(err)
		return
	}

	for _, media := range medias {
		fmt.Printf("URL: https://www.instagram.com/p/%s/\n", media.Shortcode)
	}
}
