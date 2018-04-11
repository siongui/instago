package igdl

import (
	"path"
	"strconv"
	"strings"
	"time"
)

var outputDir = "Instagram"

func SetOutputDir(s string) {
	outputDir = s
}

func formatTimestamp(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format(time.RFC3339)
}

func buildFilepath(url, dir, username, middle string, timestamp int64) string {
	ext := path.Ext(path.Base(url))
	return path.Join(dir, username+
		middle+
		formatTimestamp(timestamp)+"-"+
		strconv.FormatInt(timestamp, 10)+
		ext)
}

func GetPostFilePath(username, url string, timestamp int64) string {
	userDir := path.Join(outputDir, username)
	userPostsDir := path.Join(userDir, "posts")
	return buildFilepath(url, userPostsDir, username, "-post-", timestamp)
}

func getStoryFilePath(username, url string, timestamp int64) string {
	userDir := path.Join(outputDir, username)
	userStoriesDir := path.Join(userDir, "stories")
	return buildFilepath(url, userStoriesDir, username, "-story-", timestamp)
}

func getPostLiveFilePath(username, url, typ string, timestamp int64) string {
	userDir := path.Join(outputDir, username)
	userPostLiveDir := path.Join(userDir, "postlives")
	return buildFilepath(url, userPostLiveDir, username, "-postlive-"+typ+"-", timestamp)
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
