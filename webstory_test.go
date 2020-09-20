package instago

import (
	"fmt"
)

func ExampleGetIdFromWebStoryUrl() {
	id, err := GetIdFromWebStoryUrl("https://www.instagram.com/stories/highlights/17862445040107085/")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(id)
	// Output:
	// 25025320
}

func ExampleGetInfoFromWebStoryUrl() {
	info, err := GetInfoFromWebStoryUrl("https://www.instagram.com/stories/highlights/17862445040107085/")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(info.User.Username)
	// Output:
	// instagram
}
