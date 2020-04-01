package igdl

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/siongui/instago"
)

var cc = color.New(color.FgCyan)
var rc = color.New(color.FgRed)

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
