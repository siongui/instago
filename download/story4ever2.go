package igdl

import (
	"fmt"
	"log"
)

func (m *IGDownloadManager) AccessReelsTrayOnce2Chan(cPublicUser, cPrivateUser chan TrayInfo, ignoreMuted, verbose bool) (err error) {
	rt, err := m.GetReelsTray()
	if err != nil {
		log.Println(err)
		return
	}

	go PrintLiveBroadcasts(rt.Broadcasts)

	for index, tray := range rt.Trays {
		fmt.Print(index, ":")

		username := tray.GetUsername()
		id := tray.Id
		//items := tray.GetItems()

		if ignoreMuted && tray.Muted {
			if verbose {
				PrintUsernameIdMsg(username, id, " is muted && ignoreMuted set. no download")
			}
			continue
		}

		if IsLatestReelMediaDownloaded(username, tray.LatestReelMedia) {
			if verbose {
				PrintUsernameIdMsg(username, id, " all downloaded")
			}
			continue
		}

		if tray.HasBestiesMedia {
			PrintUsernameIdMsg(username, id, "has close friend (besties) story item(s)")
		}

		if verbose {
			UsernameIdColorPrint(username, id)
			fmt.Println(" has undownloaded items")
		}

		// 2: also download reel mentions in story item
		if tray.User.IsPrivate {
			cPrivateUser <- setupTrayInfo(id, username, 2, tray.User.IsPrivate)
		} else {
			cPublicUser <- setupTrayInfo(id, username, 2, tray.User.IsPrivate)
		}
	}

	return
}

func (m *IGDownloadManager) TwoAccountDownloadStoryForever(interval1 int, interval2 int64, ignoreMuted, verbose bool) {
	if !m.IsCleanAccountSet() {
		fmt.Println("clean account not set. exit")
		return
	}

	// usually there are at most 150 trays in reels_tray.
	// double the buffer to 300. 160 or 200 may be ok as well.
	cPublicUser := make(chan TrayInfo, 300)
	cPrivateUser := make(chan TrayInfo, 300)

	tl := NewTimeLimiter(interval2)
	go m.GetCleanAccountManager().TrayDownloader(cPublicUser, tl, verbose)

	for {
		err := m.AccessReelsTrayOnce2Chan(cPublicUser, cPrivateUser, ignoreMuted, verbose)
		if err != nil {
			log.Println(err)
		}
		PrintMsgSleep(interval1, "TwoAccountDownloadStoryForever: ")
	}
}
