package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/siongui/instago/download"
)

func main() {
	mgr, err := igdl.NewInstagramDownloadManager("auth.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	num := flag.Int("num", 12, "number of saved posts to be downloaded")
	flag.Parse()

	sleepInterval := 12 // seconds

	for {
		fmt.Println("Download saved post:", *num)
		// The following method will download given number of saved
		// posts. -1 will download all saved posts. second parameter
		// true means also download unexpired stories of the post user
		mgr.DownloadSavedPosts(*num, true)
		fmt.Println("=========================")
		fmt.Print(time.Now().Format(time.RFC3339))
		fmt.Println(": sleep ", sleepInterval, " second")
		fmt.Println("=========================")
		time.Sleep(time.Duration(sleepInterval) * time.Second)
	}
}
