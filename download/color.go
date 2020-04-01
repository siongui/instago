package igdl

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/siongui/instago"
)

var cc = color.New(color.FgCyan)
var rc = color.New(color.FgRed)

func printTimelineItemInfo(index int, item instago.IGItem) {
	// You're All Caught Up
	if item.EndOfFeedDemarcator.Title != "" {
		cc.Print(index)
		fmt.Println(": You're All Caught Up")
		return
	}

	// injected ads
	if item.Injected.Label != "" {
		cc.Print(index)
		fmt.Print(": injected ad: ")
		printItemInfo(index, &item)
		return
	}

	// suggested_user
	if item.Type == 2 {
		cc.Print(index)
		fmt.Print(": suggested_user: ")
		for i, suggestion := range item.Suggestions {
			if i > 4 {
				break
			}
			rc.Print(suggestion.User.Username)
			fmt.Print(", ")
		}
		fmt.Println("")
		return
	}

	// item from following users
	printItemInfo(index, &item)
}

func printItemInfo(index int, pi instago.PostItem) {
	cc.Print(index)
	fmt.Print(": username: ")
	rc.Print(pi.GetUsername())
	fmt.Print(" , post url: ")
	cc.Println(pi.GetPostUrl())
}

func printDownloadInfo(pi instago.PostItem, url, filepath string) {
	fmt.Print("username: ")
	cc.Println(pi.GetUsername())
	fmt.Print("time: ")
	cc.Println(formatTimestamp(pi.GetTimestamp()))
	fmt.Print("post url: ")
	cc.Println(pi.GetPostUrl())

	fmt.Print("Download ")
	rc.Print(url)
	fmt.Print(" to ")
	cc.Println(filepath)
}
