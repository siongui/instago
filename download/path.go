package igdl

import (
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/siongui/instago"
)

var outputDir = "Instagram"

func SetOutputDir(s string) {
	outputDir = s
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
