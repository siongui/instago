package instago

import (
	"fmt"
)

func ExampleGetIdFromWebStoryUrl() {
	m := NewApiManager(nil, nil)
	id, err := m.GetIdFromWebStoryUrl("https://www.instagram.com/stories/highlights/17862445040107085/")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(id)
	// Output:
	// 25025320
}

func ExampleGetInfoFromWebStoryUrl() {
	m := NewApiManager(nil, nil)
	info, err := m.GetInfoFromWebStoryUrl("https://www.instagram.com/stories/highlights/17862445040107085/")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(info.User.Username)
	// Output:
	// instagram
}
