package instago

import (
	"fmt"
	"os"
	"testing"
)

func ExampleGetAllPostCode(t *testing.T) {
	mgr := NewInstagramApiManager(
		os.Getenv("IG_DS_USER_ID"),
		os.Getenv("IG_SESSIONID"),
		os.Getenv("IG_CSRFTOKEN"))

	codes, err := mgr.GetAllPostCode(os.Getenv("IG_TEST_USERNAME"))
	if err != nil {
		t.Error(err)
		return
	}
	for _, code := range codes {
		fmt.Printf("URL: https://www.instagram.com/p/%s/\n", code)
	}
}
