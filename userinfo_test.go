package instago

import (
	"fmt"
	"os"
	"testing"
)

func ExampleGetUserInfoNoLogin() {
	user, err := GetUserInfoNoLogin("instagram")
	if err != nil {
		panic(err)
	}

	fmt.Println(user.Id)
	fmt.Println(user.Biography)
	// Output:
	// 25025320
	// Discovering — and telling — stories from around the world. Founded in 2010 by @kevin and @mikeyk.
}

func ExampleGetUserId() {
	fmt.Println(GetUserId("instagram"))
	// Output: 25025320 <nil>
}

func ExampleGetUserProfilePicUrlHd() {
	fmt.Println(GetUserProfilePicUrlHd("instagram"))
	// Output: https://instagram.fkhh1-2.fna.fbcdn.net/vp/893534d61bdc5ea6911593d3ee0a1922/5B6363AB/t51.2885-19/s320x320/14719833_310540259320655_1605122788543168512_a.jpg <nil>
}

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

func ExampleGetRecentPostCodeNoLogin(t *testing.T) {
	codes, err := GetRecentPostCodeNoLogin(os.Getenv("IG_TEST_USERNAME"))
	if err != nil {
		t.Error(err)
		return
	}
	for _, code := range codes {
		fmt.Printf("URL: https://www.instagram.com/p/%s/\n", code)
	}
}

func ExampleGetRecentPostCode(t *testing.T) {
	mgr := NewInstagramApiManager(
		os.Getenv("IG_DS_USER_ID"),
		os.Getenv("IG_SESSIONID"),
		os.Getenv("IG_CSRFTOKEN"))

	codes, err := mgr.GetRecentPostCode(os.Getenv("IG_TEST_USERNAME"))
	if err != nil {
		t.Error(err)
		return
	}
	for _, code := range codes {
		fmt.Printf("URL: https://www.instagram.com/p/%s/\n", code)
	}
}

func ExampleGetUserInfo(t *testing.T) {
	mgr := NewInstagramApiManager(
		os.Getenv("IG_DS_USER_ID"),
		os.Getenv("IG_SESSIONID"),
		os.Getenv("IG_CSRFTOKEN"))

	ui, err := mgr.GetUserInfo(os.Getenv("IG_TEST_USERNAME"))
	if err != nil {
		t.Error(err)
		return
	}
	jsonPrettyPrint(ui)
}
