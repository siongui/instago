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
		PrintItemInfo(index, &item)
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
	PrintItemInfo(index, &item)
}

func PrintItemInfo(index int, pi instago.PostItem) {
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

func PrintLiveBroadcasts(bcs []instago.IGBroadcast) {
	for _, bc := range bcs {
		fmt.Print("Live: ")
		cc.Print(bc.BroadcastOwner.Username)
		for _, cobcter := range bc.Cobroadcasters {
			fmt.Print(" + ")
			cc.Print(cobcter.Username)
		}
		fmt.Print(" , id: ")
		cc.Print(bc.Id)
		fmt.Println(" , viewer_count: ", bc.ViewerCount)
		fmt.Println("dash_playback_url: ", bc.DashPlaybackUrl)
		fmt.Println("dash_live_predictive_playback_url: ", bc.DashLivePredictivePlaybackUrl)
		fmt.Println("cover_frame_url", bc.CoverFrameUrl)
	}
}

func PrintPostLiveItem(pli instago.IGPostLiveItem) {
	fmt.Print("Postlive Item, Pk: ")
	cc.Print(pli.Pk)
	fmt.Print(" , user: ")
	cc.Print(pli.User.Username)
	fmt.Println(" , ranked_position: ", pli.RankedPosition)
}
