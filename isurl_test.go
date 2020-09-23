package instago

import (
	"fmt"
)

func ExampleIsWebStoryUrl() {
	is1 := IsWebStoryUrl("https://www.instagram.com/stories/highlights/17862445040107085/")
	fmt.Println(is1)

	is2 := IsWebStoryUrl("https://www.instagram.com/stories/instagram/")
	fmt.Println(is2)

	is3 := IsWebStoryUrl("https://www.instagram.com/story/instagram/")
	fmt.Println(is3)

	// Output:
	// true
	// true
	// false
}
