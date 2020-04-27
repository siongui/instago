package igdl

import (
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/siongui/instago"
)

var outputDir = "Instagram"
var dataDir = "Data"

func SetOutputDir(s string) {
	outputDir = s
}

func SetDataDir(s string) {
	dataDir = s
}

func formatTimestamp(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format(time.RFC3339)
}

func buildFilename(url, username, id, middle, last string, timestamp int64) string {
	url, err := instago.StripQueryString(url)
	if err != nil {
		panic(err)
	}

	ext := path.Ext(path.Base(url))
	return path.Join(username + "-" + id +
		middle +
		formatTimestamp(timestamp) + "-" +
		last +
		strconv.FormatInt(timestamp, 10) +
		ext)
}

func getPostFilePath2(username, id, code, url string, timestamp int64, taggedusers [][2]string) string {
	userDir := path.Join(outputDir, username)
	userPostsDir := path.Join(userDir, "posts")

	filename := buildFilename(url, username, id, "-post-", code+"-", timestamp)
	filename = appendUsernameToFilename(username, id, filename, taggedusers)

	return path.Join(userPostsDir, filename)
}

func getPostFilePath(username, id, code, url string, timestamp int64) string {
	userDir := path.Join(outputDir, username)
	userPostsDir := path.Join(userDir, "posts")
	return path.Join(userPostsDir, buildFilename(url, username, id, "-post-", code+"-", timestamp))
}

func getStoryFilePath(username, id, code, url string, timestamp int64) string {
	userDir := path.Join(outputDir, username)
	userStoriesDir := path.Join(userDir, "stories")
	return path.Join(userStoriesDir, buildFilename(url, username, id, "-story-", code+"-", timestamp))
}

func appendUsernameToFilename(username, id, filename string, appendIdUsernames [][2]string) string {
	prefix := username + "-" + id

	usednames := make(map[string]bool)
	usednames[username] = true
	for _, n := range appendIdUsernames {
		taggedname := n[1]
		newprefix := prefix + "-" + taggedname
		newfilename := strings.Replace(filename, prefix, newprefix, 1)

		// cannot use 256 here. will get filename too long error.
		// use 240
		if len(newfilename) > 240 {
			continue
		}

		if _, ok := usednames[taggedname]; ok {
			continue
		} else {
			usednames[taggedname] = true
		}

		prefix = newprefix
		filename = newfilename
	}

	return filename
}

// same as getStoryFilePath, except adding usernames in reel_mentions
func getStoryFilePath2(username, id, code, url string, timestamp int64, rms []instago.ItemReelMention) string {
	userDir := path.Join(outputDir, username)
	userStoriesDir := path.Join(userDir, "stories")

	filename := buildFilename(url, username, id, "-story-", code+"-", timestamp)
	appendIdUsernames := [][2]string{}
	for _, rm := range rms {
		pair := [2]string{rm.GetUserId(), rm.GetUsername()}
		appendIdUsernames = append(appendIdUsernames, pair)
	}
	filename = appendUsernameToFilename(username, id, filename, appendIdUsernames)

	return path.Join(userStoriesDir, filename)
}

func getPostLiveFilePath(username, id, url, typ string, timestamp int64) string {
	userDir := path.Join(outputDir, username)
	userPostLiveDir := path.Join(userDir, "postlives")
	return path.Join(userPostLiveDir, buildFilename(url, username, id, "-postlive-"+typ+"-", "", timestamp))
}

func getPostLiveMergedFilePath(vpath, apath string) string {
	filename := path.Base(vpath)
	filename = strings.Replace(filename, "video-", "", 1)
	return path.Join(path.Dir(vpath), filename)
}

// only for post/item with several photos/videos of the same TakenAt time
func appendIndexToFilename(filename string, index int) string {
	ext := path.Ext(filename)
	fne := strings.TrimSuffix(filename, ext)
	return fne + "-" + strconv.Itoa(index) + ext
}

func CreateFilepathDirIfNotExist(filepath string) {
	dir := path.Dir(filepath)
	err := CreateDirIfNotExist(dir)
	if err != nil {
		panic(err)
	}
}

func getUserProfilPicFilePath(username, id, url string, timestamp int64) string {
	url, err := instago.StripQueryString(url)
	if err != nil {
		panic(err)
	}

	ext := path.Ext(url)
	userDir := path.Join(outputDir, username)
	filename := username + "-" + id + "-profile_pic-" + strconv.FormatInt(timestamp, 10) + ext
	return path.Join(userDir, filename)
}

func GetIdUsernamePath(id, username string) string {
	filename := id + "-" + username
	return path.Join(dataDir, "ID-USERNAME", filename)
}

func GetReelMentionsPath(id, username string) string {
	filename := id + "-" + username
	return path.Join(dataDir, "Reel-Mentions", filename)
}

func GetScreenshotPath(id, username string) string {
	filename := username + "-" + id + "-screenshot.png"
	return path.Join(outputDir, "Auto-Screenshot", filename)
}

func getFollowingPath(id string) string {
	filename := id + "-following-" + time.Now().Format(time.RFC3339) + ".json"
	return path.Join(dataDir, "Follow", filename)
}

func getFollowersPath(id string) string {
	filename := id + "-followers-" + time.Now().Format(time.RFC3339) + ".json"
	return path.Join(dataDir, "Follow", filename)
}

func GetRFC3339String(s string) string {
	// Google search: regex rfc3339 golang
	pattern := regexp.MustCompile(`([0-9]+)-(0[1-9]|1[012])-(0[1-9]|[12][0-9]|3[01])[Tt]([01][0-9]|2[0-3]):([0-5][0-9]):([0-5][0-9]|60)(\.[0-9]+)?(([Zz])|([\+|\-]([01][0-9]|2[0-3]):[0-5][0-9]))`)
	return string(pattern.Find([]byte(s)))
}

func ExtractPostCodeFromFilename(filename string) (code string) {
	// remove ext
	f1 := strings.TrimSuffix(filename, path.Ext(filename))

	rfc3339s := GetRFC3339String(f1)
	pieces := strings.Split(f1, "-"+rfc3339s+"-")
	if len(pieces) != 2 {
		return
	}

	f2 := pieces[1]
	pieces = strings.Split(f2, "-")
	if len(pieces) < 2 {
		return
	}

	utime := pieces[len(pieces)-1]
	if len(utime) < 3 && len(pieces) > 2 {
		utime = pieces[len(pieces)-2]
	}

	pieces = strings.Split(f2, "-"+utime)
	if len(pieces) < 1 {
		return
	}

	return pieces[0]
}

func ExtractUsernameIdFromFilename(filename string) (username, id string) {
	pieces := strings.Split(filename, "-")
	if len(pieces) < 2 {
		return
	}
	username = pieces[0]
	id = pieces[1]
	return
}
