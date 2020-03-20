package instago

import (
	"fmt"
	"os"
	"testing"
)

func ExampleGetUserInfoNoLogin() {
	user, err := GetUserInfoNoLogin("instagram")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(user.Id)
	fmt.Println(user.Biography)
	// Output:
	// 25025320
	// Bringing you closer to the people and things you love. ❤️
}

func ExampleGetUserId() {
	fmt.Println(GetUserId("instagram"))
	// Output: 25025320 <nil>
}

func ExampleGetUserProfilePicUrlHd() {
	fmt.Println(GetUserProfilePicUrlHd("instagram"))
	// Output: https://instagram.ftpe1-1.fna.fbcdn.net/v/t51.2885-19/s320x320/59381178_2348911458724961_5863612957363011584_n.jpg?_nc_ht=instagram.ftpe1-1.fna.fbcdn.net&oh=40b5f2e0329d70e72aa8adef38b17291&oe=5E8D1B25 <nil>
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
	mgr, err := NewInstagramApiManager("auth.json")
	if err != nil {
		t.Error(err)
		return
	}

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
	mgr, err := NewInstagramApiManager("auth.json")
	if err != nil {
		t.Error(err)
		return
	}

	ui, err := mgr.GetUserInfo(os.Getenv("IG_TEST_USERNAME"))
	if err != nil {
		t.Error(err)
		return
	}
	jsonPrettyPrint(ui)
}

func ExampleGetSharedDataQueryHashNoLogin() {
	sd, qh, err := GetSharedDataQueryHashNoLogin("instagram")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(sd.EntryData.ProfilePage[0].GraphQL.User.Id)
	fmt.Println(qh)
	// Output:
	// 25025320
	// e769aa130647d2354c40ea6a439bfc08
}
