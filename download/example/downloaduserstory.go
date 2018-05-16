package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/siongui/instago/download"
)

func main() {
	mgr, err := igdl.NewInstagramDownloadManager(
		os.Getenv("IG_DS_USER_ID"),
		os.Getenv("IG_SESSIONID"),
		os.Getenv("IG_CSRFTOKEN"))
	if err != nil {
		fmt.Println(err)
		return
	}

	id := flag.Int64("id", 25025320, "user id")
	flag.Parse()
	mgr.DownloadUserStory(*id)
}
