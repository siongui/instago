package igdl

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/siongui/instago"
)

var cc = color.New(color.FgCyan)
var rc = color.New(color.FgRed)

func CyanPrint(a ...interface{}) {
	cc.Print(a...)
}

func RedPrint(a ...interface{}) {
	rc.Print(a...)
}

func CyanPrintln(a ...interface{}) {
	cc.Println(a...)
}

func RedPrintln(a ...interface{}) {
	rc.Println(a...)
}

func SleepSecond(interval int) {
	time.Sleep(time.Duration(interval) * time.Second)
}

func SleepLog(interval int) {
	rc.Print(time.Now().Format(time.RFC3339))
	fmt.Print(": sleep ")
	cc.Print(interval)
	fmt.Println(" second")
	SleepSecond(interval)
}

func PrintMsgSleep(sleepInterval int, msg string) {
	fmt.Print(msg)
	SleepLog(sleepInterval)
}

func UsernameIdColorPrint(username, id interface{}) {
	CyanPrint(username)
	fmt.Print(" (")
	RedPrint(id)
	fmt.Print(") ")
}

func PrintUsernameIdMsg(username, id interface{}, msg ...interface{}) {
	UsernameIdColorPrint(username, id)
	fmt.Println(msg...)
}

func PrintBestiesItemInfo(item instago.IGItem, username string) {
	PrintUsernameIdMsg(username, item.GetUserId(), "'s besties item")
}

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

func PrintReelMentionInfo(rm instago.ItemReelMention) {
	fmt.Print("Reel Mention ")
	UsernameIdColorPrint(rm.GetUsername(), rm.GetUserId())
	fmt.Print(", Display Type: ")
	CyanPrint(rm.DisplayType)
	fmt.Print(" , ")
	if rm.User.IsPrivate {
		RedPrintln("Private")
	} else {
		RedPrintln("Public")
	}
}

func PrintUserComment(user instago.User) {
	fmt.Println("// https://www.instagram.com/" + user.GetUsername() + "/")
	fmt.Println("//" + user.GetUserId() + "-" + user.GetUsername())
}

func PrintUserInfo(user instago.User) {
	fmt.Println("----")
	fmt.Print("User: ")
	UsernameIdColorPrint(user.GetUsername(), user.GetUserId())
	fmt.Print(", ")
	if user.IsPublic() {
		RedPrintln("Public")
	} else {
		RedPrintln("Private")
	}
	fmt.Println("https://www.instagram.com/" + user.GetUsername() + "/")
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
	cc.Println(instago.FormatTimestamp(pi.GetTimestamp()))
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

		fmt.Print("Prefix: " + bc.BroadcastOwner.Username)
		fmt.Print("-")
		fmt.Print(bc.BroadcastOwner.Pk)
		for _, cobcter := range bc.Cobroadcasters {
			fmt.Print("-")
			fmt.Print(cobcter.Username)
			fmt.Print("-")
			fmt.Print(cobcter.Pk)
		}
		fmt.Println("-")
	}
}

func PrintPostLiveItem(pli instago.IGPostLiveItem) {
	fmt.Print("Postlive Item, Pk: ")
	cc.Print(pli.Pk)
	fmt.Print(" , user: ")
	cc.Print(pli.User.Username)
	fmt.Println(" , ranked_position: ", pli.RankedPosition)
}

func PrintDownloadStoryLayerInfo(item instago.IGItem, username string) {
	fmt.Print("Download story of ")
	cc.Print(username)
	fmt.Print(" , id: ")
	cc.Println(item.GetUserId())
}
